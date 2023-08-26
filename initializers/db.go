package initializers

import (
	"fmt"
	"os"

	"github.com/prezeswastaken/blog/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDoDatabase() {
	DbUrl := os.Getenv("DB_URL")
	var err error
	dsn := DbUrl
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to the database :(")
	} else {
		fmt.Println("Connected to the database!")
	}
	fmt.Println(DB)
}

func SyncDb() {
	DB.AutoMigrate(&models.Post{})
}
