/*
go mod init example.com/hello
go get golang.org/x/example/hello/reverse
*/
package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Println(reverse.String("Hello"), reverse.Int(24601))
}
