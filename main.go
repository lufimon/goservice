package main

import (
	"./Config"
	"./Models"
	"./Routers"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/pdrum/swagger-automation/api"
	_ "github.com/pdrum/swagger-automation/docs"
	"os"
)

var err error

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email tanuphong.p@hotmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8089
// @BasePath /api/
// @query.collection.format multi
func main() {
	os.Setenv("DATABASE_URL", "root:root@(localhost:3306)/training?charset=utf8&parseTime=True&loc=Local")
	//os.Setenv("IP", "10.100.150.155")
	os.Setenv("IP", "0.0.0.0")
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
