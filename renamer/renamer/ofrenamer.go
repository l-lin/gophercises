package renamer

import (
	"fmt"
	"path/filepath"
)

// OfRenamer renames files in (X of X)
type OfRenamer struct {
	Max int
}

// NewOfRenamer instanciates a new renamer with pattern (X of X)
func NewOfRenamer(max int) *OfRenamer {
	return &OfRenamer{max}
}

// Rename the given fileName to something like "fileName (X of X).txt"
func (r *OfRenamer) Rename(nb int, fileName string) string {
	ext := filepath.Ext(fileName)
	fileNameWithoutExt := string([]rune(fileName)[0 : len(fileName)-len(ext)])
	if fileNameWithoutExt == "" {
		return fmt.Sprintf("(%d of %d)%s", nb, r.Max, ext)
	}
	return fmt.Sprintf("%s (%d of %d)%s", fileNameWithoutExt, nb, r.Max, ext)
}
