package km

import "github.com/go-playground/universal-translator/resources/locales"

type km struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'km' locale
func New() locales.Translator {
	return &km{
		locale:   "km",
		plurals:  []locales.PluralRule{6},
		decimal:  []byte{0x2c},
		group:    []byte{0x2e},
		minus:    []byte{0x2d},
		percent:  []byte{0x25},
		perMille: []byte{0xe2, 0x80, 0xb0},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *km) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'km'
func (t *km) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'km'
func (t *km) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}
