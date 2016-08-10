package mfe

import "github.com/go-playground/universal-translator/resources/locales"

type mfe struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'mfe' locale
func New() locales.Translator {
	return &mfe{
		locale:   "mfe",
		plurals:  nil,
		decimal:  []byte{},
		group:    []byte{0xc2, 0xa0},
		minus:    []byte{},
		percent:  []byte{},
		perMille: []byte{},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *mfe) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mfe'
func (t *mfe) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'mfe'
func (t *mfe) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}
