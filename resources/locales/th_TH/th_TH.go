package th_TH

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type th_TH struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'th_TH' locale
func New() locales.Translator {
	return &th_TH{
		locale:  "th_TH",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *th_TH) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'th_TH'
func (t *th_TH) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'th_TH'
func (t *th_TH) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
