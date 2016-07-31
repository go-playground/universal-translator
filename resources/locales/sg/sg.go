package sg

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type sg struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'sg' locale
func New() locales.Translator {
	return &sg{
		locale:  "sg",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *sg) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sg'
func (t *sg) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'sg'
func (t *sg) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
