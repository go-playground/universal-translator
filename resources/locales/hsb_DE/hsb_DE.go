package hsb_DE

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type hsb_DE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'hsb_DE' locale
func New() locales.Translator {
	return &hsb_DE{
		locale:  "hsb_DE",
		plurals: []locales.PluralRule{2, 3, 4, 6},
	}
}

// Locale returns the current translators string locale
func (t *hsb_DE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'hsb_DE'
func (t *hsb_DE) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'hsb_DE'
func (t *hsb_DE) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	f := locales.F(n, v)

	if (v == 0 && i%100 == 1) || (f%100 == 1) {
		return locales.PluralRuleOne
	} else if (v == 0 && i%100 == 2) || (f%100 == 2) {
		return locales.PluralRuleTwo
	} else if (v == 0 && i%100 >= 3 && i%100 <= 4) || (f%100 >= 3 && f%100 <= 4) {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}
