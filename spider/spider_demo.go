package main //如果包名是spider的话就会报错：GoLand runnerw.exe: CreateProcess failed with error 216

import (
	"fmt"
	"github.com/beevik/etree" //兼容性有点差，对于没有正常关闭的标签就报错
	"io/ioutil"
	"net/http"
)

func main(){
	var url string = "https://blog.csdn.net/qqqqll3/article/details/85054126"
	resp,err :=http.Get(url)
	if err!=nil{
		panic(err) //XML syntax error on line 122: unquoted or missing attribute value in element
	}
	body,_ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	doc := etree.NewDocument()
	if err := doc.ReadFromString(string(body)); err != nil {
		fmt.Println(err)
		//log.Panic(err)
	}
	//  过去数据方法和xpath规则一致
	for _, t := range doc.FindElements("//div[@id='content_views']//p") {
		fmt.Println("Title:", t.Text())
	}
	//node, err := xmlpath.Parse(body)
	//if err != nil {
	//	panic("xmlpath parse file failed!!!")
	//}
	//fmt.Println(node)

	//if doc, err := libxml2.ParseHTMLReader(body); err != nil {
	//	log.Fatal(err)
	//
	//} else {
	//	defer doc.Free()
	//	nodes, err := doc.Find("//div[@class='row download']//img/@src")
	//	fmt.Println(nodes.NodeList()[0].TextContent(), err)
	//}



	}



