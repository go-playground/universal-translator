package dsb_DE

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type dsb_DE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'dsb_DE' locale
func New() locales.Translator {
	return &dsb_DE{
		locale:  "dsb_DE",
		plurals: []locales.PluralRule{2, 3, 4, 6},
	}
}

// Locale returns the current translators string locale
func (t *dsb_DE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'dsb_DE'
func (t *dsb_DE) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'dsb_DE'
func (t *dsb_DE) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

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
