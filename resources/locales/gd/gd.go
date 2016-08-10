package gd

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type gd struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'gd' locale
func New() locales.Translator {
	return &gd{
		locale:   "gd",
		plurals:  []locales.PluralRule{2, 3, 4, 6},
		decimal:  []byte{0xd9, 0xab},
		group:    []byte{0xd9, 0xac},
		minus:    []byte{0xe2, 0x80, 0x8f, 0x2d},
		percent:  []byte{0xd9, 0xaa},
		perMille: []byte{0xd8, 0x89},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *gd) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'gd'
func (t *gd) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'gd'
func (t *gd) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 || n == 11 {
		return locales.PluralRuleOne
	} else if n == 2 || n == 12 {
		return locales.PluralRuleTwo
	} else if (n >= 3 && n <= 10) || (n >= 13 && n <= 19) {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}
