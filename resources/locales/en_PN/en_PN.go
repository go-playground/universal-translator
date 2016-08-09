package en_PN

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type en_PN struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'en_PN' locale
func New() locales.Translator {
	return &en_PN{
		locale:  "en_PN",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *en_PN) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'en_PN'
func (t *en_PN) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'en_PN'
func (t *en_PN) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
