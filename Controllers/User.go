package Controllers

import (
	"../Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//GetAllUsers godoc
//@Summary Get detail user of all
//@Description Get detail user of all
//@Accept json
//@Produce json
//@Success 200 {array} Models.Users
//@Success 204
//@Failure 400
//@Failure 404
//@Failure 500
//@Router /user [get]
func GetAllUsers(c *gin.Context) {
	var users []Models.User
	var response []Models.Users
	err := Models.GetAllUsers(&users)
	for _, user := range users {
		u := Models.Users{}
		u.ID = user.Model.ID
		u.FirstName = user.FirstName
		u.LastName = user.LastName
		u.BirthDay = user.BirthDay
		u.Age = user.Age
		contact := Models.Contact{}
		contact.Facebook = user.Facebook
		contact.LineId = user.LineId
		contact.Instagram = user.Instagram
		u.Contact = contact
		response = append(response, u)
	}
	if err != nil {
		if err.Error() == "record not found" {
			fmt.Println("RECORD NOT FOUND :", err)
			c.Status(http.StatusNoContent)
		} else {
			fmt.Println("ERROR :", err)
			c.Status(http.StatusInternalServerError)
		}
	} else {
		c.JSON(http.StatusOK, response)
	}
}

//GetUserById godoc
//@Summary Get detail user by id
//@Description Get detail user by id
//@Accept json
//@Produce json
//Param id path int true "User ID"
//@Success 200 {array} Models.Users
//@Success 204
//@Failure 400
//@Failure 404
//@Failure 500
//@Router /user/{id} [get]
func GetUserById(c *gin.Context) {
	id, e := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if e == nil {
		fmt.Println("ID : ", id)
	}
	var user Models.User
	err := Models.GetUserById(&user, uint(id))
	response := Models.Users{}
	response.ID = user.Model.ID
	response.FirstName = user.FirstName
	response.LastName = user.LastName
	response.BirthDay = user.BirthDay
	response.Age = user.Age
	contact := Models.Contact{}
	contact.Facebook = user.Facebook
	contact.LineId = user.LineId
	contact.Instagram = user.Instagram
	response.Contact = contact
	if err != nil {
		if err.Error() == "record not found" {
			fmt.Println("RECORD NOT FOUND :", err)
			c.Status(http.StatusNoContent)
		} else {
			fmt.Println("ERROR :", err)
			c.Status(http.StatusInternalServerError)
		}
	} else {
		c.JSON(http.StatusOK, response)
	}
}

func AddUser(c *gin.Context) {
	var user Models.Users
	c.BindJSON(&user)
	var u Models.User
	u.ID = user.ID
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.BirthDay = user.BirthDay
	u.Age = user.Age
	u.Facebook = user.Contact.Facebook
	u.LineId = user.Contact.LineId
	u.Instagram = user.Contact.Instagram
	err := Models.AddUser(&u)
	if err != nil {
		fmt.Println("ERROR :", err)
		c.Status(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	}
}

func UpdateUser(c *gin.Context) {
	var u Models.Users
	c.BindJSON(&u)
	var user Models.User
	err := Models.GetUserById(&user, u.ID)
	if err != nil {
		fmt.Println("ERROR :", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	user.ID = u.ID
	user.FirstName = u.FirstName
	user.LastName = u.LastName
	user.BirthDay = u.BirthDay
	user.Age = u.Age
	user.Facebook = u.Contact.Facebook
	user.LineId = u.Contact.LineId
	user.Instagram = u.Contact.Instagram
	err = Models.UpdateUser(&user)
	if err != nil {
		fmt.Println("ERROR :", err)
		c.Status(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	}
}

func DeleteUser(c *gin.Context) {
	id, e := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if e == nil {
		fmt.Println("ID : ", id)
	}
	var user Models.User
	err := Models.DeleteUser(&user, uint(id))
	if err != nil {
		fmt.Println("ERROR :", err)
		c.Status(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	}
}
