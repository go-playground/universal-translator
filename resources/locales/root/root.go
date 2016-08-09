package root

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type root struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'root' locale
func New() locales.Translator {
	return &root{
		locale:  "root",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *root) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'root'
func (t *root) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'root'
func (t *root) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
