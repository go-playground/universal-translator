package nd

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type nd struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'nd' locale
func New() locales.Translator {
	return &nd{
		locale:  "nd",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *nd) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'nd'
func (t *nd) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'nd'
func (t *nd) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
