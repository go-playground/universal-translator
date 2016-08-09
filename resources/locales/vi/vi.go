package vi

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type vi struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'vi' locale
func New() locales.Translator {
	return &vi{
		locale:  "vi",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *vi) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'vi'
func (t *vi) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'vi'
func (t *vi) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
