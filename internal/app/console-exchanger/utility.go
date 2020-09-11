package console_exchanger

import (
	"fmt"
	"github.com/valyala/fastjson"
	"strconv"
)

func toFloatAmount(amount string) (float64, error) {
	res, err := strconv.ParseFloat(amount, 64)
	return res, err
}

func getCurrentDst(bytes []byte, src string, dst string) (float64, error) {
	var p fastjson.Parser
	if src == dst {
		return 1, nil
	}
	v, err := p.ParseBytes(bytes) //req := &Response{}
	if err != nil {
		panic(err)
	}
	rates := v.GetObject("rates")
	if rates == nil {
		return 0, fmt.Errorf("can't found %q currency", src)
	}

	if src != "EUR" {
		if rates.Get(src) == nil {
			return 0, fmt.Errorf("can't found %q currency", src)
		}
	}
	currentDst := rates.Get(dst)
	if currentDst == nil {
		return 0, fmt.Errorf("can't found %q currency", dst)
	}

	res, err := currentDst.Float64()
	if err != nil {
		return 0, err
	}
	return res, nil
}
