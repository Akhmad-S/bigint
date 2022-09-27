package bigint_test

import (
	"testing"

	"github.com/uacademy/bigint/bigint"
)

func TestNewint(t *testing.T){
	res, _ := bigint.NewInt("0123456789")
	expected := bigint.Bigint{Value: "123456789"}

	if res != expected{
		t.Error("Not valid number")
	}

	res2, _ := bigint.NewInt("+++++++++++123456789")
	expected2 := bigint.Bigint{Value: "123456789"}

	if res2 != expected2{
		t.Error("Not valid number")
	}

	res3, _ := bigint.NewInt("-0003456789")
	expected3 := bigint.Bigint{Value: "-3456789"}

	if res3 != expected3{
		t.Error("Not valid number")
	}
}

func TestSet(t *testing.T){
	a := bigint.Bigint{Value: ""}
	a.Set("-160")
	a.Set("160")
}

func TestSum(t *testing.T){
	testTable := []struct{
		num1 bigint.Bigint
		num2 bigint.Bigint
		expected bigint.Bigint
	}{
		{num1: bigint.Bigint{Value: "0"}, num2: bigint.Bigint{Value: "0"}, expected: bigint.Bigint{Value: "0"}},
		{num1: bigint.Bigint{Value: "50"}, num2: bigint.Bigint{Value: "100"}, expected: bigint.Bigint{Value: "150"}},
		{num1: bigint.Bigint{Value: "-50"}, num2: bigint.Bigint{Value: "-100"}, expected: bigint.Bigint{Value: "-150"}},
		{num1: bigint.Bigint{Value: "-50"}, num2: bigint.Bigint{Value: "100"}, expected: bigint.Bigint{Value: "50"}}, 
		{num1: bigint.Bigint{Value: "50"}, num2: bigint.Bigint{Value: "-100"}, expected: bigint.Bigint{Value: "-50"}},

		{num1: bigint.Bigint{Value: "100"}, num2: bigint.Bigint{Value: "-50"}, expected: bigint.Bigint{Value: "50"}},
		{num1: bigint.Bigint{Value: "-100"}, num2: bigint.Bigint{Value: "50"}, expected: bigint.Bigint{Value: "-50"}},
		{num1: bigint.Bigint{Value: "123"}, num2: bigint.Bigint{Value: "120"}, expected: bigint.Bigint{Value: "243"}},
	}

	for _, tests := range testTable{
		total := bigint.Add(tests.num1, tests.num2)
		if total != tests.expected{
			t.Errorf("Sum was incorrect got: %v, want: %v.", total, tests.expected)
		}
	}
}

func TestSub(t *testing.T){
	testTable := []struct{
		num1 bigint.Bigint
		num2 bigint.Bigint
		expected bigint.Bigint
	}{
		{num1: bigint.Bigint{Value: "0"}, num2: bigint.Bigint{Value: "0"}, expected: bigint.Bigint{Value: "0"}},
		{num1: bigint.Bigint{Value: "50"}, num2: bigint.Bigint{Value: "100"}, expected: bigint.Bigint{Value: "-50"}},
		{num1: bigint.Bigint{Value: "-50"}, num2: bigint.Bigint{Value: "-100"}, expected: bigint.Bigint{Value: "50"}},
		{num1: bigint.Bigint{Value: "-50"}, num2: bigint.Bigint{Value: "100"}, expected: bigint.Bigint{Value: "-150"}},
		{num1: bigint.Bigint{Value: "50"}, num2: bigint.Bigint{Value: "-100"}, expected: bigint.Bigint{Value: "150"}},

		{num1: bigint.Bigint{Value: "100"}, num2: bigint.Bigint{Value: "50"}, expected: bigint.Bigint{Value: "50"}},
		{num1: bigint.Bigint{Value: "100"}, num2: bigint.Bigint{Value: "-50"}, expected: bigint.Bigint{Value: "150"}},
		{num1: bigint.Bigint{Value: "-100"}, num2: bigint.Bigint{Value: "50"}, expected: bigint.Bigint{Value: "-150"}},
		{num1: bigint.Bigint{Value: "-100"}, num2: bigint.Bigint{Value: "-50"}, expected: bigint.Bigint{Value: "-50"}},

		{num1: bigint.Bigint{Value: "123"}, num2: bigint.Bigint{Value: "120"}, expected: bigint.Bigint{Value: "3"}},
	}

	for _, tests := range testTable{
		total := bigint.Sub(tests.num1, tests.num2)
		if total != tests.expected{
			t.Errorf("Subtruct was incorrect got: %v, want: %v.", total, tests.expected)
		}
	}
}

func TestMult(t *testing.T){
	testTable := []struct{
		num1 bigint.Bigint
		num2 bigint.Bigint
		expected bigint.Bigint
	}{
		{num1: bigint.Bigint{Value: "0"}, num2: bigint.Bigint{Value: "0"}, expected: bigint.Bigint{Value: "0"}},
		{num1: bigint.Bigint{Value: "50"}, num2: bigint.Bigint{Value: "100"}, expected: bigint.Bigint{Value: "5000"}},
		{num1: bigint.Bigint{Value: "-50"}, num2: bigint.Bigint{Value: "-100"}, expected: bigint.Bigint{Value: "5000"}},
		{num1: bigint.Bigint{Value: "-50"}, num2: bigint.Bigint{Value: "100"}, expected: bigint.Bigint{Value: "-5000"}},
		{num1: bigint.Bigint{Value: "50"}, num2: bigint.Bigint{Value: "-100"}, expected: bigint.Bigint{Value: "-5000"}},

	}

	for _, tests := range testTable{
		total := bigint.Mult(tests.num1, tests.num2)
		if total != tests.expected{
			t.Errorf("Multiply was incorrect, got: %v, want: %v.", total, tests.expected)
		}
	}
}

func TestMod(t *testing.T){
	testTable := []struct{
		num1 bigint.Bigint
		num2 bigint.Bigint
		expected bigint.Bigint
	}{
		{num1: bigint.Bigint{Value: "50"}, num2: bigint.Bigint{Value: "5"}, expected: bigint.Bigint{Value: "0"}},
		{num1: bigint.Bigint{Value: "50"}, num2: bigint.Bigint{Value: "100"}, expected: bigint.Bigint{Value: "50"}},
		{num1: bigint.Bigint{Value: "35"}, num2: bigint.Bigint{Value: "15"}, expected: bigint.Bigint{Value: "5"}},
		{num1: bigint.Bigint{Value: "-50"}, num2: bigint.Bigint{Value: "3"}, expected: bigint.Bigint{Value: "-2"}},
		{num1: bigint.Bigint{Value: "50"}, num2: bigint.Bigint{Value: "-3"}, expected: bigint.Bigint{Value: "2"}},

	}

	for _, tests := range testTable{
		total := bigint.Mod(tests.num1, tests.num2)
		if total != tests.expected{
			t.Errorf("Mod was incorrect got: %v, want: %v.", total, tests.expected)
		}
	}
}

func TestAbs(t *testing.T){
	res := bigint.Bigint.Abs(bigint.Bigint{Value: "-489"})
	expected := "489"

	if res.Value != expected{
		t.Errorf("Absolute was incorrect, got: %v, want: %v", res, expected)
	}

	res2 := bigint.Bigint.Abs(bigint.Bigint{Value: "+489"})
	expected2 := "489"

	if res2.Value != expected2{
		t.Errorf("Absolute was incorrect, got: %v, want: %v", res2, expected2)
	}

	res3 := bigint.Bigint.Abs(bigint.Bigint{Value: "489"})
	expected3 := "489"

	if res2.Value != expected2{
		t.Errorf("Absolute was incorrect, got: %v, want: %v", res3, expected3)
	}
}