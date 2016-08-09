package ki_KE

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ki_KE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ki_KE' locale
func New() locales.Translator {
	return &ki_KE{
		locale:  "ki_KE",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *ki_KE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ki_KE'
func (t *ki_KE) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ki_KE'
func (t *ki_KE) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
