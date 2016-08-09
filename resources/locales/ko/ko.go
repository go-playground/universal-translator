package ko

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ko struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ko' locale
func New() locales.Translator {
	return &ko{
		locale:  "ko",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *ko) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ko'
func (t *ko) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ko'
func (t *ko) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
