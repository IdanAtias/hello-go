package main

import (
    "fmt"
    //"github.com/IdanAtias/hello-go/morestrings" // for fetching this pkg from the remote repo
    "hello-go/morestrings" // for "local dev"
)

func main() {
	fmt.Println(morestrings.ReplaceRunes("hello go!", 'l', '1'))
}
