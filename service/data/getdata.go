package data

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/ShiBinjie/chimpanzee/model"
	"github.com/go-resty/resty"
)

const (
	API      = "http://stock.finance.sina.com.cn/fundInfo/api/openapi.php/CaihuiFundInfoService.getNav?symbol=%s&datefrom=%s&dateto=%s&page=%d"
	PAGESIZE = 21
)

func FormatURL(api, fundId, from, to string, page int) (url string) {
	return fmt.Sprintf(api, fundId, from, to, page)
}

func SizeofData(url string) (size int, err error) {
	response, err := resty.R().Get(url)
	if err != nil {
		return 0, err
	}
	raw := string(response.Body())

	data := model.APIResult{}
	if err = json.Unmarshal([]byte(raw), &data); err != nil {
		return 0, err
	}

	size, _ = strconv.Atoi(data.Result.Datas.Size)
	return size, nil
}

func GetData(fundId, from, to string) ([]model.Data, error) {
	url := FormatURL(API, fundId, from, to, 0)
	size, err := SizeofData(url)
	if err != nil {
		return nil, err
	}

	var pages int
	if size%PAGESIZE > 0 {
		pages = size/PAGESIZE + 1
	} else {
		pages = size / PAGESIZE
	}

	var result = []model.Data{}
	for p := pages; p >= 0; p-- {
		pn, err := GetDataPage(fundId, from, to, p)
		if err != nil {
			return nil, err
		}
		//fmt.Println(pn)
		result = append(result, pn...)
	}

	return result, nil
}

func GetDataPage(fundId, from, to string, page int) ([]model.Data, error) {
	url := FormatURL(API, fundId, from, to, page)
	response, err := resty.R().Get(url)
	if err != nil {
		return nil, err
	}
	raw := string(response.Body())

	data := model.APIResult{}
	if err = json.Unmarshal([]byte(raw), &data); err != nil {
		return nil, err
	}

	return data.Result.Datas.Datas, nil
}
