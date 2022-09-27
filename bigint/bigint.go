package bigint

import (
	"errors"
	"strconv"
	"strings"
)

type Bigint struct {
	Value string
}

var ErrorBadInput = errors.New("bad input, please input only number")
var ErrorRange = errors.New("out of range, maximum number - 2147483647(int32)")
var ErrorDivZero = errors.New("error divide by zero")

func validNum(num string) (string, error) {
	allowed := "1234567890"
	var check bool

	for strings.HasPrefix(num, "+") {
		num = strings.TrimPrefix(num, "+")
	}

	valStr := num
	if (strings.HasPrefix(num, "-") && strings.Count(num, "-") == 1) || (strings.HasPrefix(num, "+") && strings.Count(num, "-") == 1) {
		valStr = num[1:]
	}

	arr := strings.Split(valStr, "")

	for _, v := range arr {
		if !strings.Contains(allowed, v) {
			check = true
		}
	}

	if check {
		return num, ErrorBadInput
	}

	_, err := strconv.ParseInt(num, 10, 32)
	if err != nil {
		panic(ErrorRange)
	}

	return num, nil
}

func cleanUp(num string) string {
	if strings.HasPrefix(num, "-") {
		num = strings.Replace(num, "-", "", 1)
		for strings.HasPrefix(num, "0") {
			num = strings.TrimPrefix(num, "0")
		}
		num = "-" + num
	}

	for strings.HasPrefix(num, "0") {
		num = strings.TrimPrefix(num, "0")

	}

	return num
}

func NewInt(num string) (Bigint, error) {
	num, err := validNum(num)
	if err != nil {
		return Bigint{}, err
	}

	num = cleanUp(num)

	if num == "" {
		return Bigint{Value: "0"}, nil
	}

	return Bigint{Value: num}, nil
}

func (z *Bigint) Set(num string) error {
	num, err := validNum(num)
	if err != nil {
		return err
	}

	num = cleanUp(num)

	if num == "" {
		return nil
	}

	z.Value = num

	return nil
}

func Add(a, b Bigint) Bigint {
	if a.Value[0] != '-' && b.Value[0] == '-' {
		b.Value = b.Value[1:]
		res := Sub(a, b)
		return res
	}
	if a.Value[0] == '-' && b.Value[0] != '-' {
		a.Value = a.Value[1:]
		num1,_ := strconv.Atoi(a.Value)
		num2,_ := strconv.Atoi(b.Value)
		var res Bigint
		if len(a.Value) <= len(b.Value) {
			num1, num2 = num2, num1
			a := num1 - num2
			rest := strconv.Itoa(a)
			res = Bigint{Value: rest}
		}
		// if len(a.Value) == len(b.Value){

		// }
		if len(a.Value) > len(b.Value) {
			res = Sub(a, b)
			res.Value = "-" + res.Value
		}
		return res
	}
	firstNum, _ := strconv.Atoi(a.Value)
	secondNum, _ := strconv.Atoi(b.Value)

	sum := firstNum + secondNum

	str := strconv.Itoa(sum)

	return Bigint{Value: str}
}

func Sub(a, b Bigint) Bigint {
	if a.Value[0] != '-' && b.Value[0] == '-' {
		b.Value = b.Value[1:]
		return Add(a, b)
	}
	if a.Value[0] == '-' && b.Value[0] != '-' {
		a.Value = a.Value[1:]
		res := Add(a, b)
		res.Value = "-" + res.Value
		return res
	}
	if a.Value[0] == '-' && b.Value[0] == '-' {
		firstNum, _ := strconv.Atoi(a.Value)
		secondNum, _ := strconv.Atoi(b.Value)

		sub := firstNum - secondNum

		str := strconv.Itoa(sub)

		return Bigint{Value: str}
	}

	firstNum, _ := strconv.Atoi(a.Value)
	secondNum, _ := strconv.Atoi(b.Value)
	flag := false
	if firstNum < secondNum {
		firstNum, secondNum = secondNum, firstNum
		flag = true
	}
	sub := firstNum - secondNum
	str := strconv.Itoa(sub)
	if flag {
		str = "-" + str
	}
	return Bigint{Value: str}
}

func Mult(a, b Bigint) Bigint {
	firstNum, _ := strconv.Atoi(a.Value)
	secondNum, _ := strconv.Atoi(b.Value)

	mult := firstNum * secondNum

	str := strconv.Itoa(mult)

	return Bigint{Value: str}
}

func Mod(a, b Bigint) Bigint {
	firstNum, _ := strconv.Atoi(a.Value)
	secondNum, _ := strconv.Atoi(b.Value)

	if firstNum == 0 || secondNum == 0 {
		panic(ErrorDivZero.Error())
	}

	mod := firstNum % secondNum

	str := strconv.Itoa(mod)

	return Bigint{Value: str}
}

func (x Bigint) Abs() Bigint {
	if x.Value[0] == '-' {
		return Bigint{
			Value: x.Value[1:],
		}
	}
	if x.Value[0] == '+' {
		return Bigint{
			Value: x.Value[1:],
		}
	}
	return Bigint{
		Value: x.Value,
	}
}
