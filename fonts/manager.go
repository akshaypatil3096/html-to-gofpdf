package fonts

import (
	"fmt"
	"html-to-gofpdf/language"
	"strings"
	"sync"

	"github.com/jung-kurt/gofpdf"
)

// IsCoreFont returns true if the font family name is one of the built-in PDF core typefaces.
func IsCoreFont(family string) bool {
	switch strings.ToLower(family) {
	case "arial", "helvetica", "times", "courier", "symbol", "zapfdingbats":
		return true
	}
	return false
}

// FontManager handles font registration, fallback mapping, and path configuration.
type FontManager struct {
	mu            sync.Mutex
	registered    map[string]bool
	customConfigs map[language.Language]language.FontConfig
}

// NewFontManager returns a new FontManager.
func NewFontManager() *FontManager {
	return &FontManager{
		registered:    make(map[string]bool),
		customConfigs: make(map[language.Language]language.FontConfig),
	}
}

// RegisterCustomFont registers a custom font configuration for a language.
func (fm *FontManager) RegisterCustomFont(lang language.Language, cfg language.FontConfig) {
	fm.mu.Lock()
	defer fm.mu.Unlock()
	fm.customConfigs[lang] = cfg
}

// GetFontConfig retrieves the font configuration for a language.
func (fm *FontManager) GetFontConfig(lang language.Language) language.FontConfig {
	fm.mu.Lock()
	defer fm.mu.Unlock()
	if cfg, ok := fm.customConfigs[lang]; ok {
		return cfg
	}
	return language.GetDefaultConfig(lang).Font
}

// LoadFontForStyle registers the configured font with the gofpdf instance for the requested style if not already done.
// Returns the registered font family name.
func (fm *FontManager) LoadFontForStyle(pdf *gofpdf.Fpdf, lang language.Language, style string) (string, error) {
	cfg := fm.GetFontConfig(lang)
	if cfg.Family == "" || cfg.FilePath == "" || IsCoreFont(cfg.Family) {
		if strings.ToLower(cfg.Family) == "arial" {
			return "Helvetica", nil
		}
		return cfg.Family, nil
	}

	// gofpdf does not natively support TrueType Collections (.ttc). Fall back to Helvetica core font.
	if strings.HasSuffix(strings.ToLower(cfg.FilePath), ".ttc") {
		return "Helvetica", nil
	}

	key := fmt.Sprintf("%s-%s", cfg.Family, style)
	fm.mu.Lock()
	defer fm.mu.Unlock()

	if !fm.registered[key] {
		// Use AddUTF8Font to load the font file directly.
		// Register under the requested style so SetFont can find it.
		pdf.AddUTF8Font(cfg.Family, style, cfg.FilePath)
		fm.registered[key] = true
	}

	return cfg.Family, nil
}
