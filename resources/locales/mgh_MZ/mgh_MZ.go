package mgh_MZ

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type mgh_MZ struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'mgh_MZ' locale
func New() locales.Translator {
	return &mgh_MZ{
		locale:  "mgh_MZ",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *mgh_MZ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mgh_MZ'
func (t *mgh_MZ) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'mgh_MZ'
func (t *mgh_MZ) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
