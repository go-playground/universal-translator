package twq

import "github.com/go-playground/universal-translator/resources/locales"

type twq struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'twq' locale
func New() locales.Translator {
	return &twq{
		locale:   "twq",
		plurals:  nil,
		decimal:  []byte{0x2e},
		group:    []byte{0xc2, 0xa0},
		minus:    []byte{},
		percent:  []byte{},
		perMille: []byte{},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *twq) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'twq'
func (t *twq) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'twq'
func (t *twq) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}
