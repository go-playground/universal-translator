package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"go/format"

	"gopkg.in/yaml.v2"
)

type pluralInfo struct {
	path      string
	locale    string
	plural    string
	outputDir string
}

// rules downloaded/copied from download github.com/vube/i18n/data/rules
func main() {
	rules := "data/rules"
	plurals := []pluralInfo{}
	basePlurals := map[string]string{}

	err := filepath.Walk(rules, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			panic(err)
		}

		if info.IsDir() {
			return nil
		}

		in, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		var out yaml.MapSlice
		if err = yaml.Unmarshal(in, &out); err != nil {
			panic(err)
		}

		var plural string
		for _, item := range out {
			if item.Key == "plural" {
				plural = item.Value.(string)
			}
		}

		locale := strings.Replace(info.Name(), filepath.Ext(info.Name()), "", 1)
		locale = strings.Replace(locale, "-", "_", -1)

		plurals = append(plurals, pluralInfo{
			path:      path,
			locale:    locale,
			plural:    plural,
			outputDir: "../../resources/locales/" + locale,
		})

		if plural == "" {
			return nil
		}

		basePlurals[locale] = plural

		return nil
	})

	if err != nil {
		panic(err)
	}

	for _, p := range plurals {
		if p.plural == "" {

			var ok bool

			fmt.Print("can't find plurals in ", p.path, " attempting to locate base language plural rules...")

			base := strings.SplitN(p.locale, "_", 2)

			p.plural, ok = basePlurals[base[0]]
			if !ok {
				fmt.Println("Not Found")
				continue
			}

			fmt.Println("Found")
		}

		if _, err := os.Stat(p.outputDir); err != nil {
			fmt.Println(p.outputDir, "doesn't exist")
			continue
		}
		file, err := os.Create(p.outputDir + "/plural.go")
		if err != nil {
			panic(err)
		}
		rawcodes := fmt.Sprintf(`package %s

		const pluralRule = %q
		`, p.locale, p.plural)

		var codes []byte
		if codes, err = format.Source([]byte(rawcodes)); err != nil {
			panic(err)
		}
		if _, err := file.Write(codes); err != nil {
			panic(err)
		}
	}
}
