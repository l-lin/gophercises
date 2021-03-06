// Code generated by "stringer -type=Mode"; DO NOT EDIT.

package primitive

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Combo-0]
	_ = x[Triangle-1]
	_ = x[Rect-2]
	_ = x[Ellipse-3]
	_ = x[Circle-4]
	_ = x[Rotatedrect-5]
	_ = x[Beziers-6]
	_ = x[Rotatedellipse-7]
	_ = x[Polygon-8]
}

const _Mode_name = "ComboTriangleRectEllipseCircleRotatedrectBeziersRotatedellipsePolygon"

var _Mode_index = [...]uint8{0, 5, 13, 17, 24, 30, 41, 48, 62, 69}

func (i Mode) String() string {
	if i < 0 || i >= Mode(len(_Mode_index)-1) {
		return "Mode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Mode_name[_Mode_index[i]:_Mode_index[i+1]]
}
