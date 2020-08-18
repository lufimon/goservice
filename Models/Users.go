package Models

import (
	"../Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllUsers(u *[]User) (err error) {
	if err = Config.DB.Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetUserById(u *User, id uint) (err error) {
	if err := Config.DB.Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

func AddUser(u *User) (err error) {
	if err := Config.DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(u *User) (err error) {
	fmt.Println(u)
	Config.DB.Save(u)
	return nil
}

func DeleteUser(u *User, id uint) (err error) {
	if err := Config.DB.Unscoped().Delete(u, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
