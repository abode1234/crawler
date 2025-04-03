package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector(

		colly.MaxDepth(3),
	)
	c.OnHTML("a[href]", func(h *colly.HTMLElement) {

		link := h.Attr("href")

		fmt.Println(link)
	})

	c.Visit("https://en.wikipedia.org/")

}
