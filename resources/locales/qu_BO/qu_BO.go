package qu_BO

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type qu_BO struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'qu_BO' locale
func New() locales.Translator {
	return &qu_BO{
		locale:  "qu_BO",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *qu_BO) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'qu_BO'
func (t *qu_BO) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'qu_BO'
func (t *qu_BO) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
