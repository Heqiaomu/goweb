package startup

import (
	"github.com/sun-iot/goweb/tools/mysql"
	"log"
)

func InitMySQL() {
	err := mysql.NewGormDB()
	if err != nil {
		log.Fatal(err)
	}
}
