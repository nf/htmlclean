package main

import "exp/html"

type Matcher func(n *html.Node) bool

func All() Matcher {
	return func(*html.Node) bool { return true }
}

func Tag(name string) Matcher {
	return func(n *html.Node) bool {
		return n.Type == html.ElementNode && n.Data == name
	}
}

func (m Matcher) Walk(n *html.Node, fn func(n *html.Node)) {
	for _, c := range n.Child {
		m.Walk(c, fn)
	}
	if m(n) {
		fn(n)
	}
}

func (m Matcher) Find(n *html.Node) (nodes []*html.Node) {
	for _, c := range n.Child {
		nodes = append(nodes, m.Find(c)...)
	}
	if m(n) {
		nodes = append(nodes, n)
	}
	return nodes
}

func InsertBefore(n *html.Node, nodes ...*html.Node) {
	i := childIndex(n)
	p := n.Parent
	c := make([]*html.Node, len(p.Child)+len(nodes))
	copy(c, p.Child[:i])
	copy(c[i:], nodes)
	copy(c[i+len(nodes):], p.Child[i:])
	for _, n := range nodes {
		n.Parent.Remove(n)
		n.Parent = p
	}
	p.Child = c
}

func childIndex(n *html.Node) int {
	p := n.Parent
	for i := range p.Child {
		if p.Child[i] == n {
			return i
		}
	}
	panic("Node's Parent doesn't list it as a Child")
}
