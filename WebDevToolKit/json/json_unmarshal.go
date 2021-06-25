package main

import (
	"encoding/json"
	"fmt"
)

var rcvd string

type data []struct {
	Code    int    `json:"code"`
	Descrip string `json:"Descrip"`
}

func main() {
	rcvd = `[{"code":200,"Descrip":"OK"},{"code":201,"Descrip":" Created"},{"code":202,"Descrip":"Accepted"},{"code":203,"Descrip":" Non-authoritative Information"},{"code":204,"Descrip":" No Content"},{"code":205,"Descrip":" Reset Content"},{"code":206,"Descrip":" Partial Content"},{"code":207,"Descrip":" Multi-Status"},{"code":208,"Descrip":" Already Reported"},{"code":226,"Descrip":" IM Used"}]`
	var un_m data
	err := json.Unmarshal([]byte(rcvd), &un_m)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range un_m {
		fmt.Println(v.Code, ":", v.Descrip)
	}
}
