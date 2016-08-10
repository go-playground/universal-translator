package rw_RW

import "github.com/go-playground/universal-translator/resources/locales"

type rw_RW struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'rw_RW' locale
func New() locales.Translator {
	return &rw_RW{
		locale:   "rw_RW",
		plurals:  nil,
		decimal:  []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x32, 0x63, 0x7d},
		group:    []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x32, 0x65, 0x7d},
		minus:    []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x7d},
		percent:  []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x7d},
		perMille: []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x7d},
		symbol:   []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x7d},
	}
}

// Locale returns the current translators string locale
func (t *rw_RW) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'rw_RW'
func (t *rw_RW) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'rw_RW'
func (t *rw_RW) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}
