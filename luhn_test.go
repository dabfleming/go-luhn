package luhn

import "testing"

func TestCompute(t *testing.T) {
	tests := []struct {
		Num         int
		Input       string
		Expect      string
		ExpectError bool
	}{
		{1, "7992739871", "3", false},
		{2, "411111111111111", "1", false},
		{3, "455592949368433", "7", false},
		{4, "432877535468714", "1", false},
		{5, "420493863322", "3", false},
		{6, "434969652705", "7", false},
		{7, "529718802031232", "2", false},
		{8, "550967024970065", "4", false},
		{9, "37354557488626", "6", false},
		{10, "36223376742346", "9", false},
		{11, "601138952562565", "4", false},
		{12, "601195990946658", "4", false},
		{13, "353 345", "", true},
		{14, "357a635", "", true},
		{15, "346\t343", "", true},
		{16, "36,463", "", true},
		{17, "36475.575", "", true},
		{18, "36\n45", "", true},
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
		{3, "4555929493684337", true, false},
		{4, "4328775354687141", true, false},
		{5, "4204938633223", true, false},
		{6, "4349696527057", true, false},
		{7, "5297188020312322", true, false},
		{8, "5509670249700654", true, false},
		{9, "373545574886266", true, false},
		{10, "362233767423469", true, false},
		{11, "6011389525625654", true, false},
		{12, "6011959909466584", true, false},
		{13, "4555929493684336", false, false},
		{14, "4328775354687144", false, false},
		{15, "4204938633225", false, false},
		{16, "4349696527052", false, false},
		{17, "5297188020312328", false, false},
		{18, "5509670249700658", false, false},
		{19, "373545574886267", false, false},
		{20, "362233767423463", false, false},
		{21, "6011389525625658", false, false},
		{22, "6011959909466586", false, false},
		{23, "353 345", false, true},
		{24, "357a635", false, true},
		{25, "346\t343", false, true},
		{26, "36,463", false, true},
		{27, "36475.575", false, true},
		{28, "36\n45", false, true},
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
