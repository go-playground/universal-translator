package nmg_CM

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type nmg_CM struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'nmg_CM' locale
func New() locales.Translator {
	return &nmg_CM{
		locale:  "nmg_CM",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *nmg_CM) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'nmg_CM'
func (t *nmg_CM) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'nmg_CM'
func (t *nmg_CM) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
