package main

//
import (

	//
	"net/http"

	//
	"github.com/gin-gonic/gin"
)

func main() {

	//
	r := gin.Default()

	//
	r.StaticFile("/favicon.ico", "./resources/favicon.ico")

	// Definition of the 'root' URL...
	r.GET("/", func(c *gin.Context) {

		//
		c.JSON(http.StatusOK, gin.H{"title": "time_calculous_GoAPI", "presentation": "API written in Go to use the 'time_calculous' library functions...", "description": "", "links": gin.H{"github": "https://github.com/Vicken-Ghoubiguian/time_calculous_GoAPI", "dockerhub": ""}})
	})

	// Definition of the 'wished_wday_in_choosen_month' to use the 'wished_wday_in_choosen_month' function...
	r.GET("/wished_wday_in_choosen_month/:wday/:month", func(c *gin.Context) {

		//
		/*wday := c.Param("wday")
		  month := c.Param("month")
		*/

		//
		c.JSON(http.StatusOK, gin.H{"data": "TODO"})
	})

	// Definition of the 'number_of_weeks_in_year' to use the 'number_of_weeks_in_a_year_according_to_the_iso_norm' function...
	r.GET("/number_of_weeks_in_year/:year", func(c *gin.Context) {

		//
		year := c.Param("year")
		number_of_weeks := 13

		//
		c.JSON(http.StatusOK, gin.H{"year": year, "number_of_weeks": number_of_weeks})
	})

	// Definition of the 'wished_wday_in_choosen_year' to use the 'wished_wday_in_choosen_year' function...
	r.GET("/wished_wday_in_choosen_year/:year/:wday/:number_of_weekday", func(c *gin.Context) {

		//
		/*year := c.Param("year")
		  wday := c.Param("wday")
		  number_of_weekday = c.Param("number_of_weekday")
		*/

		//
		c.JSON(http.StatusOK, gin.H{"data": "TODO"})
	})

	// Definition of the 'wished_number_in_year_is_day_in_choosen_year' to use the 'wished_number_in_year_is_day_in_choosen_year' function...
	r.GET("/wished_number_in_year_is_day_in_choosen_year/:mday/:month/:year", func(c *gin.Context) {

		//
		/*mday := c.Param("mday")
		  wday := c.Param("wday")
		  year := c.Param("year")
		*/

		//
		c.JSON(http.StatusOK, gin.H{"data": "TODO"})
	})

	// Definition of the 'calculations_on_date_and_time_from_today' to use the 'calculations_on_date_and_time_from_today' function...
	r.GET("/calculations_on_date_and_time_from_today/:milleniums/:centuries/:decades/:years/:months/:weeks/:days/:hours/:minutes/:seconds", func(c *gin.Context) {

		//
		c.JSON(http.StatusOK, gin.H{"data": "TODO"})
	})

	// Definition of the 'number_of_days_in_choosen_month_in_choosen_year' to use the 'number_of_days_in_choosen_month_in_choosen_year' function...
	r.GET("/number_of_days_in_choosen_month_in_choosen_year/:month/:year", func(c *gin.Context) {

		//
		/*year := c.Param("year")
		  month := c.Param("month")
		*/

		//
		c.JSON(http.StatusOK, gin.H{"data": "TODO"})
	})

	//
	r.Run()
}
