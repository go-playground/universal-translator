package km_KH

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type km_KH struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'km_KH' locale
func New() locales.Translator {
	return &km_KH{
		locale:  "km_KH",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *km_KH) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'km_KH'
func (t *km_KH) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'km_KH'
func (t *km_KH) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
