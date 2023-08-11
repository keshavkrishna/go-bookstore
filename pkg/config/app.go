 package config

 import{
	"github.com/jinzu/gorm"
	_ "github.com/jinzu/gorm/dialects/mysql"
 }

 var{
	db * gorm.DB
 }

 func Connect(){
	d, err := gorm.Open("mysql", "root:1924mysql/sympletest?charset=utf8&parseTime=True&loc=Local")

	if err != nil{
		panic(err)
	}
	db=d
 }

 func GetDB() *gorm.DB{
	return db
 }
