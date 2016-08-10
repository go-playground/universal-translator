package bm_ML

import "github.com/go-playground/universal-translator/resources/locales"

type bm_ML struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'bm_ML' locale
func New() locales.Translator {
	return &bm_ML{
		locale:   "bm_ML",
		plurals:  []locales.PluralRule{6},
		decimal:  []byte{},
		group:    []byte{},
		minus:    []byte{},
		percent:  []byte{},
		perMille: []byte{},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *bm_ML) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'bm_ML'
func (t *bm_ML) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'bm_ML'
func (t *bm_ML) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}
