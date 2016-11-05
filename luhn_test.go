package luhn

import "testing"

func TestVerify(t *testing.T) {
	res := verifyNumeric("79927398713")
	if res == false {
		t.Error("verifyNumeric failed.")
	}
}

func TestCompute(t *testing.T) {
	tests := []struct {
		Num         int
		Input       string
		Expect      string
		ExpectError bool
	}{
		{1, "7992739871", "3", false},
		{2, "411111111111111", "1", false},
	}

	for _, test := range tests {
		res, err := Compute(test.Input)
		if res != test.Expect {
			t.Logf("Test %v: Expected %#v, got %#v.", test.Num, test.Expect, res)
			t.Fail()
			continue
		}
		if err != nil && test.ExpectError == false {
			t.Logf("Test %v: Expected no error, got: %v", test.Num, err)
			t.Fail()
			continue
		}
		if err == nil && test.ExpectError {
			t.Logf("Test %v: Expected error but did not get one.", test.Num)
			t.Fail()
			continue
		}
		t.Logf("Test %v: Success.", test.Num)
	}
}

func TestCheck(t *testing.T) {
	tests := []struct {
		Num         int
		Input       string
		Expect      bool
		ExpectError bool
	}{
		{1, "79927398713", true, false},
		{2, "4111111111111111", true, false},
	}

	for _, test := range tests {
		res, err := Check(test.Input)
		if res != test.Expect {
			t.Logf("Test %v: Expected %#v, got %#v.", test.Num, test.Expect, res)
			t.Fail()
			continue
		}
		if err != nil && test.ExpectError == false {
			t.Logf("Test %v: Expected no error, got: %v", test.Num, err)
			t.Fail()
			continue
		}
		if err == nil && test.ExpectError {
			t.Logf("Test %v: Expected error but did not get one.", test.Num)
			t.Fail()
			continue
		}
		t.Logf("Test %v: Success.", test.Num)
	}
}
