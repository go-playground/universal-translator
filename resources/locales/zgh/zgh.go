package zgh

import "github.com/go-playground/universal-translator/resources/locales"

type zgh struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'zgh' locale
func New() locales.Translator {
	return &zgh{
		locale:   "zgh",
		plurals:  nil,
		decimal:  []byte{0x2c},
		group:    []byte{0xc2, 0xa0},
		minus:    []byte{},
		percent:  []byte{},
		perMille: []byte{},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *zgh) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'zgh'
func (t *zgh) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'zgh'
func (t *zgh) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}
