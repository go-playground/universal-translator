package ii

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ii struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ii' locale
func New() locales.Translator {
	return &ii{
		locale:  "ii",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *ii) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ii'
func (t *ii) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ii'
func (t *ii) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
