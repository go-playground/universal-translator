package vai_Latn

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type vai_Latn struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'vai_Latn' locale
func New() locales.Translator {
	return &vai_Latn{
		locale:  "vai_Latn",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *vai_Latn) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'vai_Latn'
func (t *vai_Latn) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'vai_Latn'
func (t *vai_Latn) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
