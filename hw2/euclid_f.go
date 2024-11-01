package hw2

import (
	"errors"
	"fmt"
)

type EuclidArg struct {
	B1  int
	B2  int
	B   int
	Mod int
}

func Euclid(input any) (string, error) {
	out := ""
	if in, ok := input.(EuclidArg); ok {
		gcd, _, _ := ExtendedGCD(in.B1, in.B2)
		out += fmt.Sprintf("gcd of %v and %v is %v \n", in.B1, in.B2, gcd)
		imod, err := Inverse(in.B, in.Mod)
		if err != nil {
			out += fmt.Sprintf("failed to get inverse mod of %v and %v\n", in.B1, in.B2)
		} else {
			out += fmt.Sprintf("inverse mod of %v and %v is %v \n", in.B, in.Mod, imod)
		}
	} else {
		return "", errors.New("expected euclid arg as input")
	}
	return out, nil
}

func ExtendedGCD(a, b int) (int, int, int) {
	if a == 0 {
		return b, 0, 1
	}
	gcd, x1, y1 := ExtendedGCD(b%a, a)
	x := y1 - (b/a)*x1
	y := x1
	return gcd, x, y
}

func Inverse(a, b int) (int, error) {
	gcd, x, _ := ExtendedGCD(a, b)
	if gcd != 1 || b == 0 {
		return 0, errors.New("no inverse mod found")
	}

	return ((b + (x % b)) % b), nil
}
