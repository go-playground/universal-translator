package ff_CM

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ff_CM struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ff_CM' locale
func New() locales.Translator {
	return &ff_CM{
		locale:  "ff_CM",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ff_CM) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ff_CM'
func (t *ff_CM) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ff_CM'
func (t *ff_CM) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 0 || i == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
