package qu_PE

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type qu_PE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'qu_PE' locale
func New() locales.Translator {
	return &qu_PE{
		locale:  "qu_PE",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *qu_PE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'qu_PE'
func (t *qu_PE) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'qu_PE'
func (t *qu_PE) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
