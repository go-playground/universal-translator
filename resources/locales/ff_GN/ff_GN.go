package ff_GN

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ff_GN struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ff_GN' locale
func New() locales.Translator {
	return &ff_GN{
		locale:  "ff_GN",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ff_GN) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ff_GN'
func (t *ff_GN) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ff_GN'
func (t *ff_GN) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 0 || i == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
