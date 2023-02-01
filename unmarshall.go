package pic

import (
	"errors"
	"reflect"
	"strconv"
)

// Unmarshal
func Unmarshall(data string, dest interface{}) error {
	dat := data

	t := reflect.TypeOf(dest)
	if t.Kind() != reflect.Ptr {
		return errors.New("INVALID_POINTER")
	}

	v := reflect.ValueOf(dest).Elem()
	for i := 0; i < t.Elem().NumField(); i++ {
		tag, _ := t.Elem().Field(i).Tag.Lookup("pic")
		l := newLexer(tag)
		p := newParser(l)
		fm, _ := p.parse()

		switch fm.picType {
		case "N":
			doNumberUnmarshal(tag, fm, v.Field(i), &dat)
		case "A":
			doStringUnmarshal(tag, fm, v.Field(i), &dat)
		default:
			return errors.New("INVALID_PIC")
		}
	}

	return nil
}

func doNumberUnmarshal(tag string, fm format, v reflect.Value, data *string) {
	var strNum string

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if fm.sign {
			if fm.signLeft {
				strNum = (*data)[:fm.intPartLen+1]
			} else {
				strNum = (*data)[:fm.intPartLen]
				strNum = string((*data)[fm.intPartLen+fm.decPartLen]) + strNum
			}
		} else {
			strNum = (*data)[:fm.intPartLen]
		}

		num, _ := strconv.ParseInt(strNum, 10, 64)
		v.SetInt(num)

		if fm.sign {
			*data = (*data)[fm.intPartLen+fm.decPartLen+1:]
		} else {
			*data = (*data)[fm.intPartLen+fm.decPartLen:]
		}

	case reflect.Float32, reflect.Float64:
		if fm.sign {
			if fm.signLeft {
				strNum = (*data)[:fm.intPartLen+1]
				strNum += "."
				strNum += (*data)[fm.intPartLen+1 : fm.intPartLen+fm.decPartLen+1]

			} else {
				strNum = (*data)[:fm.intPartLen]
				strNum = string((*data)[fm.intPartLen+fm.decPartLen]) + strNum
				strNum += "."
				strNum += (*data)[fm.intPartLen : fm.intPartLen+fm.decPartLen]
			}
		} else {
			strNum = (*data)[:fm.intPartLen]
			strNum += "."
			strNum += (*data)[fm.intPartLen : fm.intPartLen+fm.decPartLen]
		}

		num, _ := strconv.ParseFloat(strNum, 64)
		v.SetFloat(num)

		if fm.sign {
			*data = (*data)[fm.intPartLen+fm.decPartLen+1:]
		} else {
			*data = (*data)[fm.intPartLen+fm.decPartLen:]
		}
	}
}

func doStringUnmarshal(tag string, frmt format, v reflect.Value, data *string) {
	var str string

	switch v.Kind() {
	case reflect.String:
		str = (*data)[:frmt.strLen]
		v.SetString(str)
	}

	*data = (*data)[frmt.strLen:]
}
