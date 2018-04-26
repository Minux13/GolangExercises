/*package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func main() {
    url := "https://translate.google.com/#en/es/goal%20to%20work%0Ano%20one%0Aplease"
    response, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    responseString := string(responseData)

    fmt.Println(responseString)
}*/

package main

import (
    "fmt"
    "golang.org/x/net/html"
    "log"
    "net/http"
    "bytes"
	"io"
)

func getElementById(id string, n *html.Node) (element *html.Node, ok bool) {
    for _, a := range n.Attr {
        if a.Key == "id" && a.Val == id {
            return n, true
        }
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        if element, ok = getElementById(id, c); ok {
            return
        }
    }
    return
}

func main() {
    resp, err := http.Get("https://translate.google.com/#en/es/goal%20to%20work%0Ano%20one%0Aplease")
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    root, err := html.Parse(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    element, ok := getElementById("result_box", root)
    if !ok {
        log.Fatal("element not found")
    }
    
    //responseString := string(element)
	/*for _, node := range element {
    	fmt.Printf("%T: %v\n", node, node.Content())
	}*/
    
    
    fmt.Println(renderNode(element))
    
    
    fmt.Println(element.FirstChild)
    /*fmt.Println(element.LastChild)
    fmt.Println(element.PrevSibling)
    fmt.Println(element.NextSibling)
    fmt.Println(element.Type)
    fmt.Println(element.DataAtom)
    fmt.Println(element.Data)
    fmt.Println(element.Namespace)
    fmt.Println(element.Attr)
    
    
    /*for _, a := range element.Attr {
        if a.Key == "value" {
            fmt.Println(a.Val)
            return
        }
    }
    log.Fatal("element missing value")*/
}

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}
