package middleware

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func CalDigest(){
	 rawData := "20002392302323wdedwe"
	 res := md5.Sum([]byte(rawData))
	 fmt.Println("md5 value is: ",hex.EncodeToString(res[:]))

}

func CalDigest02(){
	rawData := "20002392302323wdedwe"
	hashCal := md5.New()
	hashCal.Write([]byte(rawData))
	res := hashCal.Sum(nil)
	fmt.Println("md5 value is: ",hex.EncodeToString(res))

}

