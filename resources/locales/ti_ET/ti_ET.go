package ti_ET

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ti_ET struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ti_ET' locale
func New() locales.Translator {
	return &ti_ET{
		locale:  "ti_ET",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ti_ET) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ti_ET'
func (t *ti_ET) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ti_ET'
func (t *ti_ET) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n >= 0 && n <= 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
