package vai_Latn_LR

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type vai_Latn_LR struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'vai_Latn_LR' locale
func New() locales.Translator {
	return &vai_Latn_LR{
		locale:  "vai_Latn_LR",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *vai_Latn_LR) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'vai_Latn_LR'
func (t *vai_Latn_LR) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'vai_Latn_LR'
func (t *vai_Latn_LR) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
