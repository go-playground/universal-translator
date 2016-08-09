package gv_IM

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type gv_IM struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'gv_IM' locale
func New() locales.Translator {
	return &gv_IM{
		locale:  "gv_IM",
		plurals: []locales.PluralRule{2, 3, 4, 5, 6},
	}
}

// Locale returns the current translators string locale
func (t *gv_IM) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'gv_IM'
func (t *gv_IM) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'gv_IM'
func (t *gv_IM) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if v == 0 && i%10 == 1 {
		return locales.PluralRuleOne
	} else if v == 0 && i%10 == 2 {
		return locales.PluralRuleTwo
	} else if v == 0 && (i%100 == 0 || i%100 == 20 || i%100 == 40 || i%100 == 60 || i%100 == 80) {
		return locales.PluralRuleFew
	} else if v != 0 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}
