package utility

import (
	"bytes"
	"context"
	"github.com/a-h/templ"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/svg"
)

func Render(component templ.Component) (string, error) {
	var rendered string
	buf := bytes.NewBufferString(rendered)
	err := component.Render(context.Background(), buf)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func MinifyHtml(str string) string {
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	str, _ = m.String("text/html", str)
	return str
}

func MinifySvg(str string) string {
	m := minify.New()
	m.AddFunc("image/svg+xml", svg.Minify)
	str, _ = m.String("image/svg+xml", str)
	return str
}
