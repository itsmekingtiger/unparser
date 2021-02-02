package main

import (
	"fmt"

	"github.com/itsmekingtiger/unparser"
)

func main() {
	var r = unparser.DefaultPathRouter()
	r.Add("/pet/$ID")
	r.Add("/pet/$ID/$ATTRNAME/$ATTRVALUE")
	r.Add("/pet/$ID/name/$NAME")
	r.Add("/pet/$ID/age/$AGE")

	r.Print()

	fmt.Println(r.Match("/pet/1"))
	fmt.Println(r.Match("/pet/2/이빨갯수/12"))
	fmt.Println(r.Match("/pet/2/발톱갯수/20"))
	fmt.Println(r.Match("/pet/3/name/애옹이"))
	fmt.Println(r.Match("/pet/4/age/4"))

	r.Add("/menu/drink/coffe/$item")
	r.Add("/menu/drink/juice/$item")

	fmt.Println(r.Match("/menu/drink/juice/orange"))
	fmt.Println(r.Match("/menu/drink/juice/apple"))
	fmt.Println(r.Match("/menu/drink/juice/amaricano"))
	fmt.Println(r.Match("/menu/drink/juice/espresso"))
	fmt.Println(r.Match("/menu/drink/pie/applepie"))
}
