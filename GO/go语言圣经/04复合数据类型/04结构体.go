package main
import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}

type Circle2 struct {
	Point //匿名成员
	Radius int
}

type Wheel2 struct {
	Circle2
	Spokes int
}


func main() {
	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 8
	w.Circle.Radius = 5
	w.Spokes = 20

	//可直接访问匿名成员下的元素
	var w2 Wheel2
	w2.X = 8
	w2.Y = 8
	w2.Radius = 5
	w2.Spokes = 20

	//w = Wheel2{8, 8, 5, 20}                       // compile error: unknown fields
	//w = Wheel2{X: 8, Y: 8, Radius: 5, Spokes: 20} // compile error: unknown fields

	fmt.Println(w)
	fmt.Println(w2)

	fmt.Println(Wheel2{Circle2{Point{8, 8}, 5}, 29})

}
