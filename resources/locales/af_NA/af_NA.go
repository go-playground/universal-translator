package af_NA

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type af_NA struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'af_NA' locale
func New() locales.Translator {
	return &af_NA{
		locale:  "af_NA",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *af_NA) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'af_NA'
func (t *af_NA) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'af_NA'
func (t *af_NA) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
