package bm

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type bm struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'bm' locale
func New() locales.Translator {
	return &bm{
		locale:  "bm",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *bm) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'bm'
func (t *bm) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'bm'
func (t *bm) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
