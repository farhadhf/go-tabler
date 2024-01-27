package avatars_test

import (
	_ "embed"
	"testing"

	"github.com/a-h/templ"
	tabler "github.com/farhadhf/go-tabler/components"
	avatars "github.com/farhadhf/go-tabler/components/avatars"
	"github.com/farhadhf/go-tabler/internal/utility"
)

//go:embed test-data/avatar-initials.html
var avatar string

func TestInitialsAvatar(t *testing.T) {
	component := avatars.AvatarWithInitials("FH", avatars.AvatarParams{})

	err := utility.CompareComponentAndHtml(component, avatar)
	if err != nil {
		t.Error(err)
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
	err := utility.CompareComponentAndHtml(component, avatarLayout)
	if err != nil {
		t.Error(err)
	}
}

//go:embed test-data/avatar-icon.html
var avatarWithIcon string

func TestAvatarWithIcon(t *testing.T) {
	component := avatars.AvatarWithIcon("2fa", avatars.AvatarParams{})
	err := utility.CompareComponentAndHtml(component, avatarWithIcon)
	if err != nil {
		t.Error(err)
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
	err := utility.CompareComponentAndHtml(component, avatarList)
	if err != nil {
		t.Error(err)
	}

	component = avatars.AvatarList(avatarComponents, true)
	err = utility.CompareComponentAndHtml(component, avatarStackedList)
	if err != nil {
		t.Error(err)
	}
}
