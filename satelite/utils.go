package satelite

import (
	"math"
)

// CirclesTouch ...
func CirclesTouch(s1 Satelite, s2 Satelite, r1 float32, r2 float32) bool {
	dx := float64(s1.X - s2.X)
	dy := float64(s1.Y - s2.Y)

	distanceBetweenCenters := math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))

	println("distanceBetweenCenters = %v", distanceBetweenCenters)

	radioDifference := math.Abs(float64(r1 - r2))

	radioSum := float64(r1 + r2)

	return distanceBetweenCenters == radioSum ||
		distanceBetweenCenters == radioDifference ||
		(radioDifference < distanceBetweenCenters && distanceBetweenCenters < radioSum)
}

func FindIntersectionPoints(s1 Satelite, s2 Satelite, r1 float32, r2 float32) (x, y, xPrime, yPrime float64) {
	x1 := s1.X
	y1 := s1.Y

	dx := float64(s2.X - s1.X)
	dy := float64(s2.Y - s1.Y)

	r1Squared := math.Pow(float64(r1), 2)
	r2Squared := math.Pow(float64(r2), 2)

	distanceBetweenCenters := math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
	/* Determine the distance from point 0 to point 2. */
	a := ((r1Squared) - (r2Squared) + (math.Pow(distanceBetweenCenters, 2))) / (2.0 * distanceBetweenCenters)

	/* Determine the coordinates of point 2. */
	ax := x1 + (dx * a / distanceBetweenCenters)
	ay := y1 + (dy * a / distanceBetweenCenters)

	/* Determine the distance from point 2 to either of the
	* intersection points.
	 */

	aSquared := math.Pow(a, 2)
	h := math.Sqrt(r1Squared - aSquared)

	/* Now determine the offsets of the intersection points from
	* point 2.
	 */
	rx := -dy * (h / distanceBetweenCenters)
	ry := dx * (h / distanceBetweenCenters)

	/* Determine the absolute intersection points. */
	xi := ax + rx
	xiPrime := ax - rx
	yi := ay + ry
	yiPrime := ay - ry
	return xi, yi, xiPrime, yiPrime
}
