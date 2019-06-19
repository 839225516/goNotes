package main

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	//声明技能结构
	type Skill struct {
		Name  string
		Level int
	}

	// 声明角色结构
	type Actor struct {
		Name string
		Age  int

		Skill []Skill
	}

	// 填充基本角色数据
	a := Actor{
		Name: "lilei",
		Age:  26,
		Skill: []Skill{
			{Name: "Roll and roll", Level: 1},
			{Name: "Flash you doy eye", Level: 2},
			{Name: "Time to have Lunch", Level: 3},
		},
	}

	if result, err := MarshalJson(a); err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}

func MarshalJson(v interface{}) (string, error) {
	var b bytes.Buffer

	// 将任意值转换成Json并输出到缓冲中
	if err := writeAny(&b, reflect.ValueOf(v)); err == nil {
		return b.String(), nil
	} else {
		return "", err
	}
}

func writeAny(buff *bytes.Buffer, value reflect.Value) error {
	switch value.Kind() {
	case reflect.String:
		buff.WriteString(strconv.Quote(value.String()))
	case reflect.Int:
		buff.WriteString(strconv.FormatInt(value.Int(), 10))
	case reflect.Slice:
		return writeSlice(buff, value)
	case reflect.Struct:
		return writeStruct(buff, value)
	default:
		return errors.New("unsupport kind: " + value.Kind().String())

	}

	return nil
}

func writeSlice(buff *bytes.Buffer, value reflect.Value) error {
	buff.WriteString("[")

	for s := 0; s < value.Len(); s++ {
		sliceValue := value.Index(s)
		writeAny(buff, sliceValue)

		if s < value.Len()-1 {
			buff.WriteString(",")
		}
	}

	buff.WriteString("]")
	return nil
}

func writeStruct(buff *bytes.Buffer, value reflect.Value) error {
	valueType := value.Type()

	buff.WriteString("{")

	for i := 0; i < value.NumField(); i++ {
		fieldValue := value.Field(i)

		fieldType := valueType.Field(i)

		buff.WriteString("\"")

		buff.WriteString(fieldType.Name)

		buff.WriteString("\":")

		writeAny(buff, fieldValue)

		if i < value.NumField()-1 {
			buff.WriteString(",")
		}
	}

	buff.WriteString("}")
	return nil
}
