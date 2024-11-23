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

	file_data_string := string(read_file(file_name));
	
	// Get the input
	reader := bufio.NewReader(os.Stdin);
	element, _ := get_input("element: ", reader)

	
	// Parce the html

	doc, err := html.Parse(strings.NewReader(file_data_string));

	if (err != nil) {
		panic(err);
	}
	new := file_data_string;
	// elements := [2]string{"<div>","</div>"}
	elements := find_element_html(doc,element);
	for j := 0; j < len(elements); j++ {
		new = select_effect(new,elements[j],"\033[31m","\033[0m");
		
		// new = select_effect(new,elements[j],"_","_a");

	}

	fmt.Println(new);
	// fmt.Println(elements)
}

func select_effect(text string, word string, effect_start string, effect_end string) string{
	return strings.ReplaceAll(text, word, effect_start+word+effect_end);
}


func arrange(n *html.Node) string{
	// Serialize the current node, including its children
	var result string

	if n.Type == html.ElementNode {
		// Construct the opening tag with attributes
		result += "<" + n.Data
		for _, attr := range n.Attr {
			result += fmt.Sprintf(" %s=\"%s\"", attr.Key, attr.Val)
		}
		result += ">"
	} else if n.Type == html.TextNode {
		// Append text content
		result += n.Data
	}

	// Process children
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		result += arrange(child)
	}

	// Close the tag if it's an element node
	if n.Type == html.ElementNode {
		result += "</" + n.Data + ">"
	}

	return result
}

func find_element_html(doc *html.Node, element string) []string {
	var elements []string


	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.DataAtom.String() == element {
			// Serialize this element and add to results
			elements = append(elements, arrange(n));
		}

		// Recurse through the DOM tree
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	// Start processing from the root node
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