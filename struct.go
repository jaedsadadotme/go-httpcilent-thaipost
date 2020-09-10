package main

type Request struct {
	Status   string   `json:"status"`
	Language string   `json:"language"`
	Barcode  []string `json:"barcode"`
}

type ResponseData struct {
	Status     string      "`json: status`"
	StatusCode int         "`json: statusCode`"
	Datas      interface{} "`json: datas`"
}
type Response struct {
	Response interface{}
}
