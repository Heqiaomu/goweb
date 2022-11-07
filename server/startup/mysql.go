package startup

import (
	"github.com/Heqiaomu/goweb/tools/mysql"
	"log"
)

func InitMySQL() {
	err := mysql.NewGormDB()
	if err != nil {
		log.Fatal(err)
	}
}
