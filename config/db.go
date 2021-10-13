package config
import(
	"os"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetDB() (*sql.DB,error){
	//Name of the database driver
	DBDriver := os.Getenv("DB_DRIVER")  
	//Database name
	DBName := os.Getenv("DB_NAME")
	//User
	DBUser := os.Getenv("DB_USER")
	//Password 
	DBPassword := os.Getenv("DB_PASSWORD")
	//DBURL is the URL of the database
	DBURL := DBUser + ":" + DBPassword + "@/" + DBName

	db,err := sql.Open(DBDriver,DBURL)
	if err != nil{
		return db,err
	}
	
	return db,err
}