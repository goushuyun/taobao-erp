/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by Elvizlai on 2016/05/24 17:30
 */

package misc

type a struct {
	Name string `json:"name"`
	Age  int64
	B    *b
	X    []string `json:"x"`
	C    *a       `json:"c"`
}

type b struct {
	Address string `json:""`
	Email   string `json:""`
	C       *c
}

type c struct {
	Location string
}

// func TestRequest2StructWithValidate(t *testing.T) {
// 	req := &pb.Config{}
//
// 	body := []byte(`{"name":"1234567","age":23, "b":{"address":"17mei","email":"sdrzlyz","c":{"location":"xyz"}}}`)
//
// 	err := Bytes2Struct(body, req, "b.c.location", "name", "c")
//
// 	fmt.Println(req)
//
// 	fmt.Println(err)
// }
