package vai

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type vai struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'vai' locale
func New() locales.Translator {
	return &vai{
		locale:  "vai",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *vai) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'vai'
func (t *vai) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'vai'
func (t *vai) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
