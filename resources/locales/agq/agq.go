package agq

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type agq struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'agq' locale
func New() locales.Translator {
	return &agq{
		locale:  "agq",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *agq) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'agq'
func (t *agq) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'agq'
func (t *agq) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
