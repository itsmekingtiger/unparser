package main

import (
	"fmt"

	"github.com/itsmekingtiger/unparser"
)

func main() {
	var r = unparser.DefaultPathRouter()
	r.Add("/user")
	r.Add("/user/Jhon")
	r.Add("/user/Jhon/red")
	r.Add("/user/Jhon/green")
	r.Add("/user/홍길동")
	r.Add("/user/홍길동/address")
	r.Add("/user/홍길동/address/detail")
	r.Add("/user/홍길동/address/detail/dong")
	r.Add("/user/홍길동/address/detail/ho")
	r.Add("/pet")
	r.Add("/pet/1")
	r.Add("/pet/1/name")
	r.Add("/pet/1/name/middle")
	r.Add("/pet/1/name/middle/sir")
	r.Add("/pet/2")
	r.Add("/pet/2/age")

	r.Print()

	fmt.Println(r.Match("/user"))
	fmt.Println(r.Match("/user/Jhon"))
	fmt.Println(r.Match("/user/Song"))
}
