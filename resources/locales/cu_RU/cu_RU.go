package cu_RU

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type cu_RU struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'cu_RU' locale
func New() locales.Translator {
	return &cu_RU{
		locale:  "cu_RU",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *cu_RU) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'cu_RU'
func (t *cu_RU) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'cu_RU'
func (t *cu_RU) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
