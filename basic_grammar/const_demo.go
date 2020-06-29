package basic_grammar

const (
	ROUTER_AUTH uint32 = iota + 5 //使用iota才能让后面的数值自动加1
	ROUTER_LOGIN
	ROUTER_SCENE
	ROUTER_CHATLOGIC
	ROUTER_ACCOUNT_AUTH

	ROUTER_MAX = 32
)

type ByteSize float64

const (
	_           = iota             // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
	MB                             // 1 << (10*2)
	GB                             // 1 << (10*3)
	TB                             // 1 << (10*4)
	PB                             // 1 << (10*5)
	EB                             // 1 << (10*6)
	ZB                             // 1 << (10*7)
	YB                             // 1 << (10*8)
)
const (
	Apple, Banana = iota + 1, iota + 2
	Cherimoya, Durian
	Elderberry, Fig
)
const (
	i = iota
	j = 3.14
	k = iota
	l
)

func ConstDemo() {
	//println(ROUTER_ACCOUNT_AUTH)
	//println("the size gb is: ", GB)
	//println(Apple, Banana, Cherimoya, Durian, Elderberry, Fig) //1 2 2 3 3 4
	//print(i, " ,j:", j, " ,k:", k, " ,l:", l)                  //0 ,j:+3.140000e+000 ,k:2 ,l:3
	////println(ROUTER_ACCOUNT_AUTH<<10 | 12)

}
