package ut

// UniversalTranslator holds all locale Translator instances
type UniversalTranslator struct {
	translators map[string]*Translator
	fallback    *Translator
}

// New creates a new UniversalTranslator instance.
func New() *UniversalTranslator {
	return &UniversalTranslator{}
}

// SetFallback registers the fallback Translator instance when no matching Translator
// can be found via a given locale
func (ut *UniversalTranslator) SetFallback(translator *Translator) {
	ut.fallback = translator
}

// LoadTranslator instantiates, if not already exists, the specified translator
// for the given locale
func (ut *UniversalTranslator) LoadTranslator(locale string) (*Translator, error) {

	if t, ok := ut.translators[locale]; ok {
		return t, nil
	}

	t, err := newTranslator(locale)
	if err != nil {
		return nil, err
	}

	ut.translators[locale] = t

	// just to ensure it get's set to something
	if ut.fallback == nil {
		ut.fallback = t
	}

	return t, nil
}
