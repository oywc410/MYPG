package main
import (
	"golang.org/x/net/html"
	"fmt"
)

func main() {

}

func soleTitle(doc *html.Node) (title string, err error) {

	type bailout struct {

	}

	defer func() {
		switch p := recover(); p {
		case nil:
		case bailout{}:
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p)
		}
	}()

	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{})
			}
			title = n.FirstChild.Data
		}
	}, nil)

	if title == "" {
		return "", fmt.Errorf("no title element")
	}

	return title, nil
}
