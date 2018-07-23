package main

import (
	"fmt"
	"net/http"
)

//func main() {
//
//	arr := []int {1,2,3,4}
//
//	fmt.Println(arr)
//
//	fmt.Println(sum(10,5))
//
//	fmt.Printf("hello, world\n")
//
//	http.HandleFunc("/",HomePage)
//
//	http.ListenAndServe(":8080", nil)
//}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func sum(x int, y int) int {
	return x + y
}
