package main

import (
	"fmt"
	"strconv"

	"github.com/dennwc/dom"
)

func createChildEl(index int) *dom.Element {
	child := dom.Doc.CreateElement("li")

	span := dom.Doc.CreateElement("span")
	span.SetTextContent("Child Element " + strconv.Itoa(index))
	child.AppendChild(span)

	return child
}
func main() {
	fmt.Println("Go Wasm Dom")
	fmt.Println("stdout prints to console")
	dom.ConsoleLog("alternatively use dom.ConsoleLog()")

	// App

	app := dom.Doc.GetElementById("app")
	app.SetAttribute("class", "container")

	appHeading := dom.Doc.CreateElement("h1")
	appHeading.SetTextContent("Go -> Wasm")
	appHeading.SetAttribute("class", "title")
	app.AppendChild(appHeading)

	// Counter Demo

	count := 0

	counterContainer := dom.Doc.CreateElement("div")
	counterContainer.SetAttribute("class", "section")
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

	// List Demo

	listChildCount := 10000

	listContainer := dom.Doc.CreateElement("div")
	listContainer.SetAttribute("class", "section")
	app.AppendChild(listContainer)

	listHeading := dom.Doc.CreateElement("h2")
	listHeading.SetTextContent("List of " + strconv.Itoa(listChildCount) + " items rendered with Go wasm")
	listHeading.SetAttribute("class", "subtitle")
	listContainer.AppendChild(listHeading)

	renderListBtn := dom.Doc.NewButton("Render items with Go")
	renderListBtn.SetAttribute("class", "button is-success is-rounded")
	listContainer.AppendChild(renderListBtn)

	listUl := dom.Doc.CreateElement("ul")
	listUl.SetAttribute("class", "list-demo-container")
	listContainer.AppendChild(listUl)

	renderCh := make(chan int, 1)
	renderListBtn.OnClick(func(_ dom.Event) {
		renderCh <- listChildCount
	})

	// Vanilla JS List Demo

	vanillaListContainer := dom.Doc.CreateElement("div")
	vanillaListContainer.SetAttribute("class", "section")
	app.AppendChild(vanillaListContainer)

	vanillaListHeading := dom.Doc.CreateElement("h2")
	vanillaListHeading.SetTextContent("List of " + strconv.Itoa(listChildCount) + " items rendered with Vanilla JS")
	vanillaListHeading.SetAttribute("class", "subtitle")
	vanillaListContainer.AppendChild(vanillaListHeading)

	vanillaFlameChartContainer := dom.Doc.CreateElement("div")
	vanillaFlameChartContainer.SetAttribute("id", "vanilla-flame-chart-container")
	app.AppendChild(vanillaFlameChartContainer)

	vanillaRenderListBtn := dom.Doc.NewButton("Render items with JS")
	vanillaRenderListBtn.SetAttribute("class", "button is-success is-rounded")
	vanillaRenderListBtn.SetAttribute("onClick", "renderVanilla("+strconv.Itoa(listChildCount)+")")
	vanillaListContainer.AppendChild(vanillaRenderListBtn)

	vanillaListUl := dom.Doc.CreateElement("ul")
	vanillaListUl.SetAttribute("id", "vanilla-list-demo-container")
	vanillaListUl.SetAttribute("class", "list-demo-container")
	vanillaListContainer.AppendChild(vanillaListUl)

	// Loop

	for {

		select {

		case count = <-counterCh:
			countValue.SetTextContent(strconv.Itoa(count))

		case count = <-renderCh:
			flameChart := dom.Doc.CreateElement("img")
			flameChart.SetAttribute("src", "./img/flame-chart.png")
			listContainer.AppendChild(flameChart)
			for i := 1; i <= count; i++ {
				listUl.AppendChild(createChildEl(i))
			}
		}
	}

	// idk what this does
	// dom.Loop()
}
