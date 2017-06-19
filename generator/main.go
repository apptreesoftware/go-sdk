package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"time"

	"fmt"

	"flag"
	"github.com/apptreesoftware/go-sdk"
)

func main() {
	configUrl := flag.String("url", "", "Url to configuration")
	outputFile := flag.String("o", "", "Output file location")
	flag.Parse()
	if *configUrl == "" || *outputFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	//configUrl := "https://accruent-cloud-dev.herokuapp.com/dataset/purchaseorders/describe"

	resp, err := http.Get(*configUrl)
	if err != nil {
		log.Fatalf("Unable to fetch configuration %s", err.Error())
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading configuration %s", err.Error())
	}
	config := apptree.Configuration{}
	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Fatalf("Unable to parse config json %s", err.Error())
	}
	structs := generateStructTemplatesFromConfig(&config, config.Name)

	file := generateCode(structs, *configUrl, *outputFile)
	formatFile(file)
}

func generateStructTemplatesFromConfig(config *apptree.Configuration, name string) []structTemplate {
	name = programmersName(name)
	parentTemplate := structTemplate{Name: name, Attributes: []structAttribute{}}
	structs := []structTemplate{}
	templateAttr := []structAttribute{}
	for _, attr := range config.Attributes {
		log.Printf("Type %s", attr.Type)
		switch attr.Type {
		case apptree.Type_Relationship:
			childConfig := attr.RelatedConfiguration
			childStructs := generateStructTemplatesFromConfig(childConfig, attr.Name)
			structs = append(structs, childStructs...)
			structAttr := structAttribute{
				Index: attr.Index,
				Name:  programmersName(attr.Name),
				Type:  fmt.Sprintf("[]%s", programmersName(attr.Name)),
			}
			templateAttr = append(templateAttr, structAttr)
			break
		case apptree.Type_SingleRelationship:
			childConfig := attr.RelatedConfiguration
			childStructs := generateStructTemplatesFromConfig(childConfig, attr.Name)
			structs = append(structs, childStructs...)
			break
		default:
			//log.Printf("Type %s", attr.Type)
			structAttr := structAttribute{
				Index: attr.Index,
				Name:  programmersName(attr.Name),
				Type:  attr.Type.ToAppTreeTypePackageName(),
			}
			templateAttr = append(templateAttr, structAttr)
		}
	}
	parentTemplate.Attributes = templateAttr
	structs = append(structs, parentTemplate)
	return structs
}

func formatFile(file *os.File) {
	path := file.Name
	arg := fmt.Sprintf("-w %s", path)
	cmd := exec.Command("gofmt", arg)
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
}

func programmersName(str string) string {

	name := str
	name = strings.Replace(name, " ", "", len(name))
	name = strings.Replace(name, ".", "", len(name))
	name = strings.Replace(name, "&", "And", len(name))
	return name
}

func generateCode(structs []structTemplate, url string, location string) *os.File {
	outputFile, err := os.Create(location)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	for i := len(structs) - 1; i >= 0; i-- {
		tmp := structs[i]
		structGenTemplate.Execute(outputFile, struct {
			Timestamp  time.Time
			URL        string
			Attributes []structAttribute
			StructName string
		}{
			Timestamp:  time.Now(),
			URL:        url,
			Attributes: tmp.Attributes,
			StructName: tmp.Name,
		})
	}
	return outputFile
}

type structTemplate struct {
	Name       string
	Attributes []structAttribute
}

type structAttribute struct {
	Index int
	Name  string
	Type  string
}

var structGenTemplate = template.Must(template.New("").Parse(`
type {{ .StructName }} struct {
	{{- range .Attributes }}
		{{ .Name }}   {{ .Type }}   ` + "`index:\"{{ .Index }}\"`" + `
	{{- end }}
}
`))
