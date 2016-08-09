package cy

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type cy struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'cy' locale
func New() locales.Translator {
	return &cy{
		locale:  "cy",
		plurals: []locales.PluralRule{1, 2, 3, 4, 5, 6},
	}
}

// Locale returns the current translators string locale
func (t *cy) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'cy'
func (t *cy) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'cy'
func (t *cy) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 0 {
		return locales.PluralRuleZero
	} else if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	} else if n == 3 {
		return locales.PluralRuleFew
	} else if n == 6 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}
