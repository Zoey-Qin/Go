package main

import (
	"fmt"
)


func main(){
	read_map()
}


func read_map(){
	var pools map[string]string
	pools = make(map[string]string)
	pools["rbdpool"] = "rbd"
	pools["mypool"] = "rgw"
	pools["testpool"] = "default"

	for _, poolapplication := range pools{
		fmt.Println(poolapplication)
	}
}
