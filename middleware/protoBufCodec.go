package middleware

import (
	"bytes"
	"github.com/funny/link"
	"github.com/golang/protobuf/proto"
	//"google.golang.org/protobuf/proto"
	"io"
	"reflect"
)

type ProtoBufProtocol struct {
	types map[string]reflect.Type
	names map[reflect.Type]string
}

func ProtoBuf() *ProtoBufProtocol {
	return &ProtoBufProtocol{
		types: make(map[string]reflect.Type),
		names: make(map[reflect.Type]string),
	}
}

func (j *ProtoBufProtocol) Register(t interface{}) {
	rt := reflect.TypeOf(t)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	name := rt.PkgPath() + "/" + rt.Name()
	j.types[name] = rt
	j.names[rt] = name
}

func (j *ProtoBufProtocol) RegisterName(name string, t interface{}) {
	rt := reflect.TypeOf(t)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	j.types[name] = rt
	j.names[rt] = name
}


type protoBufIn struct {
	Code string
	Body proto.Message
}

type protoBufOut struct {
	Code string
	Body proto.Message
}
type ProtoBufInterFace interface {

}

type protoBufCodec struct {
	p       *ProtoBufProtocol
	closer  io.Closer
	writer io.Writer
	reader io.Reader
}

func (c *protoBufCodec) Receive() (interface{}, error) {
	var in protoBufIn
	//rawIn := make([]byte,0) //这样是读不出来的
	rawIn := make([]byte, 4096)
	//rawIn := make([]byte,4096)
	//var rawIn []byte
	n,err := c.reader.Read(rawIn)
	var loc int
	for idx,c := range rawIn{
		if loc==0 &&string(c)=="|"{
			loc = idx
		}
	}
	in.Code = string(rawIn[:loc])
	if err != nil {
		return nil, err
	}
	//var body pb.HelloRequest
	var body proto.Message
	if in.Code != "" {
		if t, exists := c.p.types[in.Code]; exists {
			//body = reflect.New(t).Elem().Interface().(proto.Message)//panic: interface conversion: proto.HelloRequest is not protoiface.MessageV1: missing method ProtoMessage
			body = reflect.New(t).Interface().(proto.Message)//panic: interface conversion: proto.HelloRequest is not protoiface.MessageV1: missing method ProtoMessage
			//body = reflect.Indirect(reflect.New(t.Elem())).Addr().Interface().(proto.Message)
		}
		//fmt.Println(string(rawIn[loc+1:]))
		err = proto.Unmarshal(rawIn[loc+1:n],body)//proto: invalid field number
		if err != nil {
			return nil, err
		}
		return body, nil
	}else {
		return nil,nil
	}

}

func (c *protoBufCodec) Send(msg interface{}) (err error) {
	var out protoBufOut
	t := reflect.TypeOf(msg)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if name, exists := c.p.names[t]; exists {
		out.Code = name
	}
	outByte,_:=proto.Marshal(msg.(proto.Message))
	//
	//_,err =c.writer.Write([]byte(out.Code+"|"))
	var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲
	buffer.Write([]byte(out.Code+"|"))
	buffer.Write(outByte)
	b3 :=buffer.Bytes()
	_,err =c.writer.Write(b3)
	return err
}

func (c *protoBufCodec) Close() error {
	if c.closer != nil {
		return c.closer.Close()
	}
	return nil
}

func (j *ProtoBufProtocol) NewCodec(rw io.ReadWriter) (link.Codec, error) {

	codec := &protoBufCodec{
		p:       j,
		writer: rw,
		reader: rw,
	}
	codec.closer, _ = rw.(io.Closer)
	return codec, nil
}
