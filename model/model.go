package model

type APIResult struct {
	Result Result `json:"result"`
}

type Result struct {
	Status APIStatus `json:"status"`
	Datas  APIDatas  `json:"data"`
}

type APIStatus struct {
	Code int `json:"code"`
}

type APIDatas struct {
	Datas []Data `json:"data"`
	Size  string `json:"total_num"`
}

type Data struct {
	Time string `json:"fbrq"`
	IOPV string `json:"jjjz"` //单位净值
	TCNV string `json:"ljjz"` //累计净值
}
