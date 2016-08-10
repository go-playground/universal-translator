package dyo_SN

import "github.com/go-playground/universal-translator/resources/locales"

type dyo_SN struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'dyo_SN' locale
func New() locales.Translator {
	return &dyo_SN{
		locale:   "dyo_SN",
		plurals:  nil,
		decimal:  []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x32, 0x63, 0x7d},
		group:    []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x63, 0x32, 0x2c, 0x20, 0x30, 0x78, 0x61, 0x30, 0x7d},
		minus:    []byte{},
		percent:  []byte{},
		perMille: []byte{},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *dyo_SN) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'dyo_SN'
func (t *dyo_SN) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'dyo_SN'
func (t *dyo_SN) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}
