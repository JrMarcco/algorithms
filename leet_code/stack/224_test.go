package stack

import (
	"strconv"
	"testing"
)

func calculate(s string) int {
	var ops []string           // 操作符栈
	var inversePoland []string // 逆波兰式

	number := ""     // 参与计算数字
	needZero := true // 通过补 0 来表示数字前的正负号，正负号前为左括号或者正负号开头的情况需要补 0
	for _, ch := range s {
		if ch == ' ' { // 过滤空格
			continue
		}

		if ch >= '0' && ch <= '9' {
			number += string(ch)
			needZero = false
			continue
		}
		if number != "" {
			// 碰到操作符则将前面记录的数字添加的逆波兰式尾部
			inversePoland = append(inversePoland, number)
			number = ""
		}

		if ch == '(' {
			// 碰到左括号则压栈
			ops = append(ops, string(ch))
			needZero = true
			continue
		}
		if ch == ')' {
			// 碰到右括号，则栈内最近的一个左括号之间的操作符添加到逆波兰式尾部
			for {
				topOp := ops[len(ops)-1]
				if topOp == "(" {
					break
				}
				inversePoland = append(inversePoland, topOp)
				ops = ops[:len(ops)-1]
			}
			ops = ops[:len(ops)-1]
			needZero = false
			continue
		}

		if (ch == '+' || ch == '-') && needZero {
			inversePoland = append(inversePoland, "0")
		}
		opWeight := getOpWeight(string(ch))
		for {
			if len(ops) == 0 {
				break
			}
			topOp := ops[len(ops)-1]
			if getOpWeight(topOp) < opWeight {
				break
			}
			// 栈顶原算法权重更大，应该先进行计算，则出栈放入逆波兰式尾部
			inversePoland = append(inversePoland, topOp)
			ops = ops[:len(ops)-1]
		}
		ops = append(ops, string(ch))
	}
	if number != "" {
		inversePoland = append(inversePoland, number)
	}
	for {
		if len(ops) == 0 {
			break
		}
		inversePoland = append(inversePoland, ops[len(ops)-1])
		ops = ops[:len(ops)-1]
	}
	return calcInversePoland(inversePoland)
}

// getOpWeight 获取操作符权重
func getOpWeight(op string) int {
	if op == "+" || op == "-" {
		return 1
	}
	if op == "*" || op == "/" {
		return 2
	}
	return 0
}

// calcInversePoland 对逆波兰式求值
func calcInversePoland(inversePoland []string) int {
	var res []int
	for i := 0; i < len(inversePoland); i++ {
		if inversePoland[i] == "+" || inversePoland[i] == "-" || inversePoland[i] == "*" || inversePoland[i] == "/" {
			a := res[len(res)-2]
			b := res[len(res)-1]
			res = res[:len(res)-2]

			res = append(res, execCalc(a, b, inversePoland[i]))
			continue
		}
		val, _ := strconv.Atoi(inversePoland[i])
		res = append(res, val)
	}
	return res[len(res)-1]
}

func execCalc(a int, b int, op string) int {
	if op == "+" {
		return a + b
	}
	if op == "-" {
		return a - b
	}
	if op == "*" {
		return a * b
	}
	if op == "/" {
		return a / b
	}
	return 0
}

func TestCalc(t *testing.T) {
	t.Logf("%d", calculate("1 + 2 * 3 - (4 + 5 * 6) "))
}
