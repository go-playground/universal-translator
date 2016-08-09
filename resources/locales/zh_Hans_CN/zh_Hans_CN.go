package zh_Hans_CN

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type zh_Hans_CN struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'zh_Hans_CN' locale
func New() locales.Translator {
	return &zh_Hans_CN{
		locale:  "zh_Hans_CN",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *zh_Hans_CN) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'zh_Hans_CN'
func (t *zh_Hans_CN) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'zh_Hans_CN'
func (t *zh_Hans_CN) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
