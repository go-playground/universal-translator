package bas

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type bas struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'bas' locale
func New() locales.Translator {
	return &bas{
		locale:  "bas",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *bas) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'bas'
func (t *bas) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'bas'
func (t *bas) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
