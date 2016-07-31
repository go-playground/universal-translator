package kde

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type kde struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'kde' locale
func New() locales.Translator {
	return &kde{
		locale:  "kde",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *kde) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'kde'
func (t *kde) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'kde'
func (t *kde) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
