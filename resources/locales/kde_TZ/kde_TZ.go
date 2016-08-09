package kde_TZ

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type kde_TZ struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'kde_TZ' locale
func New() locales.Translator {
	return &kde_TZ{
		locale:  "kde_TZ",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *kde_TZ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'kde_TZ'
func (t *kde_TZ) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'kde_TZ'
func (t *kde_TZ) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
