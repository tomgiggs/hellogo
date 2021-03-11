//go:binary-only-package
package basic_grammar



func Add(a ,b int32) int32{
	return a+b
}
func Sub(a,b int32) int32{
	return a -b
}

//go build -o c_bin.a -i
