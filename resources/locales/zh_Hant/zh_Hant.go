package zh_Hant

import "github.com/go-playground/universal-translator/resources/locales"

type zh_Hant struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'zh_Hant' locale
func New() locales.Translator {
	return &zh_Hant{
		locale:   "zh_Hant",
		plurals:  []locales.PluralRule{6},
		decimal:  []byte{0x2e},
		group:    []byte{0x2c},
		minus:    []byte{0x2d},
		percent:  []byte{0x25},
		perMille: []byte{0xe2, 0x80, 0xb0},
		symbol:   []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x7d},
	}
}

// Locale returns the current translators string locale
func (t *zh_Hant) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'zh_Hant'
func (t *zh_Hant) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'zh_Hant'
func (t *zh_Hant) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}
