package main

import "fmt"

func main() {
	map1 := make(map[string]int)
	test1 := []string{"hello", "bello", "mello", "tello", "hello", "tello", "bello"}
	for index, value := range test1 {
		val, ok := map1[value]
		if ok == false {
			map1[value] = index + 0
		} else {
			map1[value] = val + index
		}
	}
	fmt.Println("Get all the keys and values of the map:")
	for key, value := range map1 {
		fmt.Println(key, value)
	}

	fmt.Println("Get all the keys of the map:")
	keys := make([]string, 0, len(map1))
	for key, _ := range map1 {
		keys = append(keys, key)
	}
	fmt.Println(keys)
}
