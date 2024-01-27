package main

import (
	"context"
	"os"

	"github.com/a-h/templ"
	tabler "github.com/farhadhf/go-tabler/components"
)

func main() {
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
