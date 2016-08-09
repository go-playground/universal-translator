package khq_ML

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type khq_ML struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'khq_ML' locale
func New() locales.Translator {
	return &khq_ML{
		locale:  "khq_ML",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *khq_ML) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'khq_ML'
func (t *khq_ML) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'khq_ML'
func (t *khq_ML) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown, nil
}
