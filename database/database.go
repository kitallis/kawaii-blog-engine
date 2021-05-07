package database

import (
	"fmt"
	"github.com/nid90/kawaii-blog-engine/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func InitDatabase() {
	var err error

	DBConn, err = gorm.Open(sqlite.Open(config.Config("DATABASE_NAME_WITH_EXT")), &gorm.Config{})

	if err != nil {
		panic("Failed to connect Database")
	}

	fmt.Println("👍🏽 Connection opened to database")
}
