package ee_TG

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ee_TG struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'ee_TG' locale
func New() locales.Translator {
	return &ee_TG{
		locale:   "ee_TG",
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
func (t *ee_TG) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ee_TG'
func (t *ee_TG) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ee_TG'
func (t *ee_TG) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
