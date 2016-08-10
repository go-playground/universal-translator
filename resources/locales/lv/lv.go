package lv

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type lv struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'lv' locale
func New() locales.Translator {
	return &lv{
		locale:   "lv",
		plurals:  []locales.PluralRule{1, 2, 6},
		decimal:  []byte{0x2c},
		group:    []byte{0xc2, 0xa0},
		minus:    []byte{0x2d},
		percent:  []byte{0x25},
		perMille: []byte{0xe2, 0x80, 0xb0},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *lv) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'lv'
func (t *lv) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'lv'
func (t *lv) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	f := locales.F(n, v)

	if (n%10 == 0) || (n%100 >= 11 && n%100 <= 19) || (v == 2 && f%100 >= 11 && f%100 <= 19) {
		return locales.PluralRuleZero
	} else if (n%10 == 1 && n%100 != 11) || (v == 2 && f%10 == 1 && f%100 != 11) || (v != 2 && f%10 == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}
