package lu

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type lu struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'lu' locale
func New() locales.Translator {
	return &lu{
		locale:  "lu",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *lu) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'lu'
func (t *lu) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'lu'
func (t *lu) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
