package sl

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type sl struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'sl' locale
func New() locales.Translator {
	return &sl{
		locale:  "sl",
		plurals: []locales.PluralRule{2, 3, 4, 6},
	}
}

// Locale returns the current translators string locale
func (t *sl) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sl'
func (t *sl) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'sl'
func (t *sl) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if v == 0 && i%100 == 1 {
		return locales.PluralRuleOne
	} else if v == 0 && i%100 == 2 {
		return locales.PluralRuleTwo
	} else if (v == 0 && i%100 >= 3 && i%100 <= 4) || (v != 0) {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}
