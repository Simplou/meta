package whatsapp

import (
	"testing"
)

func TestLanguages(t *testing.T) {
	langs := Languages()
	Languages := langs.JSON()
	expectedLength := 68
	if len(Languages) != expectedLength {
		t.Errorf("Tamanho incorreto do mapa Languages. Esperado: %d, Obtido: %d", expectedLength, len(Languages))
	}

	testCases := []struct {
		languageName         string
		expectedLanguageCode string
	}{
		{"Afrikaans", "af"},
		{"Albanian", "sq"},
		{"Arabic", "ar"},
		{"Azerbaijani", "az"},
		{"Bengali", "bn"},
		{"Bulgarian", "bg"},
		{"Catalan", "ca"},
		{"Chinese (CHN)", "zh_CN"},
		{"Chinese (HKG)", "zh_HK"},
		{"Chinese (TAI)", "zh_TW"},
		{"Croatian", "hr"},
		{"Czech", "cs"},
		{"Danish", "da"},
		{"Dutch", "nl"},
		{"English", "en"},
		{"English (UK)", "en_GB"},
		{"English (USA)", "en_US"},
		{"Estonian", "et"},
		{"Filipino", "fil"},
		{"Finnish", "fi"},
		{"French", "fr"},
		{"German", "de"},
		{"Greek", "el"},
		{"Gujarati", "gu"},
		{"Hausa", "ha"},
		{"Hebrew", "he"},
		{"Hindi", "hi"},
		{"Hungarian", "hu"},
		{"Indonesian", "id"},
		{"Irish", "ga"},
		{"Italian", "it"},
		{"Japanese", "ja"},
		{"Kannada", "kn"},
		{"Kazakh", "kk"},
		{"Korean", "ko"},
		{"Lao", "lo"},
		{"Latvian", "lv"},
		{"Lithuanian", "lt"},
		{"Macedonian", "mk"},
		{"Malay", "ms"},
		{"Malayalam", "ml"},
		{"Marathi", "mr"},
		{"Norwegian", "nb"},
		{"Persian", "fa"},
		{"Polish", "pl"},
		{"Portuguese (BR)", "pt_BR"},
		{"Portuguese (POR)", "pt_PT"},
		{"Punjabi", "pa"},
		{"Romanian", "ro"},
		{"Russian", "ru"},
		{"Serbian", "sr"},
		{"Slovak", "sk"},
		{"Slovenian", "sl"},
		{"Spanish", "es"},
		{"Spanish (ARG)", "es_AR"},
		{"Spanish (ESP)", "es_ES"},
		{"Spanish (MEX)", "es_MX"},
		{"Swahili", "sw"},
		{"Swedish", "sv"},
		{"Tamil", "ta"},
		{"Telugu", "te"},
		{"Thai", "th"},
		{"Turkish", "tr"},
		{"Ukrainian", "uk"},
		{"Urdu", "ur"},
		{"Uzbek", "uz"},
		{"Vietnamese", "vi"},
		{"Zulu", "zu"},
	}

	for _, tc := range testCases {
		actualCode, exists := Languages[tc.languageName]
		if !exists {
			t.Errorf("Código de linguagem para '%s' não encontrado", tc.languageName)
			continue
		}
		if actualCode != tc.expectedLanguageCode {
			t.Errorf("Código de linguagem incorreto para '%s'. Esperado: %s, Obtido: %s", tc.languageName, tc.expectedLanguageCode, actualCode)
		}
	}
}
