package af

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type af struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'af' locale
func New() locales.Translator {
	return &af{
		locale:  "af",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *af) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'af'
func (t *af) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'af'
func (t *af) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
