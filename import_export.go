package ut

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"io"

	"github.com/go-playground/locales"
)

type translation struct {
	Locale           string      `json:"locale"`
	Key              interface{} `json:"key"` // either string or integer
	Translation      string      `json:"trans"`
	PluralType       string      `json:"type,omitempty"`
	PluralRule       string      `json:"rule,omitempty"`
	OverrideExisting bool        `json:"override,omitempty"`
}

const (
	cardinalType = "Cardinal"
	ordinalType  = "Ordinal"
	rangeType    = "Range"
)

// ExportFormat is the format of the file import or export
type ExportFormat uint8

// supported Export Formats
const (
	JSON ExportFormat = iota
)

// Export writes the translations out to a file on disk.
//
// NOTE: this currently only works with string or int translations keys.
func (t *UniversalTranslator) Export(format ExportFormat, filename string) error {

	// build up translations
	var trans []translation

	for _, locale := range t.translators {

		for k, v := range locale.(*translator).translations {
			trans = append(trans, translation{
				Locale:      locale.Locale(),
				Key:         k,
				Translation: v.text,
			})
		}

		for k, pluralTrans := range locale.(*translator).cardinalTanslations {

			for i, plural := range pluralTrans {

				// leave enough for all plural rules
				// but not all are set for all languages.
				if plural == nil {
					continue
				}

				trans = append(trans, translation{
					Locale:      locale.Locale(),
					Key:         k.(string),
					Translation: plural.text,
					PluralType:  cardinalType,
					PluralRule:  locales.PluralRule(i).String(),
				})
			}
		}

		for k, pluralTrans := range locale.(*translator).ordinalTanslations {

			for i, plural := range pluralTrans {

				// leave enough for all plural rules
				// but not all are set for all languages.
				if plural == nil {
					continue
				}

				trans = append(trans, translation{
					Locale:      locale.Locale(),
					Key:         k.(string),
					Translation: plural.text,
					PluralType:  ordinalType,
					PluralRule:  locales.PluralRule(i).String(),
				})
			}
		}

		for k, pluralTrans := range locale.(*translator).rangeTanslations {

			for i, plural := range pluralTrans {

				// leave enough for all plural rules
				// but not all are set for all languages.
				if plural == nil {
					continue
				}

				trans = append(trans, translation{
					Locale:      locale.Locale(),
					Key:         k.(string),
					Translation: plural.text,
					PluralType:  rangeType,
					PluralRule:  locales.PluralRule(i).String(),
				})
			}
		}
	}

	var b []byte
	var err error

	switch format {
	case JSON:
		b, err = json.MarshalIndent(trans, "", "    ")
	}

	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, b, 0644)
}

// Import reads the translations out of a file or directory on disk.
//
// NOTE: this currently only works with string or int translations keys.
func (t *UniversalTranslator) Import(format ExportFormat, dirnameOrFilename string) error {

	fi, err := os.Stat(dirnameOrFilename)
	if err != nil {
		return err
	}

	processFn := func(filename string) error {

		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()

		return t.ImportByReader(format, f)
	}

	if !fi.IsDir() {
		return processFn(dirnameOrFilename)
	}

	// recursively go through directory
	walker := func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			return nil
		}

		switch format {
		case JSON:
			// skip non JSON files
			if filepath.Ext(info.Name()) != ".json" {
				return nil
			}
		}

		return processFn(path)
	}

	return filepath.Walk(dirnameOrFilename, walker)
}

// ImportByReader imports the the translations found within the contents read from the supplied reader.
//
// NOTE: generally used when assets have been embedded into the binary and are already in memory.
func (t *UniversalTranslator) ImportByReader(format ExportFormat, reader io.Reader) error {

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	var trans []translation

	switch format {
	case JSON:
		err = json.Unmarshal(b, &trans)
	}

	if err != nil {
		return err
	}

	for _, tl := range trans {

		locale, found := t.FindTranslator(tl.Locale)
		if !found {
			return &ErrMissingLocale{locale: tl.Locale}
		}

		pr := stringToPR(tl.PluralRule)

		if pr == locales.PluralRuleUnknown {

			err = locale.Add(tl.Key, tl.Translation, tl.OverrideExisting)
			if err != nil {
				return err
			}

			continue
		}

		switch tl.PluralType {
		case cardinalType:
			err = locale.AddCardinal(tl.Key, tl.Translation, pr, tl.OverrideExisting)
		case ordinalType:
			err = locale.AddOrdinal(tl.Key, tl.Translation, pr, tl.OverrideExisting)
		case rangeType:
			err = locale.AddRange(tl.Key, tl.Translation, pr, tl.OverrideExisting)
		default:
			return &ErrBadPluralDefinition{tl: tl}
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func stringToPR(s string) locales.PluralRule {

	switch s {
	case "One":
		return locales.PluralRuleOne
	case "Two":
		return locales.PluralRuleTwo
	case "Few":
		return locales.PluralRuleFew
	case "Many":
		return locales.PluralRuleMany
	case "Other":
		return locales.PluralRuleOther
	default:
		return locales.PluralRuleUnknown
	}

}
