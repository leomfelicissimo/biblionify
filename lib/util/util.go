package util

import (
	"biblionify/lib/types"
	"fmt"
)

// HandleError get a error object if it not is nil
// and throws a panic showing a description declared on event
func HandleError(e error, event string) {
	if e != nil {
		panic(fmt.Errorf("One error has ocurred on %s: %s", event, e))
	}
}

// Mapper get a arr and apply a delegate on each item resulting on a new value
func Mapper(arr []interface{}, delegate types.GenericFunction) []interface{} {
	result := make([]interface{}, len(arr))
	for i, v := range arr {
		result[i] = delegate(v)
	}

	return result
}

func convertToInterface(arr []string) []interface{} {
	iArray := make([]interface{}, len(arr))
	for i, v := range arr {
		iArray[i] = v
	}

	return iArray
}
