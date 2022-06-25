package pow_50

import "testing"

func myPow(x float64, n int) float64 {

	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1.0 / myPow(x, -n)
	}

	ans := myPow(x, n/2)
	ans = ans * ans

	if n%2 == 1 {
		return ans * x
	}
	return ans
}

func TestMyPow(t *testing.T) {
	t.Log(myPow(2.00000, 2))
	t.Log(myPow(2.10000, 3))
	t.Log(myPow(2.00000, -2))
}
