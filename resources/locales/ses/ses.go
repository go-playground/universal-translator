package ses

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ses struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ses' locale
func New() locales.Translator {
	return &ses{
		locale:  "ses",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *ses) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ses'
func (t *ses) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ses'
func (t *ses) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
