package root

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type root struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'root' locale
func New() locales.Translator {
	return &root{
		locale:  "root",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *root) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'root'
func (t *root) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'root'
func (t *root) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
