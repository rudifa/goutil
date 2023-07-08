package util

import (
	"fmt"
	"log"
	"reflect"
)

// PrintFields prints the exported fields of an object
func PrintFields(tag string, obj interface{}) {
	objValue := reflect.ValueOf(obj)
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}
	objType := objValue.Type()

	log.Printf("PrintFields fields of %s:\n", tag)
	for i := 0; i < objValue.NumField(); i++ {
		field := objValue.Field(i)
		fieldName := objType.Field(i).Name

		// Check if the field is exported
		if field.CanInterface() {
			fieldValue := field.Interface()
			fmt.Printf("%s: %v\n", fieldName, fieldValue)
		}
	}
}
