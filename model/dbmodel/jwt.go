package dbmodel

type JwtBlacklist struct {
	Model
	Jwt string `gorm:"type:text;comment:jwt"`
}
