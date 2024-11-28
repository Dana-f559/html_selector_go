package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main(){
	// Validate the args
	if (len(os.Args) != 2 || !strings.HasSuffix(os.Args[1],".html")){
		fmt.Println("USSAGE: go run main.go [file_name].html");
		return;
	}
	
	file_name := os.Args[1];

	// read the file
	file_data_string := string(read_file(file_name));

	// format to html5
	html_node_doc, err := html.Parse(strings.NewReader(file_data_string));

	if err != nil {
		fmt.Println(err)
	}
	
	var buf strings.Builder;
	err1 := html.Render(&buf, html_node_doc)

	if err1 != nil {
		fmt.Printf("Error rendering document: %v", err1)
	}

	// the doc
	string_doc := html.UnescapeString(buf.String())

	
	// Get the input
	reader := bufio.NewReader(os.Stdin);
	element, _ := get_input("element: ", reader)

	new := string_doc;
	
	elements := find_element_html(html_node_doc,element);
	// fmt.Println(elements)
	for j := 0; j < len(elements); j++ {
		new = select_effect(new,elements[j],"\033[31m","\033[0m");
		
	}
	fmt.Println(new);
}

func select_effect(text string, word string, effect_start string, effect_end string) string{
	return strings.ReplaceAll(text, word, effect_start+word+effect_end);
}

func find_element_html(doc *html.Node, element string) []string {
	var elements []string


	var f func(*html.Node)
	f = func(n *html.Node) {
		
		if n.Type == html.ElementNode && n.DataAtom.String() == element {
			
			if n.Parent.DataAtom.String() != element{
				var sb strings.Builder
				err := html.Render(&sb, n)
				if err != nil {
					fmt.Printf("Error rendering node: %v", err)
				}
				
				elements = append(elements, html.UnescapeString(sb.String()));
			}
			
		}

		// Recurse through the DOM tree
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc);

	return elements;
}


func get_input(prompt string, r *bufio.Reader) (string, error){
	fmt.Print(prompt)
	element, err := r.ReadString('\n');

	return strings.TrimSpace(element), err;
}

func read_file(file_name string) []byte{
	data, err := os.ReadFile(file_name);
	if (err != nil) {
		panic(err);
	}

	return data;
}
