package main

import(
    "fmt"
    "os"
    "encoding/xml"
    "github.com/ugifractal/xmlquery"
)

func main(){
    fmt.Println("Start")
    f, err := os.Open("./file.xml")
    doc, err := xmlquery.Parse(f)
    if err != nil {
	panic(err)
    }
    element := xmlquery.FindOne(doc, "//xs:element")
    //fmt.Printf("title: %s\n", element.SelectAttr("name"))



    newDoc := &xmlquery.Node{
	Type: xmlquery.DeclarationNode,
	Data: "xml",
	Prefix: "xs",
	Attr: []xml.Attr{
	    xml.Attr{Name: xml.Name{Local: "version"}, Value: "1.0"},
	},
    }
    root := &xmlquery.Node{
	Data: "element",
	Prefix: "xs",
	Type: xmlquery.ElementNode,
	Attr: []xml.Attr{
	    xml.Attr{Name: xml.Name{Local: "name"}, Value: "note"},
	},
    }

    schema := &xmlquery.Node{
	Data: "schema",
	Prefix: "xs",
	Type: xmlquery.ElementNode,
	Attr: []xml.Attr{
	    xml.Attr{Name: xml.Name{Local: "xmlns:xs"}, Value: "http://www.w3.org/2001/XMLSchema"},
	},
    }

    newDoc.FirstChild = schema
    schema.FirstChild = root
    root.FirstChild = element.FirstChild

    //fmt.Printf("****[%s]***(%s)*****<%s>\n", root.FirstChild.NextSibling.Type, root.FirstChild.NextSibling.SelectAttr("code"), root.FirstChild.NextSibling.Prefix)
    /*
    node := element.FirstChild
    for {
	node = node.NextSibling
	if node == nil {
	    break
	}
	fmt.Printf("--%s--%s---\n", node.Data, node.SelectAttr("code"))
	//if root.FirstChild == nil {
	//    root.FirstChild = node  
	//}
    }
	//myNode := hasChild.NextSibling
	//root.FirstChild = complexTypeNode
    */
    fmt.Println(newDoc.OutputXML(true))

}
