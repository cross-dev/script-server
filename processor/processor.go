package processor

import (
	"fmt"
	"io"
	"net/url"
	"text/template"
)

func writeError(msg string, err error, writer io.Writer) {
	writer.Write([]byte(fmt.Sprint(msg, err)))
}

func Process(values url.Values, tmpl string, writer io.Writer) {
	functions := make(template.FuncMap)
	generator := func(key string) interface{} {
		return func() string { return values.Get(key) }
	}
	for key := range values {
		functions[key] = generator(key)
	}
	theTemplate, err := template.New("").Funcs(functions).Parse(tmpl)
	if err != nil {
		writeError("Error parsing the template: ", err, writer)
	} else {
		err = theTemplate.Execute(writer, nil)
		if err != nil {
			writeError("Error executing the template: ", err, writer)
		}
	}
}
