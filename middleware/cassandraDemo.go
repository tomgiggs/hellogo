package middleware

import (
"fmt"
"time"

"github.com/hailocab/gocassa"
)

type Sale struct {
	Id         int64
	CustomerId string
	SellerId   string
	Price      int
	Created    time.Time
}

func CassandraDemo() {

	keySpace, err := gocassa.ConnectToKeySpace("demo", []string{"172.26.43.219"}, "", "")
	if err != nil {
		panic(err)
	}
	salesTable := keySpace.Table("Sale", Sale{}, gocassa.Keys{
		PartitionKeys: []string{"Id"},
	})

	// 创建表
	//err = salesTable.Create()
	//fmt.Println(err)
	// 插入数据
	err = salesTable.Set(Sale{
		Id:        300000,
		CustomerId: "customer-1",
		SellerId:   "seller-1",
		Price:      42,
		Created:    time.Now(),
	}).Run()
	if err != nil {
		panic(err)
	}

	//读取数据
	result := Sale{}
	if err := salesTable.Where(gocassa.Eq("Id", 30000)).ReadOne(&result).Run(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}