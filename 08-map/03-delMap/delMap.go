package main

import "fmt"

func main(){
	scence := make(map[string]int)

	scence["route"] = 66
	scence["ZZ"] = 96
	scence["brazil"] = 4

	// delete(map, 键)  删除map中一组键值对
	delete(scence, "brazil")

	for k,v := range scence {
		fmt.Println(k,v)
	}
}
