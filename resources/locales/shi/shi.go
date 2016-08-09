package shi

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type shi struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'shi' locale
func New() locales.Translator {
	return &shi{
		locale:  "shi",
		plurals: []locales.PluralRule{2, 4, 6},
	}
}

// Locale returns the current translators string locale
func (t *shi) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'shi'
func (t *shi) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'shi'
func (t *shi) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne
	} else if n >= 2 && n <= 10 {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}
