package middleware

/*
gorm文档地址：http://gorm.book.jasperxu.com/models.html#md
表名称是结构体名称的复数形式，列名是字段名的下划线形式，可以修改表名
*/

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserModel struct {
	gorm.Model
	Name     string `json:"user_name"`
	Age      int    `json:"age"`
	Location string `json:"location"`
	Job      string `json:"job"`
}

type UserSeq struct {
	ID int64
	//gorm.Model//不需要gorm默认的字段
	AppId      int64
	UserId     int64
	ReadSeq    int64
	ReceiveSeq int64
}

//type DUserLimit struct { //表名是下划线风格怎么办？？？
//	gorm.Model
//	uid string
//	work_limits int32
//	release_wait_time int32
//}
//
//func OrmDemo()  {
//	db, err := gorm.Open("mysql", "root:123@tcp(127.0.0.1:3306)/test?charset=utf8mb4")
//	if err != nil {
//		panic(err)
//	}
//	result := db.First(&DUserLimit{uid:"cyl222222",work_limits:1}) //报错这个：Table 'test.d_user_limits' doesn't exist？？？表名后面怎么加了s了,
//	// 修改表名后报Unknown column 'd_user_limits.deleted_at' in 'where clause' ，各种奇奇怪怪的错误。。。。
//	fmt.Println(result)
//
//}
//func OrmDemo2()  {
//	db, err := gorm.Open("mysql", "root:123@tcp(127.0.0.1:3306)/test?charset=utf8mb4")
//	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
//		return defaultTableName[:len(defaultTableName)-1]
//	}
//	if err != nil {
//		panic(err)
//	}
//	result := db.First(&UserInfo{Name:"tom"})
//	user :=&UserInfo{}
//	row:=result.Row()
//	fmt.Println(reflect.TypeOf(row))
//	id:=0
//	row.Scan(&id,&user.Name,&user.Age,&user.Location,&user.Job)
//	fmt.Println(user)
//	//result := db.First(&UserInfo{Name:"gorm",Age:25,Location:"fujian-china"})
//}
func OrmDemo3()  {
	db, err := gorm.Open("mysql", "root:123@tcp(127.0.0.1:3306)/nonsence?charset=utf8mb4")
	if err != nil {
		fmt.Println("open db error:",err)
		return
	}
	defer db.Close()
	db.SingularTable(true)// 全局禁用表名复数
	//更改默认表名
	//gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
	//	return "prefix_" + defaultTableName;
	//}
	us := &UserSeq{
		AppId: 3,
		UserId: 3,
		ReadSeq: 1,
		ReceiveSeq: 5,
	}
	result := db.Create(us)
	result.Commit()
	fmt.Println("insert result:",us.ID)

}
//
