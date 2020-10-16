package main

func Burbuja(s []int64) {
	vueltas := cap(s) - 1
	aux := int64(0)

	for i := 0; i < vueltas; i++ {
		for j := 0; j < vueltas; j++ {
			if s[j] > s[j+1] {
				aux = s[j]
				s[j] = s[j+1]
				s[j+1] = aux
			}
		}
	}
}

func main() {
	slide := []int64{5, 4, 3, 2, 1}
	Burbuja(slide)
}
