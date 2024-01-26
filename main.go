package main

import (
	"context"
	"github.com/a-h/templ"
	tabler "github.com/farhadhf/tabler-templ/components"
	"github.com/farhadhf/tabler-templ/components/icons"
	"os"
)

func main() {
	icons.Icon2Fa()
	component := tabler.Page(
		templ.Attributes{
			"dir":  "ltr",
			"lang": "en",
		},
		tabler.Head("Example", nil, nil, nil, true),
		tabler.Body(
			templ.Attributes{
				"class": "d-flex flex-column bg-white",
			},
			LoginPage(),
		),
	)
	component.Render(context.Background(), os.Stdout)
}
