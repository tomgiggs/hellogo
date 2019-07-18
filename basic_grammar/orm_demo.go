package basic_grammar

import (
	"fmt"
	"github.com/jinzhu/gorm"
_ "github.com/go-sql-driver/mysql"
)

type DUserLimit struct { //表名是下划线风格怎么办？？？
	gorm.Model
	uid string
	work_limits int32
	release_wait_time int32
}

func OrmDemo()  {
	db, err := gorm.Open("mysql", "root:123@tcp(127.0.0.1:3306)/test?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	result := db.First(&DUserLimit{uid:"cyl222222",work_limits:1}) //报错这个：Table 'test.d_user_limits' doesn't exist？？？表名后面怎么加了s了,修改表名后报Unknown column 'd_user_limits.deleted_at' in 'where clause' ，各种奇奇怪怪的错误。。。。
	fmt.Println(result)

}





