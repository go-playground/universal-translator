package ak_GH

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ak_GH struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ak_GH' locale
func New() locales.Translator {
	return &ak_GH{
		locale:  "ak_GH",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ak_GH) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ak_GH'
func (t *ak_GH) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ak_GH'
func (t *ak_GH) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n >= 0 && n <= 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
