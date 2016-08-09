package lag_TZ

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type lag_TZ struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'lag_TZ' locale
func New() locales.Translator {
	return &lag_TZ{
		locale:  "lag_TZ",
		plurals: []locales.PluralRule{1, 2, 6},
	}
}

// Locale returns the current translators string locale
func (t *lag_TZ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'lag_TZ'
func (t *lag_TZ) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'lag_TZ'
func (t *lag_TZ) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if n == 0 {
		return locales.PluralRuleZero
	} else if (i == 0 || i == 1) && n != 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
