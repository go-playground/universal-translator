package id_ID

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type id_ID struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'id_ID' locale
func New() locales.Translator {
	return &id_ID{
		locale:  "id_ID",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *id_ID) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'id_ID'
func (t *id_ID) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'id_ID'
func (t *id_ID) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	return locales.PluralRuleOther
}
