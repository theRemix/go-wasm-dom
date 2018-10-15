package main

import (
	"fmt"
	"strconv"

	"github.com/dennwc/dom"
)

func main() {
	fmt.Println("Go Wasm Dom")
	fmt.Println("stdout prints to console")
	dom.ConsoleLog("alternatively use dom.ConsoleLog()")

	count := 0

	app := dom.Doc.GetElementById("app")
	app.SetTextContent("HELLO")

	counterContainer := dom.Doc.CreateElement("p")
	app.AppendChild(counterContainer)

	counterHeading := dom.Doc.CreateElement("h2")
	counterHeading.SetTextContent("Counter ")
	counterContainer.AppendChild(counterHeading)

	countValue := dom.Doc.CreateElement("span")
	countValue.SetTextContent("0")
	counterHeading.AppendChild(countValue)

	incBtn := dom.Doc.NewButton("Increment")
	counterContainer.AppendChild(incBtn)

	decBtn := dom.Doc.NewButton("Decrement")
	counterContainer.AppendChild(decBtn)

	counterCh := make(chan int, 1)
	incBtn.OnClick(func(_ dom.Event) {
		counterCh <- count + 1
	})
	decBtn.OnClick(func(_ dom.Event) {
		counterCh <- count - 1
	})

	for {
		count = <-counterCh
		countValue.SetTextContent(strconv.Itoa(count))
	}

	// idk what this does
	// dom.Loop()
}
