package pa_Arab_PK

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type pa_Arab_PK struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'pa_Arab_PK' locale
func New() locales.Translator {
	return &pa_Arab_PK{
		locale:  "pa_Arab_PK",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *pa_Arab_PK) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'pa_Arab_PK'
func (t *pa_Arab_PK) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'pa_Arab_PK'
func (t *pa_Arab_PK) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n >= 0 && n <= 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
