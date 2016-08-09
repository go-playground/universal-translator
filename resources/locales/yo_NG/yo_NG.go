package yo_NG

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type yo_NG struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'yo_NG' locale
func New() locales.Translator {
	return &yo_NG{
		locale:  "yo_NG",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *yo_NG) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'yo_NG'
func (t *yo_NG) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'yo_NG'
func (t *yo_NG) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
