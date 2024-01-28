package forms_test

import (
	_ "embed"
	"testing"

	"github.com/a-h/templ"
	tabler "github.com/farhadhf/go-tabler/components"
	"github.com/farhadhf/go-tabler/components/forms"
	"github.com/farhadhf/go-tabler/internal/utility"
)

//go:embed test-data/basic-input-element.html
var expectedBasicInput string

func TestBasicInputElement(t *testing.T) {
	opts := forms.InputElementOpts{
		InputType:   "email",
		Name:        "email",
		Placeholder: "someone@example.com",
		Value:       "farhad@example.com",
	}

	component := forms.InputElement(opts)
	err := utility.CompareComponentAndHtml(component, expectedBasicInput)
	if err != nil {
		t.Error(err)
	}
}

//go:embed test-data/label-element.html
var expectedLabel string

func TestLabelElement(t *testing.T) {
	opts := forms.LabelElementOpts{
		Label:      tabler.Text("Email"),
		Attributes: templ.Attributes{"for": "email"},
		Classes:    "mb-2",
	}
	component := forms.LabelElement(opts)
	err := utility.CompareComponentAndHtml(component, expectedLabel)
	if err != nil {
		t.Error(err)
	}
}

//go:embed test-data/form-input.html
var expectedFormInput string

func TestFormInput(t *testing.T) {
	opts := forms.FormInputOpts{
		WrapperAttributes: templ.Attributes{"id": "name-wrapper"},
		WrapperClasses:    "wrapper name",
		InputOptions: forms.InputElementOpts{
			Name:        "name",
			Placeholder: "Bilbo Baggins",
			Value:       "Farhad Hedayatifard",
			Rounded:     true,
			Flush:       true,
			Size:        tabler.SizeMD,
			Attributes:  templ.Attributes{"id": "name"},
			Classes:     "name-input",
		},
		LabelOptions: forms.LabelElementOpts{
			Label:      tabler.Text("Name"),
			For:        "name",
			Attributes: templ.Attributes{"id": "name-label"},
			Classes:    "name-label",
		},
	}
	component := forms.FormInput(opts)
	err := utility.CompareComponentAndHtml(component, expectedFormInput)
	if err != nil {
		t.Error(err)
	}
}

//go:embed test-data/form-input-icon.html
var expectedFormInputIcon string

func TestFormInputWithIcon(t *testing.T) {
	opts := forms.FormInputOpts{
		WrapperAttributes: templ.Attributes{"id": "amount-wrapper"},
		WrapperClasses:    "wrapper amount",
		InputOptions: forms.InputElementOpts{
			Name:        "amount",
			Placeholder: "Amount",
			Value:       "100",
			Rounded:     true,
			Flush:       true,
			Size:        tabler.SizeMD,
			Attributes:  templ.Attributes{"id": "amount"},
			Classes:     "amount-input",
		},
	}

	component := forms.FormInputWithIcon(opts, forms.InputIconAddonOpts{
		StartAddonIcon: "currency-dollar",
		EndAddonIcon:   "plus",
	})
	err := utility.CompareComponentAndHtml(component, expectedFormInputIcon)
	if err != nil {
		t.Error(err)
	}
}

//go:embed test-data/form-input-loader.html
var expectedFormInputLoader string

func TestFormInputWithLoader(t *testing.T) {
	opts := forms.FormInputOpts{
		WrapperAttributes: templ.Attributes{"id": "amount-wrapper"},
		WrapperClasses:    "wrapper amount",
		InputOptions: forms.InputElementOpts{
			Name:        "amount",
			Placeholder: "Amount",
			Value:       "100",
			Rounded:     true,
			Flush:       true,
			Size:        tabler.SizeMD,
			Attributes:  templ.Attributes{"id": "amount"},
			Classes:     "amount-input",
		},
	}

	component := forms.FormInputWithLoader(opts, true)
	err := utility.CompareComponentAndHtml(component, expectedFormInputLoader)
	if err != nil {
		t.Error(err)
	}
}

//go:embed test-data/full-input-group.html
var expectedFormInputGroup string

func TestFormInputGroup(t *testing.T) {
	opts := forms.FormInputOpts{
		WrapperAttributes: templ.Attributes{"id": "amount-wrapper"},
		WrapperClasses:    "wrapper amount",
		InputOptions: forms.InputElementOpts{
			Name:        "amount",
			Placeholder: "Amount",
			Value:       "100",
			Rounded:     true,
			Flush:       true,
			Size:        tabler.SizeMD,
			Attributes:  templ.Attributes{"id": "amount"},
			Classes:     "amount-input",
		},
	}

	component := forms.FormInputGroup(opts, forms.InputGroupOpts{
		Start: tabler.Text("$"),
		End:   tabler.Text("USD"),
	})
	err := utility.CompareComponentAndHtml(component, expectedFormInputGroup)
	if err != nil {
		t.Error(err)
	}
}
