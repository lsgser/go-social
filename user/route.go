package user

import(
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine){
	api := route.Group("/api")
	{
		api.POST("add_user",AddUser)
		api.POST("sign_in",LoginUser)
		api.GET("auth/:token",CheckUser)
	}	
}