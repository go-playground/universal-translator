package agq_CM

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type agq_CM struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'agq_CM' locale
func New() locales.Translator {
	return &agq_CM{
		locale:  "agq_CM",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *agq_CM) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'agq_CM'
func (t *agq_CM) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'agq_CM'
func (t *agq_CM) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
