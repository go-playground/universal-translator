package sr_Cyrl_BA

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type sr_Cyrl_BA struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'sr_Cyrl_BA' locale
func New() locales.Translator {
	return &sr_Cyrl_BA{
		locale:  "sr_Cyrl_BA",
		plurals: []locales.PluralRule{2, 4, 6},
	}
}

// Locale returns the current translators string locale
func (t *sr_Cyrl_BA) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sr_Cyrl_BA'
func (t *sr_Cyrl_BA) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'sr_Cyrl_BA'
func (t *sr_Cyrl_BA) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	f := locales.F(n, v)

	if (v == 0 && i%10 == 1 && i%100 != 11) || (f%10 == 1 && f%100 != 11) {
		return locales.PluralRuleOne
	} else if (v == 0 && i%10 >= 2 && i%10 <= 4 && i%100 < 12 && i%100 > 14) || (f%10 >= 2 && f%10 <= 4 && f%100 < 12 && f%100 > 14) {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}
