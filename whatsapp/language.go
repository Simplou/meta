package whatsapp

import "github.com/Simplou/goxios"

type languageCodes struct {
	Afrikaans     string
	Albanian      string
	Arabic        string
	Azerbaijani   string
	Bengali       string
	Bulgarian     string
	Catalan       string
	ChineseCHN    string
	ChineseHKG    string
	ChineseTAI    string
	Croatian      string
	Czech         string
	Danish        string
	Dutch         string
	English       string
	EnglishUK     string
	EnglishUSA    string
	Estonian      string
	Filipino      string
	Finnish       string
	French        string
	German        string
	Greek         string
	Gujarati      string
	Hausa         string
	Hebrew        string
	Hindi         string
	Hungarian     string
	Indonesian    string
	Irish         string
	Italian       string
	Japanese      string
	Kannada       string
	Kazakh        string
	Korean        string
	Lao           string
	Latvian       string
	Lithuanian    string
	Macedonian    string
	Malay         string
	Malayalam     string
	Marathi       string
	Norwegian     string
	Persian       string
	Polish        string
	PortugueseBR  string
	PortuguesePOR string
	Punjabi       string
	Romanian      string
	Russian       string
	Serbian       string
	Slovak        string
	Slovenian     string
	Spanish       string
	SpanishARG    string
	SpanishESP    string
	SpanishMEX    string
	Swahili       string
	Swedish       string
	Tamil         string
	Telugu        string
	Thai          string
	Turkish       string
	Ukrainian     string
	Urdu          string
	Uzbek         string
	Vietnamese    string
	Zulu          string
}

func (lc languageCodes) JSON() goxios.GenericJSON[string] {
	return goxios.GenericJSON[string]{
		"Afrikaans":        "af",
		"Albanian":         "sq",
		"Arabic":           "ar",
		"Azerbaijani":      "az",
		"Bengali":          "bn",
		"Bulgarian":        "bg",
		"Catalan":          "ca",
		"Chinese (CHN)":    "zh_CN",
		"Chinese (HKG)":    "zh_HK",
		"Chinese (TAI)":    "zh_TW",
		"Croatian":         "hr",
		"Czech":            "cs",
		"Danish":           "da",
		"Dutch":            "nl",
		"English":          "en",
		"English (UK)":     "en_GB",
		"English (USA)":    "en_US",
		"Estonian":         "et",
		"Filipino":         "fil",
		"Finnish":          "fi",
		"French":           "fr",
		"German":           "de",
		"Greek":            "el",
		"Gujarati":         "gu",
		"Hausa":            "ha",
		"Hebrew":           "he",
		"Hindi":            "hi",
		"Hungarian":        "hu",
		"Indonesian":       "id",
		"Irish":            "ga",
		"Italian":          "it",
		"Japanese":         "ja",
		"Kannada":          "kn",
		"Kazakh":           "kk",
		"Korean":           "ko",
		"Lao":              "lo",
		"Latvian":          "lv",
		"Lithuanian":       "lt",
		"Macedonian":       "mk",
		"Malay":            "ms",
		"Malayalam":        "ml",
		"Marathi":          "mr",
		"Norwegian":        "nb",
		"Persian":          "fa",
		"Polish":           "pl",
		"Portuguese (BR)":  "pt_BR",
		"Portuguese (POR)": "pt_PT",
		"Punjabi":          "pa",
		"Romanian":         "ro",
		"Russian":          "ru",
		"Serbian":          "sr",
		"Slovak":           "sk",
		"Slovenian":        "sl",
		"Spanish":          "es",
		"Spanish (ARG)":    "es_AR",
		"Spanish (ESP)":    "es_ES",
		"Spanish (MEX)":    "es_MX",
		"Swahili":          "sw",
		"Swedish":          "sv",
		"Tamil":            "ta",
		"Telugu":           "te",
		"Thai":             "th",
		"Turkish":          "tr",
		"Ukrainian":        "uk",
		"Urdu":             "ur",
		"Uzbek":            "uz",
		"Vietnamese":       "vi",
		"Zulu":             "zu",
	}
}

func Languages() languageCodes {
	lcs := languageCodes{
		Afrikaans:     "af",
		Albanian:      "sq",
		Arabic:        "ar",
		Azerbaijani:   "az",
		Bengali:       "bn",
		Bulgarian:     "bg",
		Catalan:       "ca",
		ChineseCHN:    "zh_CN",
		ChineseHKG:    "zh_HK",
		ChineseTAI:    "zh_TW",
		Croatian:      "hr",
		Czech:         "cs",
		Danish:        "da",
		Dutch:         "nl",
		English:       "en",
		EnglishUK:     "en_GB",
		EnglishUSA:    "en_US",
		Estonian:      "et",
		Filipino:      "fil",
		Finnish:       "fi",
		French:        "fr",
		German:        "de",
		Greek:         "el",
		Gujarati:      "gu",
		Hausa:         "ha",
		Hebrew:        "he",
		Hindi:         "hi",
		Hungarian:     "hu",
		Indonesian:    "id",
		Irish:         "ga",
		Italian:       "it",
		Japanese:      "ja",
		Kannada:       "kn",
		Kazakh:        "kk",
		Korean:        "ko",
		Lao:           "lo",
		Latvian:       "lv",
		Lithuanian:    "lt",
		Macedonian:    "mk",
		Malay:         "ms",
		Malayalam:     "ml",
		Marathi:       "mr",
		Norwegian:     "nb",
		Persian:       "fa",
		Polish:        "pl",
		PortugueseBR:  "pt_BR",
		PortuguesePOR: "pt_PT",
		Punjabi:       "pa",
		Romanian:      "ro",
		Russian:       "ru",
		Serbian:       "sr",
		Slovak:        "sk",
		Slovenian:     "sl",
		Spanish:       "es",
		SpanishARG:    "es_AR",
		SpanishESP:    "es_ES",
		SpanishMEX:    "es_MX",
		Swahili:       "sw",
		Swedish:       "sv",
		Tamil:         "ta",
		Telugu:        "te",
		Thai:          "th",
		Turkish:       "tr",
		Ukrainian:     "uk",
		Urdu:          "ur",
		Uzbek:         "uz",
		Vietnamese:    "vi",
		Zulu:          "zu",
	}
	return lcs
}
