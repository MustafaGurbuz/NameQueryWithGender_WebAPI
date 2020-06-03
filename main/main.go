package main
import (
	"./controllers"
	"fmt"
	"github.com/gin-gonic/gin"
)
func Query(c *gin.Context) {
	name := c.Query("name")
	c.JSON(200,gin.H{
		"name":name,
		"probability":controllers.PROBABILITY(name),
		"gender":controllers.GENDER(name),
		"count":controllers.COUNT(name),
	})
}
func main(){
	fmt.Println("Hello world...")
	r := gin.Default()
	r.GET("/query",Query)
	r.Run()
}