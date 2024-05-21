package heuristics

import "testing"

func isEqualStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

type readNumTestCase struct {
	Formula        string
	ExpectedResult string
}

func TestReadNum(t *testing.T) {
	cases := []readNumTestCase{
		{Formula: "12+3435", ExpectedResult: "12"},
		{Formula: "15.2352*43dsgs5", ExpectedResult: "15.2352"},
	}
	for _, testCase := range cases {
		result := readNum(testCase.Formula)
		if result != testCase.ExpectedResult {
			t.Errorf("Wrong Num reading in \"%s\", expected: \"%s\", got: \"%s\"", testCase.Formula, testCase.ExpectedResult, result)
		}
	}
}

type successReadingTestCase struct {
	Formula        string
	ExpectedResult []string
}

func TestSuccessReading(t *testing.T) {
	cases := []successReadingTestCase{
		{Formula: "1+ 2", ExpectedResult: []string{"1", "2", "+"}},
		{Formula: "1 +2* 3", ExpectedResult: []string{"1", "2", "3", "*", "+"}},
		{Formula: "(1 +2)* 3", ExpectedResult: []string{"1", "2", "+", "3", "*"}},
		{Formula: "abs(2-3) * 3", ExpectedResult: []string{"2", "3", "-", "abs", "3", "*"}},
		{Formula: "min(abs(source.x-target.x), abs(source.y-target.y))", ExpectedResult: []string{"source.x", "target.x", "-", "abs", "source.y", "target.y", "-", "abs", "min"}},
	}
	for _, testCase := range cases {
		result := readExpression(testCase.Formula)
		if !isEqualStringSlices(result, testCase.ExpectedResult) {
			t.Errorf("Wrong read in \"%s\", expected: \"%#v\", got: \"%#v\"", testCase.Formula, testCase.ExpectedResult, result)
		}
	}
}

type isFirstNumTestCase struct {
	Formula        string
	ExpectedResult bool
}

func TestIsFirstNum(t *testing.T) {
	cases := []isFirstNumTestCase{
		{Formula: "12+1+2+3", ExpectedResult: true},
		{Formula: "max(12+1,2+3)", ExpectedResult: false},
		{Formula: "(12+1-2*3)", ExpectedResult: false},
		{Formula: "source.x+target.x", ExpectedResult: true},
	}
	for _, testCase := range cases {
		result := isNextNum(testCase.Formula)

		if result != testCase.ExpectedResult {
			t.Errorf("Wrong Num reading in \"%s\". Expected: %t, got %t", testCase.Formula, testCase.ExpectedResult, result)
		}
	}
}

type fullTestCase struct {
	Formula        string
	ExpectedResult float64
	Source         Params
	Target         Params
}

func TestFull(t *testing.T) {
	dummyParams := map[string]interface{}{}
	cases := []fullTestCase{
		{Formula: "12+1+2+3", ExpectedResult: 18.0, Source: dummyParams, Target: dummyParams},
		{Formula: "12+1-2+3", ExpectedResult: 14.0, Source: dummyParams, Target: dummyParams},
		{Formula: "12+1-2*3", ExpectedResult: 7.0, Source: dummyParams, Target: dummyParams},
		{Formula: "12+1-3/2", ExpectedResult: 11.5, Source: dummyParams, Target: dummyParams},
		{Formula: "12+1-3^2", ExpectedResult: 4.0, Source: dummyParams, Target: dummyParams},
		{Formula: "source.x+target.x", ExpectedResult: 3.0, Source: map[string]interface{}{"x": 1}, Target: map[string]interface{}{"x": 2}},
		{Formula: "abs(source.x-target.x)", ExpectedResult: 1.0, Source: map[string]interface{}{"x": 1}, Target: map[string]interface{}{"x": 2}},
		{Formula: "max(abs(source.x-target.x), abs(source.y-target.y))", ExpectedResult: 2.0, Source: map[string]interface{}{"x": 1, "y": 0}, Target: map[string]interface{}{"x": 2, "y": 2}},
		{Formula: "min(abs(source.x-target.x), abs(source.y-target.y))", ExpectedResult: 1.0, Source: map[string]interface{}{"x": 1, "y": 0}, Target: map[string]interface{}{"x": 2, "y": 2}},
		{Formula: "(source.x - target.x)^2+(source.y-target.y)^2", ExpectedResult: 25.0, Source: map[string]interface{}{"x": 1, "y": 1}, Target: map[string]interface{}{"x": 4, "y": 5}},
	}
	for _, testCase := range cases {
		heuristic := GetHeuristic(testCase.Formula)
		result, err := heuristic(testCase.Source, testCase.Target)
		if err != nil {
			t.Errorf("Error: %s", err)
		}

		if result != testCase.ExpectedResult {
			t.Errorf("Wrong Num reading in \"%s\". Expected: %g, got %g", testCase.Formula, testCase.ExpectedResult, result)
		}
	}
}
