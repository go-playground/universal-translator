package uz

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type uz struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'uz' locale
func New() locales.Translator {
	return &uz{
		locale:   "uz",
		plurals:  []locales.PluralRule{2, 6},
		decimal:  []byte{0xd9, 0xab},
		group:    []byte{0xd9, 0xac},
		minus:    []byte{0x2d},
		percent:  []byte{0xd9, 0xaa},
		perMille: []byte{0xd8, 0x89},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *uz) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'uz'
func (t *uz) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'uz'
func (t *uz) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
