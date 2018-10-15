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
	app.SetAttribute("class", "container")

	appHeading := dom.Doc.CreateElement("h1")
	appHeading.SetTextContent("Go -> Wasm")
	appHeading.SetAttribute("class", "title")
	app.AppendChild(appHeading)

	counterContainer := dom.Doc.CreateElement("p")
	counterContainer.SetAttribute("class", "control")
	app.AppendChild(counterContainer)

	counterHeading := dom.Doc.CreateElement("h2")
	counterHeading.SetTextContent("Counter ")
	counterHeading.SetAttribute("class", "subtitle")
	counterContainer.AppendChild(counterHeading)

	countValue := dom.Doc.CreateElement("span")
	countValue.SetTextContent("0")
	counterHeading.AppendChild(countValue)

	decBtn := dom.Doc.NewButton("Decrement")
	decBtn.SetAttribute("class", "button is-warning is-rounded")
	counterContainer.AppendChild(decBtn)

	incBtn := dom.Doc.NewButton("Increment")
	incBtn.SetAttribute("class", "button is-primary is-rounded")
	counterContainer.AppendChild(incBtn)

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
