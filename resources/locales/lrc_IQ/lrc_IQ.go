package lrc_IQ

import "github.com/go-playground/universal-translator/resources/locales"

type lrc_IQ struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'lrc_IQ' locale
func New() locales.Translator {
	return &lrc_IQ{
		locale:   "lrc_IQ",
		plurals:  nil,
		decimal:  []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x32, 0x65, 0x7d},
		group:    []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x32, 0x63, 0x7d},
		minus:    []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x32, 0x64, 0x7d},
		percent:  []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x32, 0x35, 0x7d},
		perMille: []byte{},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *lrc_IQ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'lrc_IQ'
func (t *lrc_IQ) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'lrc_IQ'
func (t *lrc_IQ) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}
