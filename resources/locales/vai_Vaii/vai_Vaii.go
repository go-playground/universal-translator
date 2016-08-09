package vai_Vaii

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type vai_Vaii struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'vai_Vaii' locale
func New() locales.Translator {
	return &vai_Vaii{
		locale:  "vai_Vaii",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *vai_Vaii) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'vai_Vaii'
func (t *vai_Vaii) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'vai_Vaii'
func (t *vai_Vaii) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
