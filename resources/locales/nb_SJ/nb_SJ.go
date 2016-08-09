package nb_SJ

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type nb_SJ struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'nb_SJ' locale
func New() locales.Translator {
	return &nb_SJ{
		locale:  "nb_SJ",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *nb_SJ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'nb_SJ'
func (t *nb_SJ) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'nb_SJ'
func (t *nb_SJ) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
