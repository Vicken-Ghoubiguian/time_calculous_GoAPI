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
    c.JSON(http.StatusOK, gin.H{"title": "", "presentation": ""})
  })

  //
  r.GET("/wished_wday_in_choosen_month/:wday/:month", func(c *gin.Context) {

    //
    c.JSON(http.StatusOK, gin.H{"data": "TODO"})
  })

  //
  r.GET("/number_of_weeks_in_year/:year", func(c *gin.Context) {

    //
    year := c.Param("year")
    number_of_weeks := 13
 
    //
    c.JSON(http.StatusOK, gin.H{"year": year, "number_of_weeks": number_of_weeks})
  })

  //
  r.Run()
}
