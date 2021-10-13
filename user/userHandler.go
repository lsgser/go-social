package user

import(
	"github.com/gin-gonic/gin"
	"github.com/badoux/checkmail"
	"strings"
)

/*
	This function will handle data from the add_user endpoint
*/
func AddUser(c *gin.Context){
	user := NewUser()

	if err := c.ShouldBindJSON(&user); err != nil{
		c.JSON(500,gin.H{
			"error":err.Error(),
		})
		return
	}

	//Check if the required strings are not empty
	if strings.TrimSpace(user.Name) == "" && strings.TrimSpace(user.Surname) == "" && strings.TrimSpace(user.Username) == "" && strings.TrimSpace(user.Email) == "" && strings.TrimSpace(user.Password) == "" {
		c.JSON(400,gin.H{
			"error":"Fill in all the require fields",
		})
		return
	} 

	//Make the input strings more presentable before storing it to the database
	user.Name = strings.Title(strings.TrimSpace(strings.ToLower(user.Name)))
	user.Surname = strings.Title(strings.TrimSpace(strings.ToLower(user.Surname)))
	user.Username = strings.TrimSpace(user.Username)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Password = strings.TrimSpace(user.Password)

	//Check if the email format is valid
	if checkmail.ValidateFormat(user.Email) != nil{
		c.JSON(400,gin.H{
			"error":"Invalid email format",
		})
		return	
	}

	/*
		Save the user to the database using the SaveUser method
		that we created in the user.go file
	*/
	err := user.SaveUser()
	if err != nil{
		if err.Error() == "User already exists"{
			c.JSON(400,gin.H{
				"error":err.Error(),
			})
			return
		}

		//Other errors from the SaveUser() method are internal server errors
		c.JSON(500,gin.H{
			"error":err.Error(),
		})
		return
	}

	c.JSON(200,gin.H{
		"message":"User registered successfully",
	})
}