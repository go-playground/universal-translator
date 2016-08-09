package ar_DJ

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type ar_DJ struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ar_DJ' locale
func New() locales.Translator {
	return &ar_DJ{
		locale:  "ar_DJ",
		plurals: []locales.PluralRule{1, 2, 3, 4, 5, 6},
	}
}

// Locale returns the current translators string locale
func (t *ar_DJ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ar_DJ'
func (t *ar_DJ) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ar_DJ'
func (t *ar_DJ) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 0 {
		return locales.PluralRuleZero
	} else if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	} else if n%100 >= 3 && n%100 <= 10 {
		return locales.PluralRuleFew
	} else if n%100 >= 11 && n%100 <= 99 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}
