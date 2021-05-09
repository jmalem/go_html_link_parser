package main

import (
	"fmt"
	"os"
	"golang.org/x/net/html"
)

var g []linkNode

type linkNode struct {
	href string
	data string
}

func listNode(n *html.Node) {

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, obj := range n.Attr {
			link := linkNode{obj.Val, n.FirstChild.Data}
			g = append(g, link)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		listNode(c)
	}
}

func main() {
	htmlFile, err := os.Open("ex4.html")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	root, err := html.Parse(htmlFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	// fmt.Println(root.Data)
	// fmt.Println(root.Type)
	listNode(root)

	fmt.Println(g)
}