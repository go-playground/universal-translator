package ln_CD

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ln_CD struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ln_CD' locale
func New() locales.Translator {
	return &ln_CD{
		locale:  "ln_CD",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ln_CD) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ln_CD'
func (t *ln_CD) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ln_CD'
func (t *ln_CD) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n >= 0 && n <= 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
