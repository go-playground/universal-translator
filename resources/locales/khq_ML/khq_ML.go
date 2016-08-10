package khq_ML

import "github.com/go-playground/universal-translator/resources/locales"

type khq_ML struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'khq_ML' locale
func New() locales.Translator {
	return &khq_ML{
		locale:   "khq_ML",
		plurals:  nil,
		decimal:  []byte{},
		group:    []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x63, 0x32, 0x2c, 0x20, 0x30, 0x78, 0x61, 0x30, 0x7d},
		minus:    []byte{},
		percent:  []byte{},
		perMille: []byte{},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *khq_ML) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'khq_ML'
func (t *khq_ML) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'khq_ML'
func (t *khq_ML) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}
