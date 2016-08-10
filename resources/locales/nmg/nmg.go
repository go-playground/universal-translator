package nmg

import "github.com/go-playground/universal-translator/resources/locales"

type nmg struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'nmg' locale
func New() locales.Translator {
	return &nmg{
		locale:   "nmg",
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
func (t *nmg) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'nmg'
func (t *nmg) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'nmg'
func (t *nmg) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}
