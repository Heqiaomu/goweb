package email

import (
	"gitee.com/goweb/config"
	"github.com/gin-gonic/gin"
)

type emailPlugin struct{}

func (e emailPlugin) Register(group *gin.RouterGroup) {
	group.Group("")
}

func (e emailPlugin) RouterPath() string {
	return "email"
}

func CreateEmailPlug(To, From, Host, Secret, Nickname string, Port int, IsSSL bool) *emailPlugin {
	config.GetConfig().Email.To = To
	config.GetConfig().Email.From = From
	config.GetConfig().Email.Host = Host
	config.GetConfig().Email.Secret = Secret
	config.GetConfig().Email.Nickname = Nickname
	config.GetConfig().Email.Port = Port
	config.GetConfig().Email.IsSSL = IsSSL

	return &emailPlugin{}

}
