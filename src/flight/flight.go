package flight

import(
	"math"
	"graph"
)

type Time struct {
	h	int
	m	int
	s	int
}

/**
 * This function calculates the flight statistics, given a particular path
 * @param g - the graph on which the path was calculated
 * @param path - the path of nodes through the graph
 * @return h, m, s - representing hours, minutes, and seconds of flight and overlay
 * @return sum - the cost of the flight path
 */
func CalculateFlightStat(g *graph.Graph, path []graph.Node) (int, int, int, float64) {

	var edge graph.Edge

	time := 0.0
	cost := 0.35
	sum  := 0.0

	for j := 0; j < len(path) - 1; j++ {

		if j != 0 && j != len(path) - 1 {
			time += (7200.0 - (float64(len(path[j].OutEdges)) * 600.0))
		}

		for i := 0; i < len(path[j].OutEdges); i++ {
			edge = path[j].OutEdges[i]
			if edge.Destination.Code == path[j + 1].Code { break }
		}

		distance := float64(edge.Len)
		a := float64(750.0 / 3600.0)

		sum = distance * cost
		if cost > 0 { cost -= 0.05 }

		if distance > 400 {
			time += 2.0 * math.Sqrt((400.0) / a)
			distance -= 400
			time += ((distance / 750.0) * 3600.0)
		}else{
			distance /= 2
			time += 2.0 * math.Sqrt((2.0 * distance) / a)
		}
	}

	t := ConvertSecToTime(time)

	return t.h, t.m, t.s, sum
}

/**
 * Converts the t to the time struct to {hours, minutes, seconds}
 * @param t - time struct
 */
func ConvertSecToTime(t float64) Time{

	hours := int(t / 3600.0)

	t = (t / 3600) - float64(hours)
	min   := int(t * 60)

	t = (t * 60) - float64(min)
	sec   := int(t * 60)

	new_t := Time{hours, min, sec}

	return new_t
}
