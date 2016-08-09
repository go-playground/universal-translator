package ms_SG

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ms_SG struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ms_SG' locale
func New() locales.Translator {
	return &ms_SG{
		locale:  "ms_SG",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *ms_SG) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ms_SG'
func (t *ms_SG) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ms_SG'
func (t *ms_SG) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
