package ut

import "fmt"

type translation struct {
	singular  string
	plural    string
	hasPlural bool
}

// Translator holds the locale translation instance
type Translator struct {
	locale       *Locale
	pluralFunc   PluralRulerFunc
	translations map[string]*translation
	groups       map[string][]*translation
}

func newTranslator(locale string) (*Translator, error) {
	return nil, nil
}

// Add registers a new translation to the Translator using the
// key for a unique translation and group allows for retrieving
// translations for dev purposes later i.e. load all translations
// for the "homepage" group
// plural is optional, can be blank, but allows adding a plural translation
// at the same time as a singular
func (t *Translator) Add(group string, key string, singular string, plural string) {

	trans := &translation{
		singular:  singular,
		plural:    plural,
		hasPlural: plural != "",
	}

	if _, ok := t.translations[key]; ok {
		panic(fmt.Sprintf("Translation with key '%s' already exists", key))
	}

	t.translations[key] = trans

	if _, ok := t.groups[group]; !ok {
		t.groups[group] = make([]*translation, 0)
	}

	t.groups[group] = append(t.groups[group], trans)
}

func (t *Translator) pluralRule(count NumberValue) (rule PluralRule) {
	// l, ok := GetLocale(locale)
	// if !ok {
	// 	return PluralRuleOther
	// }

	// ruler, ok := pluralRules[t.locale.PluralRule]

	// if !ok {
	// 	return PluralRuleOther
	// }
	//  ruler.FindRule(count)
	return t.pluralFunc(count)
}

// T translates the text associated with the given key with the
// arguments passed in
func (t *Translator) T(key string, a ...interface{}) {

}
