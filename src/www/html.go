package www

import (
	"log"
	"io/ioutil"
	"fmt"
	"golang.org/x/net/html"
	"gopkg.in/xmlpath.v2"
	"strings"
	"bytes"
)


func ParseAndGet(file string, xPath string) (*xmlpath.Iter, error) {
	fd, e := ioutil.ReadFile(file)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        return nil, e	
    }

    reader := strings.NewReader(string(fd))
    root, err := html.Parse(reader)

    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    var b bytes.Buffer
    html.Render(&b, root)
    fixedHtml := b.String()

    reader = strings.NewReader(fixedHtml)
    xmlroot, xmlerr := xmlpath.ParseHTML(reader)

    if xmlerr != nil {
        log.Fatal(xmlerr)
    	return nil, xmlerr
    }

    return xmlpath.MustCompile(xPath).Iter(xmlroot), nil
}

func GetAttr(attr string, node *xmlpath.Node) string {
    iter := GetChilds("@" + attr, node) 
    if !iter.Next() {
        return ""
    }
    return iter.Node().String()
}

func GetChilds(xPath string, node *xmlpath.Node) *xmlpath.Iter {
    return xmlpath.MustCompile(xPath).Iter(node)
}

func GetValue(node* xmlpath.Node) string {
    return node.String()
}