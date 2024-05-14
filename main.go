//
package main

//
import (

  //
  "net/http"

  //
  "github.com/gin-gonic/gin"
)

//
func main() {

  //
  r := gin.Default()

  //
  r.GET("/", func(c *gin.Context) {

    //
    c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
  })

  //
  r.GET("/wished_wday_in_choosen_month", func(c *gin.Context) {

    //
    c.JSON(http.StatusOK, gin.H{"data": "TODO"})
  })

  //
  r.GET("/number_of_weeks_in_a_year_according_to_the_iso_norm", func(c *gin.Context) {
 
    //
    c.JSON(http.StatusOK, gin.H{"data": "TODO"})
  })

  //
  r.Run()
}
