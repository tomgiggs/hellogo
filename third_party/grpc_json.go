package third_party

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/johanbrandhorst/grpc-json-example/codec"
	_ "github.com/johanbrandhorst/grpc-json-example/codec"
	"google.golang.org/grpc"
	"time"
)

func StartGrpcJson()  {
	_,err := grpc.Dial("localhost:9000",
		grpc.WithDefaultCallOptions(grpc.CallContentSubtype(codec.JSON{}.Name())),
	)
	if err != nil {
		fmt.Println("StartGrpcJson",err)
	}
}


/*
Echo -en '\x00\x00\x00\x00\x17{"id":1,"role":"ADMIN"}' | curl -ss -k --http2 \
        -H "Content-Type: application/grpc+json" \
        -H "TE:trailers" \
        --data-binary @- \
        https://localhost:10000/example.UserService/AddUser | od -bc
*/

func RequestWithTimeout(client *resty.Client, params interface{}, url string, timeout int) (*resty.Response, error) {
	var (
		resp *resty.Response
		err  error
	)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(timeout))
	defer cancel()
	go func(ctx context.Context) {
		resp, err = client.R().SetBody(params).Post(url)
		if err != nil {
			fmt.Println("err",err)
		}
	}(ctx)

	for {
		select {
		case <-ctx.Done():
			return resp, nil
		case <-time.After(time.Second * time.Duration(timeout)):
			return resp, fmt.Errorf("timeout")
		}
	}

}
