package to_TO

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type to_TO struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'to_TO' locale
func New() locales.Translator {
	return &to_TO{
		locale:  "to_TO",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *to_TO) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'to_TO'
func (t *to_TO) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'to_TO'
func (t *to_TO) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
