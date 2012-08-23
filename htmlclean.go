package main

import (
	"exp/html"
	"log"
	"os"
)

func main() {
	root, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	body := Tag("body").Find(root)[0]
	collapseSpans(body)
	cleanClasses(body)
	html.Render(os.Stdout, body)
}

func collapseSpans(n *html.Node) {
	Tag("span").Walk(n, func(n *html.Node) {
		if len(n.Child) > 0 {
			InsertBefore(n, n.Child...)
		}
		n.Parent.Remove(n)
	})
}

func cleanClasses(n *html.Node) {
	All().Walk(n, func(n *html.Node) {
		RemoveAttr(n, "class")
	})
}
