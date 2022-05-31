package math

import (
	"fmt"
	"testing"
)

func TestMyRoundFloat(t *testing.T) {
	x := MyRoundFloat(12.15807659924030304,5)
	fmt.Printf("%T:%#v",x,x)
}
