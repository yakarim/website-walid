package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/jinzhu/gorm"

	// database ...
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB ... gorm
var DB *gorm.DB
var once sync.Once

func init() {
	var db *gorm.DB
	portln := os.Getenv("PORT")
	dbport, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	if len(portln) == 0 {
		db = pqsl("localhost", "postgres", "1234", "walid", "disable", 5432)
	} else if !strings.HasPrefix(":", portln) {
		db = pqsl(os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), "require", dbport)
	}
	DB = db.AutoMigrate(&User{}, &Auth{}, &Post{})
}

func pqsl(host, user, pass, database, sslmode string, port int) *gorm.DB {
	var db *gorm.DB

	psql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s ", host, port, user, pass, database, sslmode)
	b, err := gorm.Open("postgres", psql)
	if err != nil {
		log.Println(err)
	}
	db = b

	return db
}
