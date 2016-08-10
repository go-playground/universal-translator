package be_BY

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type be_BY struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'be_BY' locale
func New() locales.Translator {
	return &be_BY{
		locale:   "be_BY",
		plurals:  []locales.PluralRule{2, 4, 5, 6},
		decimal:  []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x32, 0x63, 0x7d},
		group:    []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x63, 0x32, 0x2c, 0x20, 0x30, 0x78, 0x61, 0x30, 0x7d},
		minus:    []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x32, 0x64, 0x7d},
		percent:  []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x32, 0x35, 0x7d},
		perMille: []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x65, 0x32, 0x2c, 0x20, 0x30, 0x78, 0x38, 0x30, 0x2c, 0x20, 0x30, 0x78, 0x62, 0x30, 0x7d},
		symbol:   []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x7d},
	}
}

// Locale returns the current translators string locale
func (t *be_BY) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'be_BY'
func (t *be_BY) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'be_BY'
func (t *be_BY) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n%10 == 1 && n%100 != 11 {
		return locales.PluralRuleOne
	} else if n%10 >= 2 && n%10 <= 4 && n%100 < 12 && n%100 > 14 {
		return locales.PluralRuleFew
	} else if (n%10 == 0) || (n%10 >= 5 && n%10 <= 9) || (n%100 >= 11 && n%100 <= 14) {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}
