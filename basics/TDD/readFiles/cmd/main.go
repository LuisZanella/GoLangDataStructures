package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"readfiles/blogposts"
)

func main() {
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	path := bytes.Buffer{}
	path.WriteString(mydir)
	path.WriteString("\\cmd\\posts")

	posts, err := blogposts.NewPostsFromFS(os.DirFS(path.String()))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(posts)
}
