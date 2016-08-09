package sbp_TZ

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type sbp_TZ struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'sbp_TZ' locale
func New() locales.Translator {
	return &sbp_TZ{
		locale:  "sbp_TZ",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *sbp_TZ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sbp_TZ'
func (t *sbp_TZ) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'sbp_TZ'
func (t *sbp_TZ) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
