package start

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"fmt"
	gormModels "lesson/gorm/start/models"
)

func Test1() {
	db, err := gorm.Open("mysql", "root:123456@/bbs_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println(db.HasTable("novel_list"))
	listOne := new(gormModels.NovelList)
	db.Model(&gormModels.NovelList{}).Where("id=?","1").First(listOne);
	fmt.Println(listOne)
}
