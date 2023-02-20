package shortenurl

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect(database string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	if err != nil {
		log.Panicln("Database Not initilized!!!!!")
		log.Panic("Error: ", err.Error())
		return nil, err
	}
	fmt.Println("Database Initialized!!!! ")
	return db, nil
}
