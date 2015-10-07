package processor

import (
	"io"
	"net/url"
	"text/template"
)

func Process(values url.Values, tmpl string, writer io.Writer) error {
	var err error
	functions := make(template.FuncMap)
	generator := func(key string) interface{} {
		return func() string { return values.Get(key) }
	}
	for key := range values {
		functions[key] = generator(key)
	}
	var theTemplate *template.Template
	if theTemplate, err = template.New("").Funcs(functions).Parse(tmpl); err == nil {
		err = theTemplate.Execute(writer, nil)
	}
	return err
}
