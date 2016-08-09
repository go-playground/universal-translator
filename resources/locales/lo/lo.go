package lo

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type lo struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'lo' locale
func New() locales.Translator {
	return &lo{
		locale:  "lo",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *lo) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'lo'
func (t *lo) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'lo'
func (t *lo) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
