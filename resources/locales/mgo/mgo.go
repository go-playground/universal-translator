package mgo

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type mgo struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'mgo' locale
func New() locales.Translator {
	return &mgo{
		locale:  "mgo",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *mgo) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mgo'
func (t *mgo) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'mgo'
func (t *mgo) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
