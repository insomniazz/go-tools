package math

import (
	"fmt"
	"testing"
)

func TestMyRoundFloat(t *testing.T) {
	num := 0.15807659924030304
	x := MyRoundFloat(num, 5)
	y := MyRoundFloat2(num, 5)
	fmt.Printf("%T:%#[1]v\n", x)
	fmt.Printf("%T:%#[1]v\n", y)
}
