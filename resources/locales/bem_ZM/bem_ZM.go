package bem_ZM

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type bem_ZM struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'bem_ZM' locale
func New() locales.Translator {
	return &bem_ZM{
		locale:  "bem_ZM",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *bem_ZM) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'bem_ZM'
func (t *bem_ZM) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'bem_ZM'
func (t *bem_ZM) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
