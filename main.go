package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/ShiBinjie/chimpanzee/lib/log"
	"github.com/ShiBinjie/chimpanzee/service/data"
)

func main() {
	fundId := flag.String("fund", "090010", "fund id")
	timeFrom := flag.String("from", "2018-01-01", "fund id")
	timeTo := flag.String("to", "2019-01-01", "fund id")
	data, err := data.GetData(*fundId, *timeFrom, *timeTo)
	if err != nil {
		log.Logger.Errorf("get data fail! err:%s", err.Error())
	}

	file, err := os.Create("data.txt")
	if err != nil {
		log.Logger.Errorf("create file fail! err:%s", err.Error())
	}
	defer file.Close()
	for _, onedata := range data {
		iopv, _ := strconv.ParseFloat(onedata.IOPV, 32)
		tcnv, _ := strconv.ParseFloat(onedata.TCNV, 32)
		str := fmt.Sprintf("time:%s, 单位净值: %.3f, 累计净值: %.3f \n", onedata.Time, iopv, tcnv)
		toWrite := []byte(str)
		file.Write(toWrite)
	}
}
