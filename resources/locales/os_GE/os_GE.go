package os_GE

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type os_GE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'os_GE' locale
func New() locales.Translator {
	return &os_GE{
		locale:  "os_GE",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *os_GE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'os_GE'
func (t *os_GE) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'os_GE'
func (t *os_GE) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
