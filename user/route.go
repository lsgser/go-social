package user

import(
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine){
	api := route.Group("/api")
	{
		api.POST("add_user",AddUser)
	}
}