package ka

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ka struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ka' locale
func New() locales.Translator {
	return &ka{
		locale:  "ka",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ka) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ka'
func (t *ka) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ka'
func (t *ka) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
