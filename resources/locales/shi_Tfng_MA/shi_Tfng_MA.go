package shi_Tfng_MA

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type shi_Tfng_MA struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'shi_Tfng_MA' locale
func New() locales.Translator {
	return &shi_Tfng_MA{
		locale:  "shi_Tfng_MA",
		plurals: []locales.PluralRule{2, 4, 6},
	}
}

// Locale returns the current translators string locale
func (t *shi_Tfng_MA) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'shi_Tfng_MA'
func (t *shi_Tfng_MA) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'shi_Tfng_MA'
func (t *shi_Tfng_MA) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne
	} else if n >= 2 && n <= 10 {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}
