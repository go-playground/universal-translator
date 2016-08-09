package bas_CM

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type bas_CM struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'bas_CM' locale
func New() locales.Translator {
	return &bas_CM{
		locale:  "bas_CM",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *bas_CM) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'bas_CM'
func (t *bas_CM) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'bas_CM'
func (t *bas_CM) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
