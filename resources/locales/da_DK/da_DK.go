package da_DK

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type da_DK struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'da_DK' locale
func New() locales.Translator {
	return &da_DK{
		locale:  "da_DK",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *da_DK) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'da_DK'
func (t *da_DK) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'da_DK'
func (t *da_DK) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	t := locales.T(n, v)

	if (n == 1) || (t != 0 && (i == 0 || i == 1)) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
