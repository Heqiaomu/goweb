package startup

import (
	"gitee.com/goweb/tools/mysql"
	"log"
)

func InitMySQL() {
	err := mysql.NewGormDB()
	if err != nil {
		log.Fatal(err)
	}
}
