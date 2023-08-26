package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDoDatabase() {
	DbUrl := os.Getenv("DB_URL")
	var err error
	dsn := DbUrl
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to database :(")
	} else {
		fmt.Println("Connected? I guess...")
	}
	println(DB)
}
