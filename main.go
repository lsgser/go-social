package main

import(
	"github.com/joho/godotenv"
	S "github.com/lsgser/go-social/server"
)

func init(){
	godotenv.Load()
}

func main(){
	S.Server()	
}
