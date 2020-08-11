package main

import(
	"go/ast"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mssql"
)

func main() {
	r := setupRouter()
	r.Run()
}


func setupRouter() *gin.Engine {
	r := gin.Default()
	h := DBHandler{}
	h.Initialize()
	r.GET("/user",)
	r.GET("/user/:id", )
	r.POST("/user",)
	r.PUT("/user",)
	r.DELETE("/user/:id",)
	return r
}

type DBHandler struct {
	DB *gorm.DB
}

func (h * DBHandler) Initialize() {
	db, err := gorm.Open("mysql", )
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate()
	defer db.Close()
	h.DB = db
}

type User struct {
	id		uint	'gorm:"primary_key" json:"id"'
}


