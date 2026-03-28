package sync

import (
	_ "embed"
	"encoding/json"
	"strings"
	"sync"
)

//go:embed data/psp_gamedb.json
var pspGameDBData []byte

var (
	pspGameDB     map[string]string
	pspGameDBOnce sync.Once
)

func loadPSPGameDB() map[string]string {
	pspGameDBOnce.Do(func() {
		pspGameDB = make(map[string]string)
		json.Unmarshal(pspGameDBData, &pspGameDB)
	})
	return pspGameDB
}

// LookupPSPTitle resolves a PPSSPP save folder name to a game title.
// Folder names vary in format:
//   - "UCUS98662_GameData0" (underscore-separated suffix)
//   - "ULUS10088010000" (save slot appended directly)
//   - "UCUS98662" (plain Game ID)
//
// The standard Game ID is 4 letters + 5 digits (9 chars). The function
// tries an underscore split first, then extracts the 9-char prefix.
func LookupPSPTitle(folderName string) (string, bool) {
	db := loadPSPGameDB()

	gameID := strings.ReplaceAll(folderName, "-", "")

	// Try underscore split first (e.g., UCUS98662_GameData0)
	if idx := strings.Index(gameID, "_"); idx > 0 {
		if title, ok := db[gameID[:idx]]; ok {
			return title, true
		}
	}

	// Try standard 9-char Game ID prefix (e.g., ULUS10088 from ULUS10088010000)
	if len(gameID) >= 9 {
		if title, ok := db[gameID[:9]]; ok {
			return title, true
		}
	}

	// Try exact match as fallback
	if title, ok := db[gameID]; ok {
		return title, true
	}

	return "", false
}
