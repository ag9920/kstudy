package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
)

func main() {
	ctx := context.Background()
	// 本地文件 idl 解析
	// YOUR_IDL_PATH thrift 文件路径: 举例 ./idl/example.thrift
	// includeDirs: 指定 include 路径，默认用当前文件的相对路径寻找 include
	p, err := generic.NewThriftFileProvider("../example.thrift")
	if err != nil {
		panic(err)
	}
	// 构造 map 请求和返回类型的泛化调用
	g, err := generic.MapThriftGeneric(p)
	if err != nil {
		panic(err)
	}
	cli, err := genericclient.NewClient("destServiceName", g, client.WithHostPorts("127.0.0.1:8888"))
	if err != nil {
		panic(err)
	}
	// 'ExampleMethod' 方法名必须包含在 idl 定义中
	resp, err := cli.GenericCall(ctx, "ExampleMethod", map[string]interface{}{
		"Msg": "hello",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	// resp is a map[string]interface{}
}
