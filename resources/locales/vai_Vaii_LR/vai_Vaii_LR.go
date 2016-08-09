package vai_Vaii_LR

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type vai_Vaii_LR struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'vai_Vaii_LR' locale
func New() locales.Translator {
	return &vai_Vaii_LR{
		locale:  "vai_Vaii_LR",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *vai_Vaii_LR) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'vai_Vaii_LR'
func (t *vai_Vaii_LR) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'vai_Vaii_LR'
func (t *vai_Vaii_LR) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
