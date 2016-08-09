package zh_Hans_HK

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type zh_Hans_HK struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'zh_Hans_HK' locale
func New() locales.Translator {
	return &zh_Hans_HK{
		locale:  "zh_Hans_HK",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *zh_Hans_HK) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'zh_Hans_HK'
func (t *zh_Hans_HK) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'zh_Hans_HK'
func (t *zh_Hans_HK) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
