package server

import(
	"os"
	"github.com/gin-gonic/gin"
	"github.com/lsgser/go-social/hello"
	"github.com/lsgser/go-social/user"
	"fmt"
)

/*
	Runs our server and houses all of our route endpoints
*/
func Server(){
	if(os.Getenv("BUILD") == "PRODUCTION"){
		gin.SetMode(gin.ReleaseMode)
		fmt.Println("BUILD: "+os.Getenv("BUILD"))
		fmt.Println("PORT: "+os.Getenv("PORT"))
	}

	router := gin.Default()
	hello.Routes(router)
	user.Routes(router)
	router.Run(":"+os.Getenv("PORT"))
}