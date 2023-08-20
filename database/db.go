package database

import (
	"assigment_project_rest_api/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func ConnectDB() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("DB_PORT")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	name := os.Getenv("DB_NAME")

	// config := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, name)
	db, err = gorm.Open(mysql.Open(config), &gorm.Config{})
	if err != nil {
		fmt.Println("can't connect to database")
		panic(err)
	}
	fmt.Println("connected to database")

	err = db.AutoMigrate(models.Student{}, models.Score{})
	if err != nil {
		fmt.Println(err)
		return
	}

}

func GetDB() *gorm.DB {
	return db
}
