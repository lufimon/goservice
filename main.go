package main

import (
	"./Config"
	"./Models"
	"./Routers"
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

var err error

func main() {
	os.Setenv("DATABASE_URL", "root:root@(localhost:3306)/training?charset=utf8&parseTime=True&loc=Local")
	os.Setenv("IP", "10.100.150.155")
	os.Setenv("PORT", "8089")

	Config.DB, err = gorm.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("statuse: ", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})
	r := Routers.SetupRouter()
	r.Run(os.Getenv("IP") + ":" + os.Getenv("PORT"))
}
