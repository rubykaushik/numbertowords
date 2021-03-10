package numbertowords

import "testing"

func TestConvert(t *testing.T) {
	result, err := Convert(-1)
	if result != "" || err == nil {
		t.Fatal("test failed for -1")
	}

	result, err = Convert(100000)
	if result != "" || err == nil {
		t.Fatal("test failed for 100000")
	}
}

func TestUnits(t *testing.T) {
	testCases := map[int]string{
		0: "zero",
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}

	for n, word := range testCases {
		result, err := Convert(n)
		if result != word || err != nil {
			t.Fail()
		}
	}
}
