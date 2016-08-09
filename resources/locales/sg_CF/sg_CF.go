package sg_CF

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type sg_CF struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'sg_CF' locale
func New() locales.Translator {
	return &sg_CF{
		locale:  "sg_CF",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *sg_CF) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sg_CF'
func (t *sg_CF) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'sg_CF'
func (t *sg_CF) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
