package dyo

import "github.com/go-playground/universal-translator/resources/locales"

type dyo struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'dyo' locale
func New() locales.Translator {
	return &dyo{
		locale:   "dyo",
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
func (t *dyo) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'dyo'
func (t *dyo) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'dyo'
func (t *dyo) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}
