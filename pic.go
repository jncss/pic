package pic

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
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

func doNumber(tag string, fmt format, v reflect.Value) string {
	var fmtNum string

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmtNum = strconv.FormatInt(v.Int(), 10)
		// If len of int part is less than intPartLen, pad with 0
		for len(fmtNum) < fmt.intPartLen {
			fmtNum = "0" + fmtNum
		}
		// If len of int part is greater than intPartLen, left truncate
		if len(fmtNum) > fmt.intPartLen {
			fmtNum = fmtNum[len(fmtNum)-fmt.intPartLen:]
		}
		// If decPartLen is > 0, add decPartLen 0s to the end
		for i := 0; i < fmt.decPartLen; i++ {
			fmtNum += "0"
		}

		// if sign add sign
		if fmt.sign && fmt.signLeft {
			if v.Int() < 0 {
				fmtNum = "-" + fmtNum
			} else {
				fmtNum = "+" + fmtNum
			}
		} else if fmt.sign && !fmt.signLeft {
			if v.Int() < 0 {
				fmtNum += "-"
			} else {
				fmtNum += "+"
			}
		}

	case reflect.Float32, reflect.Float64:
		fmtNum = strconv.FormatFloat(v.Float(), 'f', -1, 64)
		// Extract int part
		intPart := fmtNum[:strings.Index(fmtNum, ".")]
		// If len of int part is less than intPartLen, pad with 0
		for len(intPart) < fmt.intPartLen {
			intPart = "0" + intPart
		}
		// If len of int part is greater than intPartLen, left truncate
		if len(intPart) > fmt.intPartLen {
			intPart = intPart[len(intPart)-fmt.intPartLen:]
		}
		// Extract dec part
		decPart := fmtNum[strings.Index(fmtNum, ".")+1:]
		// If len of dec part is less than decPartLen, pad with 0
		for len(decPart) < fmt.decPartLen {
			decPart += "0"
		}
		// If len of dec part is greater than decPartLen, right truncate
		if len(decPart) > fmt.decPartLen {
			decPart = decPart[:fmt.decPartLen]
		}

		fmtNum = intPart + decPart

		// if sign add sign
		if fmt.sign && fmt.signLeft {
			if v.Float() < 0 {
				fmtNum = "-" + fmtNum
			} else {
				fmtNum = "+" + fmtNum
			}
		} else if fmt.sign && !fmt.signLeft {
			if v.Float() < 0 {
				fmtNum += "-"
			} else {
				fmtNum += "+"
			}
		}
	default:
		// All 0s
		for i := 0; i < fmt.intPartLen+fmt.decPartLen; i++ {
			fmtNum += "0"
		}
		if fmt.sign && fmt.signLeft {
			fmtNum = "+" + fmtNum
		}
		if fmt.sign && !fmt.signLeft {
			fmtNum += "+"
		}
	}

	return fmtNum
}

func doString(tag string, fmt format, v reflect.Value) string {
	var fmtStr string

	switch v.Kind() {
	case reflect.String:
		fmtStr = v.String()
		// If len of string is less than strLen, pad with spaces
		for len(fmtStr) < fmt.strLen {
			fmtStr += " "
		}
		// If len of string is greater than strLen, right truncate
		if len(fmtStr) > fmt.strLen {
			fmtStr = fmtStr[:fmt.strLen]
		}

	default:
		// All spaces
		for i := 0; i < fmt.strLen; i++ {
			fmtStr += " "
		}
	}

	return fmtStr
}
