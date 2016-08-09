package bg_BG

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type bg_BG struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'bg_BG' locale
func New() locales.Translator {
	return &bg_BG{
		locale:  "bg_BG",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *bg_BG) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'bg_BG'
func (t *bg_BG) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'bg_BG'
func (t *bg_BG) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
