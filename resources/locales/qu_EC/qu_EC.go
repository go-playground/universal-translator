package qu_EC

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type qu_EC struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'qu_EC' locale
func New() locales.Translator {
	return &qu_EC{
		locale:  "qu_EC",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *qu_EC) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'qu_EC'
func (t *qu_EC) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'qu_EC'
func (t *qu_EC) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
