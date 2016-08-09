package ak

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ak struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ak' locale
func New() locales.Translator {
	return &ak{
		locale:  "ak",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ak) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ak'
func (t *ak) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ak'
func (t *ak) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n >= 0 && n <= 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
