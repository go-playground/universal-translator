package ja_JP

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ja_JP struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ja_JP' locale
func New() locales.Translator {
	return &ja_JP{
		locale:  "ja_JP",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *ja_JP) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ja_JP'
func (t *ja_JP) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ja_JP'
func (t *ja_JP) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
