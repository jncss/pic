package pic

import (
	"errors"
	"math"
	"reflect"
	"strconv"
	"strings"
)

// Marshall
func Marshall(data interface{}) (string, error) {
	var result string

	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	for i := 0; i < t.NumField(); i++ {
		tag, _ := t.Field(i).Tag.Lookup("pic")
		l := newLexer(tag)
		p := newParser(l)
		fm, _ := p.parse()

		switch fm.picType {
		case "N":
			result += doNumber(tag, fm, v.Field(i))
		case "A":
			result += doString(tag, fm, v.Field(i))
		default:
			return "", errors.New("INVALID_PIC")

		}
	}

	return result, nil
}

func doNumber(tag string, fm format, v reflect.Value) string {
	var fmNum string

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if v.Int() < 0 {
			fmNum = strconv.FormatInt(v.Int()*-1, 10)
		} else {
			fmNum = strconv.FormatInt(v.Int(), 10)
		}

		for len(fmNum) < fm.intPartLen {
			fmNum = "0" + fmNum
		}

		if len(fmNum) > fm.intPartLen {
			fmNum = fmNum[len(fmNum)-fm.intPartLen:]
		}

		for i := 0; i < fm.decPartLen; i++ {
			fmNum += "0"
		}

		if fm.sign && fm.signLeft {
			if v.Int() < 0 {
				fmNum = "-" + fmNum
			} else {
				fmNum = "+" + fmNum
			}
		} else if fm.sign && !fm.signLeft {
			if v.Int() < 0 {
				fmNum += "-"
			} else {
				fmNum += "+"
			}
		}

	case reflect.Float32, reflect.Float64:
		fmNum = strconv.FormatFloat(math.Abs(v.Float()), 'f', -1, 64)

		intPart := fmNum[:strings.Index(fmNum, ".")]

		for len(intPart) < fm.intPartLen {
			intPart = "0" + intPart
		}

		if len(intPart) > fm.intPartLen {
			intPart = intPart[len(intPart)-fm.intPartLen:]
		}

		decPart := fmNum[strings.Index(fmNum, ".")+1:]

		for len(decPart) < fm.decPartLen {
			decPart += "0"
		}

		if len(decPart) > fm.decPartLen {
			decPart = decPart[:fm.decPartLen]
		}

		fmNum = intPart + decPart

		if fm.sign && fm.signLeft {
			if v.Float() < 0 {
				fmNum = "-" + fmNum
			} else {
				fmNum = "+" + fmNum
			}
		} else if fm.sign && !fm.signLeft {
			if v.Float() < 0 {
				fmNum += "-"
			} else {
				fmNum += "+"
			}
		}
	default:
		for i := 0; i < fm.intPartLen+fm.decPartLen; i++ {
			fmNum += "0"
		}
		if fm.sign && fm.signLeft {
			fmNum = "+" + fmNum
		}
		if fm.sign && !fm.signLeft {
			fmNum += "+"
		}
	}

	return fmNum
}

func doString(tag string, fm format, v reflect.Value) string {
	var fmStr string

	switch v.Kind() {
	case reflect.String:
		fmStr = v.String()
		for len(fmStr) < fm.strLen {
			fmStr += " "
		}
		if len(fmStr) > fm.strLen {
			fmStr = fmStr[:fm.strLen]
		}

	default:
		for i := 0; i < fm.strLen; i++ {
			fmStr += " "
		}
	}

	return fmStr
}
