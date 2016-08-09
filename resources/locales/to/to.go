package to

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type to struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'to' locale
func New() locales.Translator {
	return &to{
		locale:  "to",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *to) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'to'
func (t *to) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'to'
func (t *to) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
