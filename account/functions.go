package account

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatbase() *gorm.DB {
	name := os.Getenv("database_name")
	pass := os.Getenv("database_pass")
	user := os.Getenv("database_user")
	host := os.Getenv("database_host")
	ssl := os.Getenv("databse_ssl")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host, port, user, name, pass, ssl)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("err")
		return
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Println(err)
	}

	return db
}
