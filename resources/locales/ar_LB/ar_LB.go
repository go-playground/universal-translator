package ar_LB

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ar_LB struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'ar_LB' locale
func New() locales.Translator {
	return &ar_LB{
		locale:   "ar_LB",
		plurals:  []locales.PluralRule{1, 2, 3, 4, 5, 6},
		decimal:  []byte{0x2c},
		group:    []byte{0x2e},
		minus:    []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x65, 0x32, 0x2c, 0x20, 0x30, 0x78, 0x38, 0x30, 0x2c, 0x20, 0x30, 0x78, 0x38, 0x66, 0x2c, 0x20, 0x30, 0x78, 0x32, 0x64, 0x7d},
		percent:  []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x64, 0x39, 0x2c, 0x20, 0x30, 0x78, 0x61, 0x61, 0x7d},
		perMille: []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x64, 0x38, 0x2c, 0x20, 0x30, 0x78, 0x38, 0x39, 0x7d},
		symbol:   []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x7d},
	}
}

// Locale returns the current translators string locale
func (t *ar_LB) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ar_LB'
func (t *ar_LB) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ar_LB'
func (t *ar_LB) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 0 {
		return locales.PluralRuleZero
	} else if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	} else if n%100 >= 3 && n%100 <= 10 {
		return locales.PluralRuleFew
	} else if n%100 >= 11 && n%100 <= 99 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}
