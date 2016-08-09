package fo

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type fo struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'fo' locale
func New() locales.Translator {
	return &fo{
		locale:  "fo",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *fo) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'fo'
func (t *fo) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'fo'
func (t *fo) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
