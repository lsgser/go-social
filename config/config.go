package config

import(
	"golang.org/x/crypto/bcrypt"
)

// Hash a user password
func HashPassword(password string) ([]byte,error){
	return bcrypt.GenerateFromPassword([]byte(password),bcrypt.MinCost)
}


//Check if the password is valid
func CheckPassword(hashedPassword,password string) error{
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err
}