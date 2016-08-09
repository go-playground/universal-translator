package cu

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type cu struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'cu' locale
func New() locales.Translator {
	return &cu{
		locale:  "cu",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *cu) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'cu'
func (t *cu) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'cu'
func (t *cu) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
