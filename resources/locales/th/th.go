package th

import "github.com/go-playground/universal-translator/resources/locales"

type th struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'th' locale
func New() locales.Translator {
	return &th{
		locale:   "th",
		plurals:  []locales.PluralRule{6},
		decimal:  []byte{0x2e},
		group:    []byte{0x2c},
		minus:    []byte{0x2d},
		percent:  []byte{0x25},
		perMille: []byte{0xe2, 0x80, 0xb0},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *th) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'th'
func (t *th) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'th'
func (t *th) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}
