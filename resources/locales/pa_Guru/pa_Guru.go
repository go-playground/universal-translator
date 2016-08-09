package pa_Guru

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type pa_Guru struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'pa_Guru' locale
func New() locales.Translator {
	return &pa_Guru{
		locale:  "pa_Guru",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *pa_Guru) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'pa_Guru'
func (t *pa_Guru) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'pa_Guru'
func (t *pa_Guru) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n >= 0 && n <= 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
