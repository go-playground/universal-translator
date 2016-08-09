package ln_CG

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ln_CG struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ln_CG' locale
func New() locales.Translator {
	return &ln_CG{
		locale:  "ln_CG",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ln_CG) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ln_CG'
func (t *ln_CG) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ln_CG'
func (t *ln_CG) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n >= 0 && n <= 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
