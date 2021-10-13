package hello

import(
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine){
	route.GET("/",ShowIndex)
	api := route.Group("/api")
	{
		api.GET("/hello",ShowIndex)
	}
}