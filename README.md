# HTML Element Highlighter in Go
This Go program reads an HTML file, parses it into a DOM, and highlights all occurrences of a specified HTML element by wrapping them with terminal color codes (red) when printed. It helps visualize and emphasize specific elements from an HTML document in the terminal.

## Features
* Validates command line argument to ensure a single .html file input.

* Parses the input HTML file into a DOM tree using golang.org/x/net/html.

* Finds and extracts all distinct elements matching a user-specified tag.

* Highlights these elements by injecting terminal color escape sequences.

* Prints the modified HTML string to standard output.

## Prerequisites
* Go 1.18+

* golang.org/x/net/html package (can be installed with go get)

## Usage
```bash
go run main.go your_file.html
```
After running, the program prompts:

```text
element:
```
Type an HTML tag name (e.g., div, p, span) and press Enter. The program will output the HTML with all matching elements highlighted in red.

### Code Overview
* Argument Validation: Checks that exactly one .html file is provided.

* File Reading: Reads the entire HTML file content to a string.

* HTML Parsing: Uses html.Parse from golang.org/x/net/html to build a DOM tree.

* Element Search: Recursively traverses the DOM to find elements with matching tag names and collects their HTML.

* Highlighting: Replaces each found element occurrence in the output string by wrapping it with red ANSI color codes (\033[31m and \033[0m).

* User Input: Reads element tag choice from standard input.

* Output: Prints the final highlighted HTML document.

### Example
Given an HTML file example.html:

```xml
<html>
  <body>
    <p>Hello</p>
    <div>World</div>
    <p>Another Paragraph</p>
  </body>
</html>
```
Running `go run main.go example.html` and entering p will highlight both <p> elements in red in the output.
