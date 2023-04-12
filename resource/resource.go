package resource

type Point struct {
	point int
}

func NewPoint() *Point {
	return &Point{}
}

func (p *Point) AddPoint(point int) {
	p.point += point
}

func (p *Point) Point() int {
	return p.point
}
