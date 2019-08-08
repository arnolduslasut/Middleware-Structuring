// package main

// import (
// 	"MVC/controllers"
// 	"net/http"

// 	"github.com/labstack/echo"
// )

// func main() {

// }


package main

import (
  "MVC/config"
  "MVC/models"
  "MVC/routes"

  m "MVC/middlewares"
)

func main() {
  e := routes.New()

  // implement middleware
  m.LogMiddlewares(e)
  // migration database
  // InitialMigration()

  e.Start(":8000")
}

func InitialMigration() {
  var db = config.DB
  if !db.HasTable(&models.User{}) { // check database exist or not
    db.AutoMigrate(&models.User{})
  }
}
