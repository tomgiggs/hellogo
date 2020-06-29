package middleware

//import (
//	"database/sql"
//	"fmt"
//	//_ "github.com/go-sql-driver/mysql"
//)
//
//const (
//	DB_Driver = "root:123@tcp(127.0.0.1:3306)/test?charset=utf8mb4"
//)
//
//func OpenDB() (success bool, db *sql.DB) {
//
//	var isOpen bool
//	db, err := sql.Open("mysql", DB_Driver)
//	if err != nil {
//		isOpen = false
//	} else {
//		isOpen = true
//	}
//	// CheckErr(err)
//	return isOpen, db
//}
//
//func QueryFromDB(db *sql.DB) interface{} {
//	rows, err := db.Query("SELECT * FROM user_info")
//	if err != nil {
//		fmt.Println("error:", err)
//		return nil
//	}
//	//var userDatas = make([]interface{},20,20)
//	userDatas := []interface{}{}
//	for rows.Next() {
//		var user = UserInfo{}
//		var id = 0
//		err = rows.Scan(&id, &user.Name, &user.Age, &user.Location, &user.Job)
//		//fmt.Println(user)
//		//userData,_:= json.Marshal(user)
//		//fmt.Println(string(userData))
//		userDatas = append(userDatas, user)
//	}
//	//userDataJson,_:= json.Marshal(userDatas)
//	//fmt.Println(string(userDataJson))
//	return userDatas
//
//}
//
//type Server struct {
//	ServerName string
//	ServerIP   string
//}
//
//func GetDBData() {
//	opend, db := OpenDB()
//	if opend {
//		fmt.Println("open success")
//	} else {
//		fmt.Println("open faile:")
//	}
//	QueryFromDB(db)
//}
