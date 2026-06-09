package main

func RectPerimeter(w, h int) int {
	if w < 0 || h < 0 {
		return -1
	}

	perimeter := 2 * (w + h)
	return perimeter
}
