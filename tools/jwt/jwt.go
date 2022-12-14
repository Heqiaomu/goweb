package jwt

import (
	"errors"
	"github.com/Heqiaomu/goweb/config"
	"github.com/Heqiaomu/goweb/tools/timer"
	uuid "github.com/satori/go.uuid"
	"time"

	jwtV4 "github.com/golang-jwt/jwt/v4"

	"golang.org/x/sync/singleflight"
)

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwtV4.RegisteredClaims
}

type BaseClaims struct {
	UUID        uuid.UUID
	UserID      uint
	Username    string
	NickName    string
	AuthorityId uint // 用户角色ID
}

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

var sg *singleflight.Group

func init() {
	sg = &singleflight.Group{}
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(config.GetConfig().JWT.SigningKey),
	}
}

func (j *JWT) CreateClaims(jwtCfg config.JWT, baseClaims BaseClaims) CustomClaims {
	bf, _ := timer.ParseDuration(jwtCfg.BufferTime)
	ep, _ := timer.ParseDuration(jwtCfg.ExpiresTime)
	claims := CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(bf / time.Second), // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		RegisteredClaims: jwtV4.RegisteredClaims{
			NotBefore: jwtV4.NewNumericDate(time.Now().Add(time.Second * -1000)), // 签名生效时间
			ExpiresAt: jwtV4.NewNumericDate(time.Now().Add(ep)),                  // 过期时间 7天  配置文件
			Issuer:    jwtCfg.Issuer,                                             // 签名的发行者
		},
	}
	return claims
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwtV4.NewWithClaims(jwtV4.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims CustomClaims) (string, error) {
	v, err, _ := sg.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwtV4.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwtV4.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwtV4.ValidationError); ok {
			if ve.Errors&jwtV4.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwtV4.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwtV4.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}
