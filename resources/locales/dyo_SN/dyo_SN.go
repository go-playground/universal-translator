package dyo_SN

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type dyo_SN struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'dyo_SN' locale
func New() locales.Translator {
	return &dyo_SN{
		locale:  "dyo_SN",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *dyo_SN) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'dyo_SN'
func (t *dyo_SN) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'dyo_SN'
func (t *dyo_SN) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
