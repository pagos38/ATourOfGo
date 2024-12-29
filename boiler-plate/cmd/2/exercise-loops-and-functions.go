package main

import "fmt"

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		if z2 := z - (z*z-x)/(2*z); z2 == z {
			fmt.Printf("%d回で推定完了\n", i+1)
			break
		} else {
			z = z2
		}
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
