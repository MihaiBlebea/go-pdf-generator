package diploma

import (
	"bytes"
	"embed"
	"html/template"
)

//go:embed template/*.gohtml
var temp embed.FS

func parseTemplate(data interface{}) (string, error) {
	t, err := template.ParseFS(temp, "template/*.gohtml")
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = t.ExecuteTemplate(buf, "index", data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
