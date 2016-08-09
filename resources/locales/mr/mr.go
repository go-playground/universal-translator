package mr

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type mr struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'mr' locale
func New() locales.Translator {
	return &mr{
		locale:  "mr",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *mr) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mr'
func (t *mr) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'mr'
func (t *mr) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
