package rn

import "github.com/go-playground/universal-translator/resources/locales"

type rn struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'rn' locale
func New() locales.Translator {
	return &rn{
		locale:   "rn",
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
func (t *rn) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'rn'
func (t *rn) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'rn'
func (t *rn) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}
