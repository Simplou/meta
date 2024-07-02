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
		"Afrikaans":        lc.Afrikaans,
		"Albanian":         lc.Albanian,
		"Arabic":           lc.Arabic,
		"Azerbaijani":      lc.Azerbaijani,
		"Bengali":          lc.Bengali,
		"Bulgarian":        lc.Bulgarian,
		"Catalan":          lc.Catalan,
		"Chinese (CHN)":    lc.ChineseCHN,
		"Chinese (HKG)":    lc.ChineseHKG,
		"Chinese (TAI)":    lc.ChineseTAI,
		"Croatian":         lc.Croatian,
		"Czech":            lc.Czech,
		"Danish":           lc.Danish,
		"Dutch":            lc.Dutch,
		"English":          lc.English,
		"English (UK)":     lc.EnglishUK,
		"English (USA)":    lc.EnglishUSA,
		"Estonian":         lc.Estonian,
		"Filipino":         lc.Filipino,
		"Finnish":          lc.Finnish,
		"French":           lc.French,
		"German":           lc.German,
		"Greek":            lc.Greek,
		"Gujarati":         lc.Gujarati,
		"Hausa":            lc.Hausa,
		"Hebrew":           lc.Hebrew,
		"Hindi":            lc.Hindi,
		"Hungarian":        lc.Hungarian,
		"Indonesian":       lc.Indonesian,
		"Irish":            lc.Irish,
		"Italian":          lc.Italian,
		"Japanese":         lc.Japanese,
		"Kannada":          lc.Kannada,
		"Kazakh":           lc.Kazakh,
		"Korean":           lc.Korean,
		"Lao":              lc.Lao,
		"Latvian":          lc.Latvian,
		"Lithuanian":       lc.Lithuanian,
		"Macedonian":       lc.Macedonian,
		"Malay":            lc.Malay,
		"Malayalam":        lc.Malayalam,
		"Marathi":          lc.Marathi,
		"Norwegian":        lc.Norwegian,
		"Persian":          lc.Persian,
		"Polish":           lc.Polish,
		"Portuguese (BR)":  lc.PortugueseBR,
		"Portuguese (POR)": lc.PortuguesePOR,
		"Punjabi":          lc.Punjabi,
		"Romanian":         lc.Romanian,
		"Russian":          lc.Russian,
		"Serbian":          lc.Serbian,
		"Slovak":           lc.Slovak,
		"Slovenian":        lc.Slovenian,
		"Spanish":          lc.Spanish,
		"Spanish (ARG)":    lc.SpanishARG,
		"Spanish (ESP)":    lc.SpanishESP,
		"Spanish (MEX)":    lc.SpanishMEX,
		"Swahili":          lc.Swahili,
		"Swedish":          lc.Swedish,
		"Tamil":            lc.Tamil,
		"Telugu":           lc.Telugu,
		"Thai":             lc.Thai,
		"Turkish":          lc.Turkish,
		"Ukrainian":        lc.Ukrainian,
		"Urdu":             lc.Urdu,
		"Uzbek":            lc.Uzbek,
		"Vietnamese":       lc.Vietnamese,
		"Zulu":             lc.Zulu,
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
