package ln_CF

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ln_CF struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ln_CF' locale
func New() locales.Translator {
	return &ln_CF{
		locale:  "ln_CF",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ln_CF) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ln_CF'
func (t *ln_CF) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ln_CF'
func (t *ln_CF) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n >= 0 && n <= 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
