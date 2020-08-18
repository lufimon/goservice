package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

func main() {
	r := setupRouter()
	r.Run()
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	h := DBHandler{}
	h.Initialize()
	r.GET("/user", h.GetAllUser)
	r.GET("/user/:id", h.GetUser)
	r.POST("/user", h.SaveUser)
	r.PUT("/user", h.UpdateUser)
	r.DELETE("/user/:id", h.DeleteUser)
	return r
}

type DBHandler struct {
	DB *gorm.DB
}

func (h *DBHandler) Initialize() {
	db, err := gorm.Open("mysql", "root:root@/training?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.AutoMigrate(&User{})
	h.DB = db
}

type User struct {
	gorm.Model
	Id        uint    `gorm:"primary_key;AUTO_INCREMENT" json:"userId"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	BirthDay  string  `json:"birthDay"`
	Age       int     `json:"age"`
	Contact   Contact `gorm:"-" json:"contact"`
	Facebook  string
	LineId    string
	Instagram string
}

type Contact struct {
	Facebook  string `json:"facebook"`
	LineId    string `json:"lineId"`
	Instagram string `json:"instagram"`
}

func (h *DBHandler) GetAllUser(c *gin.Context) {
	var users []User
	h.DB.Find(&users)
	for _, user := range users {
		contact := Contact{}
		contact.Facebook = user.Facebook
		contact.Instagram = user.Instagram
		contact.LineId = user.LineId
		user.Contact = contact
	}
	c.JSON(http.StatusOK, users)
}

func (h *DBHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user := User{}
	contact := Contact{}
	if err := h.DB.Find(&user, id).Error; err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	contact.Facebook = user.Facebook
	contact.Instagram = user.Instagram
	contact.LineId = user.LineId
	user.Contact = contact
	c.JSON(http.StatusOK, user)
}

func (h *DBHandler) SaveUser(c *gin.Context) {
	user := User{}
	user.Facebook = user.Contact.Facebook
	user.LineId = user.Contact.LineId
	user.Instagram = user.Contact.Instagram
	if err := c.ShouldBindJSON(&user); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	if err := h.DB.Save(&user).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *DBHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	user := User{}
	user.Facebook = user.Contact.Facebook
	user.LineId = user.Contact.LineId
	user.Instagram = user.Contact.Instagram
	if err := h.DB.Find(&user, id).Error; err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	if err := h.DB.Save(&user).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *DBHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	user := User{}

	if err := h.DB.Find(&user, id).Error; err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	if err := h.DB.Delete(&user).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)
}
