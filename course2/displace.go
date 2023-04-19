package course2

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, v0, s0 *float64) func(*float64) float64 {
	fn := func(t *float64) float64 {
		return (0.5)*(*a)*math.Pow(*t, 2) + (*v0)*(*t) + *s0
	}
	return fn
}

func GenDisplaceFnRunner() {
	var a, v0, s0, t float64
	fmt.Println("Enter space separated values for acceleration, initial velocity, and initial displacement. eg: 10 2 1")
	n, err := fmt.Scanf("%f %f %f", &a, &v0, &s0)
	fmt.Println(a, v0, s0)
	if err != nil || n != 3 {
		fmt.Println("All values not entered or read err!")
		panic(err)
	}
	var dispFn func(*float64) float64
	dispFn = GenDisplaceFn(&a, &v0, &s0)
	fmt.Println("Enter time for displacement computation: eg 3")
	_, err = fmt.Scanln(&t)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Displacement for the entered time value: %f is %f:\n", t, dispFn(&t))
}
