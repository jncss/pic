package pic

import (
	"errors"
	"log"
	"reflect"
	"strconv"
)

func Marshall(data interface{}) (string, error) {
	var result string

	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	for i := 0; i < t.NumField(); i++ {
		tag, _ := t.Field(i).Tag.Lookup("pic")
		l := newLexer(tag)
		p := newParser(l)
		res, _ := p.parse()

		switch res.picType {
		case "N":
			result += doNumber(tag, res, v.Field(i))
		case "A":
			result += doString(tag, res, v.Field(i))
		default:
			return "", errors.New("INVALID_PIC")

		}
	}

	return result, nil
}

func doNumber(tag string, res result, v reflect.Value) string {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		strNum := strconv.FormatInt(v.Int(), 10)
		// If len of int part is less than intPartLen, pad with 0
		for len(strNum) < res.intPartLen {
			strNum = "0" + strNum
		}
		// If len of int part is greater than intPartLen, left truncate
		if len(strNum) > res.intPartLen {
			strNum = strNum[len(strNum)-res.intPartLen:]
		}
		// If decPartLen is > 0, add decPartLen 0s to the end
		for i := 0; i < res.decPartLen; i++ {
			strNum += "0"
		}

		return strNum

	case reflect.Float32, reflect.Float64:
		log.Println(tag, res, v.Float())
	}

	return ""
}

func doString(tag string, res result, v reflect.Value) string {
	switch v.Kind() {
	case reflect.String:
		log.Println(tag, res, v.String())
	}

	return ""
}
