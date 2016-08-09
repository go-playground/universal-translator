package so_SO

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type so_SO struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'so_SO' locale
func New() locales.Translator {
	return &so_SO{
		locale:  "so_SO",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *so_SO) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'so_SO'
func (t *so_SO) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'so_SO'
func (t *so_SO) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
