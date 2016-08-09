package nmg

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type nmg struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'nmg' locale
func New() locales.Translator {
	return &nmg{
		locale:  "nmg",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *nmg) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'nmg'
func (t *nmg) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'nmg'
func (t *nmg) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
