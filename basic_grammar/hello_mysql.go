package basic_grammar

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_Driver = "root:123@tcp(127.0.0.1:3306)/test?charset=utf8mb4"
)

func OpenDB() (success bool, db *sql.DB) {

	var isOpen bool
	db, err := sql.Open("mysql", DB_Driver)
	if err != nil {
		isOpen = false
	} else {
		isOpen = true
	}
	// CheckErr(err)
	return isOpen, db
}

func QueryFromDB(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM d_user_limit")
	// CheckErr(err)
	if err != nil {
		fmt.Println("error:", err)
	} else {
	}
	for rows.Next() {
		var uid string
		var works_limits string
		var release_wait_time string
		// CheckErr(err)
		err = rows.Scan(&uid, &works_limits, &release_wait_time)
		fmt.Println(uid)
		fmt.Println(works_limits)
		fmt.Println(release_wait_time)
	}
}

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func GetDBData() {
	opend, db := OpenDB()
	if opend {
		fmt.Println("open success")
	} else {
		fmt.Println("open faile:")
	}
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},
            {"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`

	json.Unmarshal([]byte(str), &s)

	QueryFromDB(db)
	fmt.Println(s)
}
