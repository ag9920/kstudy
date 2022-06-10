package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
)

type GenericCallConfig struct {
	IDLPath         string
	DestServiceName string
	Method          string
}

func main() {
	ctx := context.Background()
	conf := loadGenericCallConfig()
	// 本地文件idl解析
	// YOUR_IDL_PATH thrift文件路径: 举例 ./idl/example.thrift
	// includeDirs: 指定include路径，默认用当前文件的相对路径寻找include
	p, err := generic.NewThriftFileProvider(conf.IDLPath)
	if err != nil {
		panic(err)
	}
	// 构造http类型的泛化调用
	g, err := generic.HTTPThriftGeneric(p)
	if err != nil {
		panic(err)
	}
	cli, err := genericclient.NewClient(conf.DestServiceName, g, client.WithHostPorts("127.0.0.1:8888"))
	if err != nil {
		panic(err)
	}

	req := fetchReq()

	// customReq *generic.HttpRequest
	// 由于http泛化的method是通过bam规则从http request中获取的，所以填空就行
	resp, err := cli.GenericCall(ctx, conf.Method, req)
	if err != nil {
		panic(err)
	}
	realResp := resp.(*generic.HTTPResponse)
	fmt.Printf("%+#v", realResp)
	//realResp.Write(w) // 写回ResponseWriter，用于http网关
}

func loadGenericCallConfig() GenericCallConfig {
	return GenericCallConfig{
		IDLPath:         "../kstudy_generic.thrift",
		DestServiceName: "destServiceName",
		Method:          "BizMethod1",
	}
}

func fetchReq() *generic.HTTPRequest {
	// 构造request，或者从ginex获取
	body := map[string]interface{}{
		"text": "text",
		"some": map[string]interface{}{
			"id":   1,
			"text": "text",
		},
		"req_items_map": map[string]interface{}{
			"1": map[string]interface{}{
				"id":   1,
				"text": "text",
			},
		},
	}
	data, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	url := "http://example.com/life/client/1/1?v_int64=1&req_items=item1,item2,itme3&cids=1,2,3&vids=1,2,3"
	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	req.Header.Set("token", "1")
	customReq, err := generic.FromHTTPRequest(req) // 考虑到业务有可能使用第三方http request，可以自行构造转换函数
	if err != nil {
		panic(err)
	}
	return customReq
}
