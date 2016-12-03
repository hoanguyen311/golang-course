package myMath

func CreateAdd(y int) func(x int) int {
	return func(x int) int {
		return x + y;
	}
}
