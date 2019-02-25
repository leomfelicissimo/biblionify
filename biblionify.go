package main

import (
	"biblionify/lib/util"
	"fmt"
)

func applyLastName(i interface{}) interface{} {
	firstName := i.(string)
	return firstName + " Felicissimo"
}

func convertToInterface(arr []string) []interface{} {
	s := make([]interface{}, len(arr))
	for i, v := range arr {
		s[i] = v
	}

	return s
}

func main() {
	content := util.ParseJSONFile("./biblefiles/mini.json")
	fmt.Println(content)

	arr := []string{"Leonardo", "Guilherme", "Isabella", "Lucas", "Jéssica"}
	mapResult := util.Mapper(convertToInterface(arr), applyLastName)
	fmt.Println(mapResult)
}
