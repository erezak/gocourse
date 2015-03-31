package calculator

type Calculator struct {
}

func (c Calculator) Add(first float64, second float64) float64 {
	return first + second
}

func (c Calculator) Substract(first float64, second float64) float64 {
	return first - second
}

func (c Calculator) Multiply(first float64, second float64) float64 {
	return first * second
}

func (c Calculator) Divide(first float64, second float64) (float64, error) {
	var err error
	var answer float64
	if second != 0 {
		answer = first / second
	}
	return answer, err
}
