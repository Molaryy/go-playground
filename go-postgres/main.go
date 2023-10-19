package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "post-go/controllers"
)

func router(r *gin.Engine)  {
   r.GET("/", controllers.HelloWorld)
}

func main() {
  r := gin.Default()

  router(r)
  err := r.Run()

  if err != nil {
    fmt.Println("Couldn't Run the app!, try again or verify your connection")
    return
  }
}

