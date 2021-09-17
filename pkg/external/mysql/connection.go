package mysql

import (
	"log"

	"github.com/awesome-linus/go-graphql-echo-boilerplate/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", config.GetDsn())

	if err != nil {
		log.Fatal(err, "Unable to connect to MySQL server.")
	}

	// DBエンジンを「InnoDB」に設定
	db.Set("gorm:table_options", "ENGINE=InnoDB")

	// 詳細なログを表示
	db.LogMode(true)

	return db
}

func Close() {
	defer func() {
		err := db.Close()
		if err == nil {
			return
		}
		log.Fatalln(err)
	}()
}
