package models

import (
  "MVC/config"
  "github.com/jinzhu/gorm"
)

var db = config.DB

type User struct {
  gorm.Model
  Name     string `json:"name" form:"name" query:"name"`
  Email    string `json:"email" form:"email" query:"email"`
  Password string `json:"password" form:"password" query:"password"`
}

// get all user
func GetUsers() (interface{}, error) {
  var users []User
  if err := db.Find(&users).Error; err != nil {
    return nil, err
  }
  return users, nil
}

// get single user by id
func GetUser(id int) (interface{}, error) {
  var user User
  if err := db.First(&user, id).Error; err != nil {
    return nil, err
  }
  return user, nil
}

// create new user
func CreateUser(user *User) (interface{}, error) {
  if err := db.Create(&user).Error; err != nil {
    return nil, err
  }
  return user, nil
}

// remove user by id
func DeleteUser(id int) (interface{}, error) {
  // your code here
  var user User
  if err := db.First(&user, id).Error; err != nil {
	return nil,err
  }
  if err := db.Delete(&user).Error; err != nil {
	return nil,err
  }

  return nil,nil
}

// update user by id
func UpdateUser(id int, newUser User) (interface{}, error) {
  // your code here
  var user User

  if err := db.First(&user, id).Error; err != nil {
	return nil,err
  }
  user.Name = newUser.Name
  user.Email = newUser.Email
  user.Password = newUser.Password
  if err := db.Save(&user).Error; err != nil {
	return nil,err
  }

  return user,nil
}
