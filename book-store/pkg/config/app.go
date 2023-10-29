package config
import(
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db * gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql", "root:cm9vdDEyMw0K@tcp(localhost:3307)/mysql")
	if err != nil{
			panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}
