package mas

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type mas struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'mas' locale
func New() locales.Translator {
	return &mas{
		locale:  "mas",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *mas) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mas'
func (t *mas) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'mas'
func (t *mas) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
