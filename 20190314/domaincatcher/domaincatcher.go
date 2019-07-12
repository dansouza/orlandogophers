# https://ayan.net/employment-challenge/
package main

import "fmt"

func ExtractDomains(s []string) []string {

}

func main() {
	input := []string{
		"there's a http://golang.org/foo/bar, https://www.google.com/, http://web.foo.bar/",
		"then there is https://foo.com",
	}
	output := ExtractDomains(input)
	fmt.Println(output)
}
