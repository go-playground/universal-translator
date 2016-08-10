package ar

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ar struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'ar' locale
func New() locales.Translator {
	return &ar{
		locale:   "ar",
		plurals:  []locales.PluralRule{1, 2, 3, 4, 5, 6},
		decimal:  []byte{0xd9, 0xab},
		group:    []byte{0xd9, 0xac},
		minus:    []byte{0xe2, 0x80, 0x8f, 0x2d},
		percent:  []byte{0xd9, 0xaa},
		perMille: []byte{0xd8, 0x89},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *ar) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ar'
func (t *ar) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ar'
func (t *ar) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

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
