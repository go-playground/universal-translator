package pt_CH

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type pt_CH struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'pt_CH' locale
func New() locales.Translator {
	return &pt_CH{
		locale:  "pt_CH",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *pt_CH) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'pt_CH'
func (t *pt_CH) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'pt_CH'
func (t *pt_CH) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n >= 0 && n <= 2 && n != 2 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
