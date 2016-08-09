package et_EE

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type et_EE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'et_EE' locale
func New() locales.Translator {
	return &et_EE{
		locale:  "et_EE",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *et_EE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'et_EE'
func (t *et_EE) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'et_EE'
func (t *et_EE) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
