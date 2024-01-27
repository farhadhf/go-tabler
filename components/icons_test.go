package tabler_test

import (
	_ "embed"
	"testing"

	"github.com/a-h/templ/generator/htmldiff"
	tabler "github.com/farhadhf/go-tabler/components"
	"github.com/farhadhf/go-tabler/internal/utility"
)

//go:embed test-data/icon-2fa.svg
var icon string

func TestIcon(t *testing.T) {
	component := tabler.Icon("2fa")
	got, err := utility.Render(component)
	if err != nil {
		t.Fatal(err)
	}

	diff, err := htmldiff.DiffStrings(utility.MinifySvg(icon), utility.MinifySvg(got))
	if err != nil {
		t.Fatal(err)
	}
	if diff != "" {
		t.Error(diff)
	}
}
