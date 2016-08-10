package nd_ZW

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type nd_ZW struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'nd_ZW' locale
func New() locales.Translator {
	return &nd_ZW{
		locale:   "nd_ZW",
		plurals:  []locales.PluralRule{2, 6},
		decimal:  []byte{},
		group:    []byte{},
		minus:    []byte{},
		percent:  []byte{},
		perMille: []byte{},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *nd_ZW) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'nd_ZW'
func (t *nd_ZW) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'nd_ZW'
func (t *nd_ZW) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
