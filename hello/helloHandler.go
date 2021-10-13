package hello

import(
	"github.com/gin-gonic/gin"
)

func ShowIndex(c *gin.Context){
	c.JSON(200,gin.H{
		"message":"Hello World",
	})
}