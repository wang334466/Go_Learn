package main

import (
	// "crypto/cipher"
	"fmt"
)

func printMap(cityMap map[string]string) {
	for key, value := range cityMap {
		fmt.Println("key=", key, "value=", value)
	}

}
func changeMap(cityMap map[string]string) {
	cityMap["USA"] = "DC"
	delete(cityMap, "china")
}

func main() {
	cityMap := make(map[string]string)
	cityMap["china"] = "beijing"
	cityMap["japan"] = "tokyo"
	cityMap["USA"] = "NewYork"
	changeMap(cityMap)
	printMap(cityMap)

}
