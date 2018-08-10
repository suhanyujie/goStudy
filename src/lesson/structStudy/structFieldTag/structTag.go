package structFieldTag

import (
	"reflect"
	"fmt"
)

type TagType struct {
	Field1 bool "this is a bool type"
	Filed2 string "this is a string type"
	Filed3 int "this is a int type"
}

func RefTag(tt TagType, fieldIndex int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(fieldIndex)
	fmt.Printf("%v\n", ixField.Tag)
}
