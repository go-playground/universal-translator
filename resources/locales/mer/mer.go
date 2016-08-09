package mer

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type mer struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'mer' locale
func New() locales.Translator {
	return &mer{
		locale:  "mer",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *mer) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mer'
func (t *mer) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'mer'
func (t *mer) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
