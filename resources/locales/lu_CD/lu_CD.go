package lu_CD

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type lu_CD struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'lu_CD' locale
func New() locales.Translator {
	return &lu_CD{
		locale:  "lu_CD",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *lu_CD) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'lu_CD'
func (t *lu_CD) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'lu_CD'
func (t *lu_CD) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
