package language

// TextDirection represents layout orientation.
type TextDirection string

const (
	DirectionLTR TextDirection = "LTR"
	DirectionRTL TextDirection = "RTL"
)

// Language represents the document languages supported.
type Language string

const (
	LangEnglish Language = "English"
	LangUrdu    Language = "Urdu"
	LangArabic  Language = "Arabic"
	LangHindi   Language = "Hindi"
	LangMarathi Language = "Marathi"
)

// FontConfig configures the typeface details.
type FontConfig struct {
	Family   string
	Style    string // "", "B", "I", "BI"
	FilePath string // Absolute or relative path to the .ttf font file
}

// LanguageConfig aggregates layout and font properties per language.
type LanguageConfig struct {
	Name      Language
	Direction TextDirection
	Font      FontConfig
	Locale    string
}

// GetDefaultConfig returns predefined settings for a supported language.
func GetDefaultConfig(lang Language) LanguageConfig {
	switch lang {
	case LangUrdu:
		return LanguageConfig{
			Name:      LangUrdu,
			Direction: DirectionRTL,
			Font: FontConfig{
				Family:   "NotoNastaliqUrdu",
				Style:    "",
				FilePath: "/System/Library/Fonts/NotoNastaliq.ttc", // Fallback, overrideable
			},
			Locale: "ur",
		}
	case LangArabic:
		return LanguageConfig{
			Name:      LangArabic,
			Direction: DirectionRTL,
			Font: FontConfig{
				Family:   "Amiri",
				Style:    "",
				FilePath: "/System/Library/Fonts/SFArabic.ttf", // Fallback, overrideable
			},
			Locale: "ar",
		}
	case LangHindi:
		return LanguageConfig{
			Name:      LangHindi,
			Direction: DirectionLTR,
			Font: FontConfig{
				Family:   "NotoSansDevanagari",
				Style:    "",
				FilePath: "/System/Library/Fonts/Supplemental/Devanagari Sangam MN.ttc", // Fallback, overrideable
			},
			Locale: "hi",
		}
	case LangMarathi:
		return LanguageConfig{
			Name:      LangMarathi,
			Direction: DirectionLTR,
			Font: FontConfig{
				Family:   "NotoSansDevanagari",
				Style:    "",
				FilePath: "/System/Library/Fonts/Supplemental/Devanagari Sangam MN.ttc", // Fallback, overrideable
			},
			Locale: "mr",
		}
	default: // Default is English LTR
		return LanguageConfig{
			Name:      LangEnglish,
			Direction: DirectionLTR,
			Font: FontConfig{
				Family:   "Helvetica",
				Style:    "",
				FilePath: "", // Uses PDF standard core font
			},
			Locale: "en",
		}
	}
}
