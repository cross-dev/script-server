package processor

import (
	"net/url"
	"text/template"
	"fmt"
	"io"
)

func extract(name string, values *url.Values) string {
	return values.Get(name)
}

type Processor interface {
	Process(values *url.Values, writer io.Writer)
}

type theProcessor struct {
	theTemplate *template.Template
}

type brokenProcessor struct {
	err error
}

func NewProcessor(tmpl string) Processor {
	functions := template.FuncMap{
		"extract": extract,
	}
	theTemplate, err := template.New("replaceArgs").Funcs(functions).Parse(tmpl)
	if err != nil {
		proc := brokenProcessor{err: err}
		return &proc
	}
	proc := theProcessor{theTemplate}
	return &proc
}

func (this *brokenProcessor) Process(values *url.Values, writer io.Writer) {
	writer.Write([]byte(fmt.Sprintf("Error building the template: %s", this.err)))
}

func (this *theProcessor) Process(values *url.Values, writer io.Writer) {
	err := this.theTemplate.Execute(writer, values)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf("Error processing values: %s", err)))
	}
}
