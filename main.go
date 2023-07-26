package main

import (
	"fmt"
	"math"
	"time"
)

type Agent struct {
	Name            string
	CurrentLocation Point
}

type Point struct {
	X int
	Y int
}

var agents = []Agent{
	{
		Name:            "A",
		CurrentLocation: Point{},
	},
	{
		Name:            "B",
		CurrentLocation: Point{},
	},
}

var points = []Point{

	{
		X: 3,
		Y: 6,
	},
	{
		X: 8,
		Y: 9,
	},
	{
		X: 2,
		Y: 3,
	},
}

func main() {

Loop:
	for {
		for _, agent := range agents {
			// all point reached
			if len(points) == 0 {
				fmt.Println("There is no point to traverse", "Good Bye!")
				break Loop
			}

			nearest := 0
			for i := 1; i < len(points); i++ {
				if agent.CurrentLocation.DistanceToPoint(points[i]) > agent.CurrentLocation.DistanceToPoint(points[nearest]) {
					nearest = i
				}
			}
			agent.Navigate(points[nearest])
			points = DropAPoint(points, nearest)
		}
	}
}

func (a *Agent) Navigate(p Point) {
	time.Sleep(time.Millisecond * 300)
	a.CurrentLocation = p
	fmt.Printf("Agent:%s went to X:%d,Y:%d\n", a.Name, p.X, p.Y)
}

func (a *Agent) String() string {
	return fmt.Sprintf("AgentName : %s (x : %d, y : %d)", a.Name, a.CurrentLocation.X, a.CurrentLocation.Y)
}

func (p *Point) DistanceToPoint(point Point) int {
	a := p.X - point.X
	b := p.Y - point.Y
	return int(math.Max(float64(a), float64(b)))

}

func DropAPoint(s []Point, i int) []Point {
	// bring element to remove at the end if it's not there yet
	if i != len(s)-1 {
		s[i] = s[len(s)-1]
	}

	// drop the last element
	return s[:len(s)-1]
}
