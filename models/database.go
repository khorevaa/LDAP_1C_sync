package models

import (
	// ORM framework
	"strings"

	"github.com/jinzhu/gorm"
	//HTTP framework
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB ..
var DB *gorm.DB

//ConnectDB ...
func ConnectDB(rootConfigs string) {
	configs := strings.Split(rootConfigs, "/")
	dialect := configs[0]
	params := configs[1]
	db, err := gorm.Open(dialect, params)
	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}
	db.AutoMigrate(&Company{}, &Department{}, &Position{}, &Employee{})
	DB = db
	DB.LogMode(true)
}
