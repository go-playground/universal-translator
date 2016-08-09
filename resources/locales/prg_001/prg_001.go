package prg_001

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type prg_001 struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'prg_001' locale
func New() locales.Translator {
	return &prg_001{
		locale:  "prg_001",
		plurals: []locales.PluralRule{1, 2, 6},
	}
}

// Locale returns the current translators string locale
func (t *prg_001) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'prg_001'
func (t *prg_001) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'prg_001'
func (t *prg_001) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	f := locales.F(n, v)

	if (n%10 == 0) || (n%100 >= 11 && n%100 <= 19) || (v == 2 && f%100 >= 11 && f%100 <= 19) {
		return locales.PluralRuleZero
	} else if (n%10 == 1 && n%100 != 11) || (v == 2 && f%10 == 1 && f%100 != 11) || (v != 2 && f%10 == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
