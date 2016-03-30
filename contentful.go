package main

import (
	"os"
	"fmt"
	"time"
	"sort"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"html/template"
	"path/filepath"
	"github.com/russross/blackfriday"
)

type Item struct {
	Sys struct {
		CreatedAt time.Time `json: "createdAt"`
	} `json: "sys"`
	Fields struct {
		Title string `json: "title"`
		Body string	`json: "body"`
	} `json: "fields"`
}

type Entries struct {
	Items []Item `json: "items"`
}

type ByTime []Item

func (s ByTime) Len() int {
    return len(s)
}
func (s ByTime) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s ByTime) Less(i, j int) bool {
    return s[i].Sys.CreatedAt.After(s[j].Sys.CreatedAt)
}

func NiceTime(t time.Time) string {
	return t.Local().Format("Mon Jan 2 15:04")
}

func Title(s string) template.HTML {
	interpolatedString := fmt.Sprintf("<h1>%s</h1>", s)
	return template.HTML(interpolatedString)
}

func ToByteThenMD(s string) template.HTML {
	byteString := []byte(s)
	return template.HTML(blackfriday.MarkdownBasic(byteString)[:])
}

func main() {
	space_id := os.Getenv("SPACE_ID")
	access_token := os.Getenv("ACCESS_TOKEN")
	project_path := os.Getenv("PROJECT_PATH")
	url := fmt.Sprintf("https://cdn.contentful.com/spaces/%s/entries?access_token=%s", space_id, access_token)
	res, err := http.Get(url)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var post Entries
	err = json.Unmarshal(body, &post)
	if err != nil {
		fmt.Println("error:", err)
	}
	file, err := os.Create(filepath.Join(project_path,"./dist/output.html"))
	templatePath := filepath.Join(project_path, "./templates/template.html")
	funcMap := template.FuncMap {
		"Title" : 	Title,
		"MD":		ToByteThenMD,
		"NiceTime":	NiceTime,
    }
	t := template.Must(template.New("template.html").Funcs( funcMap ).ParseFiles( templatePath ))
	sort.Sort(ByTime(post.Items))
	t.Execute(file, post.Items)
}
