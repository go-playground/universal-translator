package et

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type et struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'et' locale
func New() locales.Translator {
	return &et{
		locale:  "et",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *et) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'et'
func (t *et) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'et'
func (t *et) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
