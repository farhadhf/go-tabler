package tabler_test

import (
	_ "embed"
	"github.com/a-h/templ/generator/htmldiff"
	tabler "github.com/farhadhf/tabler-templ/components"
	"github.com/farhadhf/tabler-templ/internal/utility"
	"testing"
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
