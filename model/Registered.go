package model
import (
	"g7love/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "fmt"
	_ "time"
	"time"
	"fmt"
)
type Registered struct {
	gorm.Model
	Email		string	`gorm:"size:255"`
	School		int `sql:"type:int(4);NOT NULL;COMMENT:'学校'"`
	Name         string  `gorm:"size:255"`       // string默认长度为255, 使用这种tag重设。
}

type School struct {
	Id			int
	Name		string
	Parentid	int
	Createtime	time.Time
	Updatetime	time.Time
	Deleted 	int
}

func Creat(id string) *gorm.DB {
	fmt.Println(id)
	db := database.GetDB()
	ac := db.Find(&School{})
	return ac
}