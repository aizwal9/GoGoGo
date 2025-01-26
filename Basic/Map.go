package main

import "fmt"


func main(){

	var map1 map[string]int

	map1 = make(map[string]int)
	map1["one"] = 1
	map1["two"] = 2
	map1["three"] = 3
	fmt.Println(map1)
	fmt.Println(map1["four"])
	_,ok := map1["four"]
	fmt.Println(ok)

	// map2 := map[string]bool{"one":true,"two":false}
	// fmt.Println(map2["three"])
	// fmt.Println(map2["one"])
	// delete(map2,"one")
	// fmt.Println(map2, len(map2))


}
