package main

import (
	"fmt"
	"sort"
)

func main() {
	var drink map[string]string
	drink = make(map[string]string,5)
	drink["cola"] = "可乐"
	drink["coffee"]= "咖啡"
	drink["tea"]  = "茶"
	drink["wine"] = "红酒"
	for key,value := range drink {
		fmt.Printf("key : %v,value : %v\n",key,value)
	}
	keys := make([]string,len(drink))
	i := 0
	for k,_ := range drink {
		keys[i] = k
		i ++
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("key : %v , value : %v\n",k,drink[k])
	}

}
