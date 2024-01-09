package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	// dsn := os.Getenv("DB_URL")
	dsn := "host=castor.db.elephantsql.com user=kyyxbwlw password=WahI74qtkYwSIgVljyCQG1NUOjDe94d3 dbname=kyyxbwlw port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Fail connect to database")
	}
}
