package renamer

// Renamer renames files according to a specific naming pattern
type Renamer interface {
	Rename(nb int, fileName string) string
}
