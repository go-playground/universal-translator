package zu_ZA

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type zu_ZA struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'zu_ZA' locale
func New() locales.Translator {
	return &zu_ZA{
		locale:  "zu_ZA",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *zu_ZA) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'zu_ZA'
func (t *zu_ZA) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'zu_ZA'
func (t *zu_ZA) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
