//go:generate stringer -type=Mode
package primitive

// Mode of the primitive to transform
type Mode int

const (
	// Combo mode
	Combo Mode = iota
	// Triangle mode
	Triangle
	// Rect mode
	Rect
	// Ellipse mode
	Ellipse
	// Circle mode
	Circle
	// Rotatedrect mode
	Rotatedrect
	// Beziers mode
	Beziers
	// Rotatedellipse mode
	Rotatedellipse
	// Polygon mode
	Polygon
)

// Modes contains all the possible primitive modes
var Modes = [...]Mode{Combo, Triangle, Rect, Beziers}

// Too many modes, so we are reducing it to have something faster
//var Modes = [...]Mode{Combo, Triangle, Rect, Ellipse, Circle, Rotatedrect, Beziers, Rotatedellipse, Polygon}

// ToSlice transforms the modes into a slice
func ToSlice() []string {
	result := []string{}
	for _, m := range Modes {
		result = append(result, m.String())
	}
	return result
}
