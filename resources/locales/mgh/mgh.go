package mgh

import "github.com/go-playground/universal-translator/resources/locales"

type mgh struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'mgh' locale
func New() locales.Translator {
	return &mgh{
		locale:   "mgh",
		plurals:  nil,
		decimal:  []byte{0x2c},
		group:    []byte{0x2e},
		minus:    []byte{},
		percent:  []byte{},
		perMille: []byte{},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *mgh) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mgh'
func (t *mgh) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'mgh'
func (t *mgh) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}
