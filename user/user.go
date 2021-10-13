package user

import(
	CO "github.com/lsgser/go-social/config"
	"errors"
	"database/sql"
)

type User struct{
	ID int64 `json:"-"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Username string `json:"username"`
	Email string `json:"email"`
	EmailVerifiedAt string `json:"email_verified_at,omitempty"`
	Password string `json:"password,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

func NewUser() *User{
	return new(User)
}

/*
	Saves a new user to the database
*/
func (u *User) SaveUser() error{
	db,err := CO.GetDB()

	if err != nil{
		err = errors.New("DB connection error")
		return err
	}

	defer db.Close()

	/*
		Check if the user already exists
	*/
	var (
		username string 
		email string
	) 

	stmt,err := db.Prepare("SELECT username,email FROM users WHERE username = ? OR email = ?")

	if err != nil{
		return err
	}

	defer stmt.Close()

	err = stmt.QueryRow(u.Username,u.Email).Scan(&username,&email)

	if err != nil{
		//The user does not exist
		if err == sql.ErrNoRows{
			//Hash the users password
			hashedPassword,err := CO.HashPassword(u.Password)
			if err != nil {
				return err
			}
			//Add the new user
			insert_stmt,err := db.Prepare("INSERT users (name,surname,username,email,password) VALUES (?,?,?,?,?)")

			if err != nil{
				return err
			}

			defer insert_stmt.Close()

			_,err = insert_stmt.Exec(u.Name,u.Surname,u.Username,u.Email,hashedPassword)
			
			return err
		}else{
			return err
		}
	}else{
		err = errors.New("User already exists")
		return err
	}

	return err
}