package alerts_test

import (
	_ "embed"
	"testing"

	"github.com/a-h/templ/generator/htmldiff"
	alerts "github.com/farhadhf/go-tabler/components/alerts"
	"github.com/farhadhf/go-tabler/internal/utility"
)

//go:embed test-data/alert.html
var alert string

func TestAlert(t *testing.T) {
	alertParams := alerts.AlertParams{
		Title:       "title",
		Description: "description",
		Color:       "",
		Dismissible: false,
		Important:   false,
	}
	component := alerts.Alert(alertParams)
	got, err := utility.Render(component)
	if err != nil {
		t.Fatal(err)
	}

	diff, err := htmldiff.DiffStrings(utility.MinifyHtml(alert), utility.MinifyHtml(got))
	if err != nil {
		t.Fatal(err)
	}
	if diff != "" {
		t.Error(diff)
	}
}

//go:embed test-data/alert-color.html
var alertColor string

func TestAlertColor(t *testing.T) {
	alertParams := alerts.AlertParams{
		Title:       "title",
		Description: "description",
		Color:       "danger",
		Dismissible: false,
		Important:   false,
	}
	component := alerts.Alert(alertParams)
	got, err := utility.Render(component)
	if err != nil {
		t.Fatal(err)
	}
	diff, err := htmldiff.DiffStrings(utility.MinifyHtml(alertColor), utility.MinifyHtml(got))
	if err != nil {
		t.Fatal(err)
	}
	if diff != "" {
		t.Error(diff)
	}
}

//go:embed test-data/alert-important.html
var alertImportant string

func TestAlertImportant(t *testing.T) {
	alertParams := alerts.AlertParams{
		Title:       "title",
		Description: "description",
		Color:       "",
		Dismissible: false,
		Important:   true,
	}
	component := alerts.Alert(alertParams)
	got, err := utility.Render(component)
	if err != nil {
		t.Fatal(err)
	}
	diff, err := htmldiff.DiffStrings(utility.MinifyHtml(alertImportant), utility.MinifyHtml(got))
	if err != nil {
		t.Fatal(err)
	}
	if diff != "" {
		t.Error(diff)
	}
}

//go:embed test-data/alert-important-dismissible.html
var alertImportantDismissible string

func TestAlertImportantDismissible(t *testing.T) {
	alertParams := alerts.AlertParams{
		Title:       "title",
		Description: "description",
		Color:       "",
		Dismissible: true,
		Important:   true,
	}
	component := alerts.Alert(alertParams)
	got, err := utility.Render(component)
	if err != nil {
		t.Fatal(err)
	}
	diff, err := htmldiff.DiffStrings(utility.MinifyHtml(alertImportantDismissible), utility.MinifyHtml(got))
	if err != nil {
		t.Fatal(err)
	}
	if diff != "" {
		t.Error(diff)
	}
}

//go:embed test-data/alert-icon.html
var alertIcon string

func TestAlertIcon(t *testing.T) {
	alertParams := alerts.AlertParams{
		Title:       "title",
		Description: "description",
		Color:       "",
		Dismissible: false,
		Important:   false,
	}
	component := alerts.AlertWithIcon(alertParams, "2fa")
	got, err := utility.Render(component)
	t.Log(got)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(got)
	t.Log(alertIcon)
	diff, err := htmldiff.DiffStrings(alertIcon, got)
	if err != nil {
		t.Fatal(err)
	}
	if diff != "" {
		t.Error(diff)
	}
}

//go:embed test-data/alert-icon-dismissible.html
var alertIconDismissible string

func TestAlertIconDismissible(t *testing.T) {
	alertParams := alerts.AlertParams{
		Title:       "title",
		Description: "description",
		Color:       "",
		Dismissible: true,
		Important:   false,
	}
	component := alerts.AlertWithIcon(alertParams, "2fa")
	got, err := utility.Render(component)
	if err != nil {
		t.Fatal(err)
	}
	diff, err := htmldiff.DiffStrings(utility.MinifyHtml(alertIconDismissible), utility.MinifyHtml(got))
	if err != nil {
		t.Fatal(err)
	}
	if diff != "" {
		t.Error(diff)
	}
}

//
////go:embed test-data/alert-avatar.html
//var alertAvatar string
//
//func TestAlertAvatar(t *testing.T) {
//	component := tabler.AlertWithAvatar("title", "description", "", false, false)
//	diff, err := htmldiff.Diff(component, alertAvatar)
//	if err != nil {
//		t.Fatal(err)
//	}
//	if diff != "" {
//		t.Error(diff)
//	}
//}
//
////go:embed test-data/alert-avatar-dismissible.html
//var alertAvatarDismissible string
//
//func TestAlertAvatarDismissible(t *testing.T) {
//	component := tabler.Alert("title", "description", "", true, false)
//	diff, err := htmldiff.Diff(component, alertAvatarDismissible)
//	if err != nil {
//		t.Fatal(err)
//	}
//	if diff != "" {
//		t.Error(diff)
//	}
//}
