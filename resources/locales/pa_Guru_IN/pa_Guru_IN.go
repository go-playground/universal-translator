package pa_Guru_IN

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type pa_Guru_IN struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'pa_Guru_IN' locale
func New() locales.Translator {
	return &pa_Guru_IN{
		locale:  "pa_Guru_IN",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *pa_Guru_IN) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'pa_Guru_IN'
func (t *pa_Guru_IN) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'pa_Guru_IN'
func (t *pa_Guru_IN) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n >= 0 && n <= 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
