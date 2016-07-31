package th

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type th struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'th' locale
func New() locales.Translator {
	return &th{
		locale:  "th",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *th) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'th'
func (t *th) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'th'
func (t *th) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
