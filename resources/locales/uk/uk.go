package uk

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type uk struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'uk' locale
func New() locales.Translator {
	return &uk{
		locale:   "uk",
		plurals:  []locales.PluralRule{2, 4, 5, 6},
		decimal:  []byte{0x2c},
		group:    []byte{0xc2, 0xa0},
		minus:    []byte{0x2d},
		percent:  []byte{0x25},
		perMille: []byte{0xe2, 0x80, 0xb0},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *uk) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'uk'
func (t *uk) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'uk'
func (t *uk) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if v == 0 && i%10 == 1 && i%100 != 11 {
		return locales.PluralRuleOne
	} else if v == 0 && i%10 >= 2 && i%10 <= 4 && i%100 < 12 && i%100 > 14 {
		return locales.PluralRuleFew
	} else if (v == 0 && i%10 == 0) || (v == 0 && i%10 >= 5 && i%10 <= 9) || (v == 0 && i%100 >= 11 && i%100 <= 14) {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}
