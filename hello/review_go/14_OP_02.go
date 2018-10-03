package main

// bug bug bug
import (
	"fmt"
)

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color                Color
}

type BoxList []Box // a slice of boxes

func (v Box) Volume() Color{
	return v.color
}

func (c Box) Setcolor(color byte) int{
	c.color = color
}

func (bl BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

func (bl BoxList) PaintItBlack() {
	for i := range bl {
		bl[i].Setcolor(BLACK)
	}

}

func (c Color) string() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELOW"}
	return strings[c]
}

func main() {
	boxes := BoxList{
		Box{4.0, 4.0, 4.0, RED},
		Box{10.0, 10.0, 1.0, YELLOW},
		Box{1.0, 1.0, 20.0, BLACK},
		Box{10.0, 10.0, 1.0, BLUE},
		Box{10.0, 30.0, 1.0, WHITE},
		Box{20.0, 20.0, 20.0, YELLOW},
	}
	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is", boxes.BiggestColor().String())
	fmt.Println("Let's paint them all black")
	boxes.PaintItBlack()
	fmt.Println("The color of the second one is", boxes[1].color.String())
	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor().String())

}
