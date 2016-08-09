package ii

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ii struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ii' locale
func New() locales.Translator {
	return &ii{
		locale:  "ii",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *ii) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ii'
func (t *ii) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ii'
func (t *ii) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
