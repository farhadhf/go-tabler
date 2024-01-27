package avatars_test

import (
	_ "embed"
	"testing"

	"github.com/a-h/templ"
	"github.com/a-h/templ/generator/htmldiff"
	tabler "github.com/farhadhf/go-tabler/components"
	avatars "github.com/farhadhf/go-tabler/components/avatars"
	"github.com/farhadhf/go-tabler/internal/utility"
)

//go:embed test-data/avatar-initials.html
var avatar string

func TestInitialsAvatar(t *testing.T) {
	component := avatars.AvatarWithInitials("FH", avatars.AvatarParams{})
	got, err := utility.Render(component)
	if err != nil {
		t.Fatal(err)
	}

	diff, err := htmldiff.DiffStrings(utility.MinifyHtml(avatar), utility.MinifyHtml(got))
	if err != nil {
		t.Fatal(err)
	}
	if diff != "" {
		t.Error(diff)
	}
}

//go:embed test-data/avatar-layout.html
var avatarLayout string

func TestAvatarLayout(t *testing.T) {
	params := avatars.AvatarParams{
		BackgroundColor: "red-lt",
		Size:            tabler.SizeSM,
		Shape:           avatars.Circle,
	}
	component := avatars.AvatarWithInitials("FH", params)
	got, err := utility.Render(component)
	if err != nil {
		t.Fatal(err)
	}

	diff, err := htmldiff.DiffStrings(utility.MinifyHtml(avatarLayout), utility.MinifyHtml(got))
	if err != nil {
		t.Fatal(err)
	}
	if diff != "" {
		t.Error(diff)
	}
}

//go:embed test-data/avatar-icon.html
var avatarWithIcon string

func TestAvatarWithIcon(t *testing.T) {
	component := avatars.AvatarWithIcon("2fa", avatars.AvatarParams{})
	got, err := utility.Render(component)
	if err != nil {
		t.Fatal(err)
	}

	diff, err := htmldiff.DiffStrings(utility.MinifyHtml(avatarWithIcon), utility.MinifyHtml(got))
	if err != nil {
		t.Fatal(err)
	}
	if diff != "" {
		t.Error(diff)
	}
}

//go:embed test-data/avatar-list.html
var avatarList string

//go:embed test-data/avatar-stacked-list.html
var avatarStackedList string

func TestAvatarList(t *testing.T) {
	avatarComponents := []templ.Component{
		avatars.AvatarWithInitials(
			"FH",
			avatars.AvatarParams{
				BackgroundColor: "red-lt",
				Size:            tabler.SizeSM,
				Shape:           avatars.Circle,
			},
		),
		avatars.AvatarWithIcon(
			"2fa",
			avatars.AvatarParams{
				BackgroundColor: "red-lt",
				Size:            tabler.SizeSM,
				Shape:           avatars.Circle,
			},
		),
	}
	component := avatars.AvatarList(avatarComponents, false)
	got, err := utility.Render(component)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(got)
	t.Log(avatarList)
	diff, err := htmldiff.DiffStrings(avatarList, got)
	if err != nil {
		t.Fatal(err)
	}
	if diff != "" {
		t.Error(diff)
	}

	component = avatars.AvatarList(avatarComponents, true)
	got, err = utility.Render(component)
	if err != nil {
		t.Fatal(err)
	}

	diff, err = htmldiff.DiffStrings(utility.MinifyHtml(avatarStackedList), utility.MinifyHtml(got))
	if err != nil {
		t.Fatal(err)
	}
	if diff != "" {
		t.Error(diff)
	}
}
