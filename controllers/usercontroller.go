package controllers

import (
  "net/http"
  "strconv"

  "MVC/models"
  "github.com/labstack/echo"
)

// get all users
func GetUsersController(c echo.Context) error {
  users, err := models.GetUsers()
  if err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
  return c.JSON(http.StatusOK, users)
}


// get single user by id
func GetUserController(c echo.Context) error {
  id, _ := strconv.Atoi(c.Param("id"))
  user, err := models.GetUser(id)

  if err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
  return c.JSON(http.StatusOK, user)
}

// create new user
func CreateUserController(c echo.Context) error {
  var user models.User
  c.Bind(&user)

  result, err := models.CreateUser(&user)
  if err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
  return c.JSON(http.StatusCreated, result)
}

//remove user by id
func DeleteUserController(c echo.Context) error {
  // your code here
  id, _ := strconv.Atoi(c.Param("id"))

 _, err := models.DeleteUser(id)
  if err != nil {
	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
  return c.JSON(http.StatusOK, map[string]interface{}{
	"message": "success delete user",
  })
 
}

func UpdateUserController(c echo.Context) error {
  // your code here
  id, _ := strconv.Atoi(c.Param("id"))
  var user models.User
  c.Bind(&user)

  _, err := models.UpdateUser(id,user)
   if err != nil {
	 return echo.NewHTTPError(http.StatusBadRequest, err.Error())
   }
   return c.JSON(http.StatusOK, map[string]interface{}{
	 "message": "success update user",
	 "user": user,

   })
  
}
