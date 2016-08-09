package gsw_FR

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type gsw_FR struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'gsw_FR' locale
func New() locales.Translator {
	return &gsw_FR{
		locale:  "gsw_FR",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *gsw_FR) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'gsw_FR'
func (t *gsw_FR) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'gsw_FR'
func (t *gsw_FR) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
