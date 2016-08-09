package guz

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type guz struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'guz' locale
func New() locales.Translator {
	return &guz{
		locale:  "guz",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *guz) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'guz'
func (t *guz) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'guz'
func (t *guz) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
