package main

import (
	"fmt"

	"github.com/itsmekingtiger/unparser"
)

func main() {
	var r = unparser.DefaultPathRouter()
	r.Add("/pet/$ID", func(mRes *unparser.MatchResult) {
		fmt.Println("ID 매칭:" + mRes.PathVariables["ID"])
	})
	r.Add("/pet/$ID/$ATTRNAME/$ATTRVALUE", func(mRes *unparser.MatchResult) {
		fmt.Println("ID매칭:" + mRes.PathVariables["ID"])
		fmt.Println("ATTRNAME매칭:" + mRes.PathVariables["ATTRNAME"])
		fmt.Println("ATTRVALUE매칭:" + mRes.PathVariables["ATTRVALUE"])
	})
	r.Add("/pet/$ID/name/$NAME", func(mRes *unparser.MatchResult) {
		fmt.Println("ID매칭:" + mRes.PathVariables["ID"])
		fmt.Println("NAME매칭:" + mRes.PathVariables["NAME"])
	})
	r.Add("/pet/$ID/age/$AGE", func(mRes *unparser.MatchResult) {
		fmt.Println("ID매칭:" + mRes.PathVariables["ID"])
		fmt.Println("AGE매칭:" + mRes.PathVariables["AGE"])
	})

	r.Print()

	r.Match("/pet/1")
	fmt.Println()
	r.Match("/pet/2/이빨갯수/12")
	fmt.Println()
	r.Match("/pet/2/발톱갯수/20")
	fmt.Println()
	r.Match("/pet/3/name/애옹이")
	fmt.Println()
	r.Match("/pet/4/age/4")
	fmt.Println()

	r.Add("/menu/drink/coffe/$item", func(mRes *unparser.MatchResult) {
		fmt.Println(mRes.PathVariables["item"])
	})
	r.Add("/menu/drink/juice/$item", func(mRes *unparser.MatchResult) {
		fmt.Println(mRes.PathVariables["item"])
	})

	r.Match("/menu/drink/juice/orange")
	r.Match("/menu/drink/juice/apple")
	r.Match("/menu/drink/juice/amaricano")
	r.Match("/menu/drink/juice/espresso")
	r.Match("/menu/drink/pie/applepie")
}
