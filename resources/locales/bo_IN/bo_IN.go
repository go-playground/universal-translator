package bo_IN

import "github.com/go-playground/universal-translator/resources/locales"

type bo_IN struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'bo_IN' locale
func New() locales.Translator {
	return &bo_IN{
		locale:   "bo_IN",
		plurals:  []locales.PluralRule{6},
		decimal:  []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x32, 0x65, 0x7d},
		group:    []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x32, 0x63, 0x7d},
		minus:    []byte{},
		percent:  []byte{},
		perMille: []byte{},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *bo_IN) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'bo_IN'
func (t *bo_IN) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'bo_IN'
func (t *bo_IN) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}
