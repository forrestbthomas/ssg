package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"html/template"
	"os"
	"path/filepath"
)

type Entries struct {
	Items []struct {
		Fields struct {
			Title string `json: "title"`
			Body string	`json: "body"`
		} `json: "fields"`
	} `json: "items"`
}


func main() {
	space_id := os.Getenv("SPACE_ID")
	access_token := os.Getenv("ACCESS_TOKEN")
	url := fmt.Sprintf("https://cdn.contentful.com/spaces/%s/entries?access_token=%s", space_id, access_token)
	res, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// handle error
	}
	// fmt.Printf("%s", body)
	var post Entries
	err = json.Unmarshal(body, &post)
	if err != nil {
		fmt.Println("error:", err)
	}

	cwd, _ := os.Getwd()
	file, err := os.Create(filepath.Join(cwd,"./static-site-generator/output.html"))
	if err != nil {
		fmt.Println("error:", err)
	}
	var templatePath = filepath.Join(cwd, "./static-site-generator/post-template.html")
	t, err := template.ParseFiles( templatePath )

	if err != nil {
		fmt.Println("error:", err)
	}
	t.Execute(file, post.Items[0].Fields)
	fmt.Println(post)
}
