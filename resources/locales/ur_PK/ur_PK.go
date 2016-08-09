package ur_PK

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ur_PK struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ur_PK' locale
func New() locales.Translator {
	return &ur_PK{
		locale:  "ur_PK",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ur_PK) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ur_PK'
func (t *ur_PK) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ur_PK'
func (t *ur_PK) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
