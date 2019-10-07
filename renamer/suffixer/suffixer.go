package suffixer

// Suffixer parses a fileName and extracts its nb and its base name
type Suffixer interface {
	Extract(fileName string) (base, ext string, nb int)
}
