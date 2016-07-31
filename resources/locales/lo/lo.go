package lo

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type lo struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'lo' locale
func New() locales.Translator {
	return &lo{
		locale:  "lo",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *lo) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'lo'
func (t *lo) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'lo'
func (t *lo) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
