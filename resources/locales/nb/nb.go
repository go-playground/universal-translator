package nb

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type nb struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'nb' locale
func New() locales.Translator {
	return &nb{
		locale:  "nb",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *nb) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'nb'
func (t *nb) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'nb'
func (t *nb) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
