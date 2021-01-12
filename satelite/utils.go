package satelite

import (
	"math"
)

// CirclesIntersect finds if two circles intersect ...
func CirclesIntersect(s1 Satelite, s2 Satelite, r1 float32, r2 float32) bool {
	dx := float64(s1.X - s2.X)
	dy := float64(s1.Y - s2.Y)

	distanceBetweenCenters := math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))

	radioDifference := math.Abs(float64(r1 - r2))

	radioSum := float64(r1 + r2)

	return distanceBetweenCenters == radioSum ||
		distanceBetweenCenters == radioDifference ||
		(radioDifference < distanceBetweenCenters && distanceBetweenCenters < radioSum)
}

// FindIntersectionPoints find coordinates of points where two circles intersect.
// Original from: http://paulbourke.net/geometry/circlesphere/
// and http://paulbourke.net/geometry/circlesphere/tvoght.c
func FindIntersectionPoints(s1 Satelite, s2 Satelite, r1 float32, r2 float32) (x, y, xPrime, yPrime int) {
	x1 := float64(s1.X)
	y1 := float64(s1.Y)

	dx := float64(s2.X - s1.X)
	dy := float64(s2.Y - s1.Y)

	r1Squared := math.Pow(float64(r1), 2)
	r2Squared := math.Pow(float64(r2), 2)

	distanceBetweenCenters := math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))

	/* 'point 2' is the point where the line through the circle
	* intersection points crosses the line between the circle
	* centers.
	 */
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
	xi := int(ax + rx)
	xiPrime := int(ax - rx)
	yi := int(ay + ry)
	yiPrime := int(ay - ry)
	return xi, yi, xiPrime, yiPrime
}
