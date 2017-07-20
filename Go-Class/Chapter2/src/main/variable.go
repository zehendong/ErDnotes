package main

import "go/ast"

func main() {

	/**
	申明变量
	 */
	var v1 int
	var v2 string
	var v3 [12]int
	var v4 []int
	var v5 struct{
		f int
	}
	var v6 *int   //指针
	var v7 map[string]int    //map,key is string type and value is int type
	var v8 func(a int) int
	//define varible together
	var (
		v9 int
		v10 string
	)
}
