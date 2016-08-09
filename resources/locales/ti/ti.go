package ti

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ti struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ti' locale
func New() locales.Translator {
	return &ti{
		locale:  "ti",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ti) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ti'
func (t *ti) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ti'
func (t *ti) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n >= 0 && n <= 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
