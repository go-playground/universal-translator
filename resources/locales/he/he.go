package he

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type he struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'he' locale
func New() locales.Translator {
	return &he{
		locale:  "he",
		plurals: []locales.PluralRule{2, 3, 5, 6},
	}
}

// Locale returns the current translators string locale
func (t *he) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'he'
func (t *he) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'he'
func (t *he) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	} else if i == 2 && v == 0 {
		return locales.PluralRuleTwo
	} else if v == 0 && n < 0 && n > 10 && n%10 == 0 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}
