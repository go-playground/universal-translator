package vi_VN

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type vi_VN struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'vi_VN' locale
func New() locales.Translator {
	return &vi_VN{
		locale:  "vi_VN",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *vi_VN) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'vi_VN'
func (t *vi_VN) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'vi_VN'
func (t *vi_VN) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
