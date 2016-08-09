package bem

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type bem struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'bem' locale
func New() locales.Translator {
	return &bem{
		locale:  "bem",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *bem) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'bem'
func (t *bem) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'bem'
func (t *bem) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
