package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	name     = "gotodoapp"
	user     = "testuser"
	password = "password"
)

func Database() *gorm.DB {
	credentials := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, name)

	//open a db connection
	db, err := gorm.Open("mysql", credentials)

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
