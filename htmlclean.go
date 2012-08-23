package main

import (
	"exp/html"
	"os"
	"log"
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
		if len(n.Child) == 0 {
			// empty span
			n.Parent.Remove(n)
		} else {
			InsertBefore(n, n.Child...)
			n.Parent.Remove(n)
		}
	})
}

func cleanClasses(n *html.Node) {
	All().Walk(n, func(n *html.Node) {
		for i, a := range n.Attr {
			if a.Key == "class" {
				copy(n.Attr[i:], n.Attr[i+1:])
				n.Attr = n.Attr[:len(n.Attr)-1]
				return
			}
		}
	})
}
