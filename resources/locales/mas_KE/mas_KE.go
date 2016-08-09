package mas_KE

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type mas_KE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'mas_KE' locale
func New() locales.Translator {
	return &mas_KE{
		locale:  "mas_KE",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *mas_KE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mas_KE'
func (t *mas_KE) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'mas_KE'
func (t *mas_KE) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
