package sv_SE

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type sv_SE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'sv_SE' locale
func New() locales.Translator {
	return &sv_SE{
		locale:  "sv_SE",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *sv_SE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sv_SE'
func (t *sv_SE) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'sv_SE'
func (t *sv_SE) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
