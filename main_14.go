package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

var tformat = time.RFC850

type Purchase struct {
	FirstName, LastName string
	Address             string
	Phone               string
	Age                 float32
	Male                bool
	Item                string
	ShipDate            string
}

var purchases = []Purchase{
	{LastName: "Feigenbaum", ShipDate: time.Now().Format(tformat), Male: true, Item: "Go for Java Programs"},
}

const purchaseTemplate = `
Dear {{if.Male}} Mr.{{else}}Ms.{{end}} {{.LastName}},
Your purchase request for "{{.Item}}" has been received.
It will be shipped on {{.ShipDate}}.
Thank you for your purchase.
`

func runTemplate(p *Purchase, f *os.File, t *template.Template) (err error) {
	err = t.Execute(f, *p)
	return
}

func genTemplate(prefix string) {
	t := template.Must(template.New("purchases").Parse(purchaseTemplate))

	for i, p := range purchases {
		f, err := os.Create(fmt.Sprintf("%s-%d.tmpl", prefix, i))
		if err != nil {
			log.Fatalf("failed to create file, cause:", err)
		}
		err = runTemplate(&p, f, t)
		if err != nil {
			log.Fatalf("failed to run template, cause:", err)
		}
	}
}

var functionMap = template.FuncMap{

	"formatDate": formatDate,
	"formatTime": formatTime,
}

var parsedTemplate *template.Template

func loadTemplate(path string) {
	parsedTemplate = template.New("tod")
	parsedTemplate.Funcs(functionMap)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicf("failed reading template %s :%v", path, err)
	}

	if _, err := parsedTemplate.Parse(string(data)); err != nil {
		log.Panicf("failed parsing template: %s :%v", path, err)
	}
}

func formatDate(dt time.Time) string {
	return dt.Format("Mon Jan _2 2006")
}

func formatTime(dt time.Time) string {
	return dt.Format("15:04:05 MST")
}

type TODData struct {
	TOD time.Time
}

func processTODRequest(w http.ResponseWriter, r *http.Request) {
	var data = &TODData{time.Now()}
	parsedTemplate.Execute(w, data) // 假设不能失败
}

var serverPort = 8085

func timeServer() {
	loadTemplate("D:\\aa-my\\1\\my-go-server\\tod.tmpl")
	http.HandleFunc("/tod", processTODRequest)
	spec := fmt.Sprintf(":%d", serverPort)
	if err := http.ListenAndServe(spec, nil); err != nil {
		log.Fatalf("failed to start server on port %d: %v", serverPort, err)
	}
	log.Println("server exited")
}

func main() {
	timeServer()
}
