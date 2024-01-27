package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/svg"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var m = minify.New()

func loadIcon(name string) string {
	icon, err := os.ReadFile("tabler-icons/icons/" + name)
	if err != nil {
		log.Fatal(err)
	}
	return string(icon)
}

func minifySvg(filename string) string {
	fileContents := loadIcon(filename)
	fileContents, _ = m.String("image/svg+xml", fileContents)
	return fileContents
}

func capitalizeName(name string) string {
	parts := strings.Split(name, "-")
	result := ""
	for _, part := range parts {
		result += cases.Title(language.English, cases.Compact).String(part)
	}
	return result
}

func main() {
	m.AddFunc("image/svg+xml", svg.Minify)

	files, err := os.ReadDir("tabler-icons/icons")
	if err != nil {
		log.Fatal(err)
	}

	icons := make(map[string]string)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		iconName := strings.Replace(file.Name(), ".svg", "", 1)
		minifiedIcon := strings.TrimSpace(loadIcon(file.Name()))
		icons[iconName] = minifiedIcon
	}

	os.Remove("components/icons.templ")
	f, err := os.OpenFile("components/icons.templ", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.Write([]byte("package tabler\n\n"))
	f.Write([]byte("var icons map[string]string\n"))
	f.Write([]byte("func init() {\n"))
	f.Write([]byte("\ticons = map[string]string{\n"))
	re := regexp.MustCompile(`>\s*<`)
	for iconName, iconComponent := range icons {
		icon := re.ReplaceAll([]byte(strings.ReplaceAll(iconComponent, "\"", "\\\"")), []byte("><"))
		f.Write([]byte(fmt.Sprintf("\t\t\"%s\": \"%s\",\n", iconName, icon)))
	}
	f.Write([]byte(fmt.Sprintf("\t}\n}\n\n")))
	f.Write([]byte(fmt.Sprintf("templ Icon(name string) {\n")))
	f.Write([]byte(fmt.Sprintf("\tif icons[name] != \"\" {\n")))
	f.Write([]byte(fmt.Sprintf("\t\t{ icons[name] }\n")))
	f.Write([]byte(fmt.Sprintf("\t} else {\n")))
	f.Write([]byte(fmt.Sprintf("\t\tINVALID ICON NAME\n")))
	f.Write([]byte(fmt.Sprintf("\t}\n")))
	f.Write([]byte(fmt.Sprintf("}\n\n")))
}
