package mt_MT

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type mt_MT struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'mt_MT' locale
func New() locales.Translator {
	return &mt_MT{
		locale:  "mt_MT",
		plurals: []locales.PluralRule{2, 4, 5, 6},
	}
}

// Locale returns the current translators string locale
func (t *mt_MT) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mt_MT'
func (t *mt_MT) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'mt_MT'
func (t *mt_MT) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	} else if (n == 0) || (n%100 >= 2 && n%100 <= 10) {
		return locales.PluralRuleFew
	} else if n%100 >= 11 && n%100 <= 19 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}
