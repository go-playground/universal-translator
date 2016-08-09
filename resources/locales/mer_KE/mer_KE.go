package mer_KE

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type mer_KE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'mer_KE' locale
func New() locales.Translator {
	return &mer_KE{
		locale:  "mer_KE",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *mer_KE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mer_KE'
func (t *mer_KE) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'mer_KE'
func (t *mer_KE) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
