package khq

import "github.com/go-playground/universal-translator/resources/locales"

type khq struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'khq' locale
func New() locales.Translator {
	return &khq{
		locale:   "khq",
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
func (t *khq) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'khq'
func (t *khq) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'khq'
func (t *khq) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}
