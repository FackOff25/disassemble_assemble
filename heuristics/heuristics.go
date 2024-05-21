package heuristics

import (
	"errors"
	"log"
	"regexp"
	"strings"
)

type Heuristic func(source Params, target Params) (float64, error)

func GetHeuristic(functionString string) Heuristic {
	yard := readExpression(functionString)
	mainStages := yardToStages(yard)

	return func(source Params, target Params) (float64, error) {
		operands := make([]float64, 0)
		stages := make([]stage, len(mainStages))
		copy(stages, mainStages)

		for len(stages) > 0 {
			stage := stages[0]
			stages = stages[1:]
			operands = append([]float64{stage(&operands, source, target)}, operands...)
		}

		return operands[0], nil
	}
}

var (
	sourceMark            = "source"
	targetMark            = "target"
	numRegExp             = `[0-9.]+`
	sourceTargetNumRegExp = `(` + sourceMark + `|` + targetMark + `)\.` + `[a-zA-Z0-9]+`
	funcRegExp            = `(abs|max|min)`
)

func getOperatorsPriorities() map[string]int {
	return map[string]int{
		"+":   0,
		"-":   0,
		"*":   1,
		"/":   1,
		"^":   2,
		"abs": 3,
		"max": 3,
		"min": 3,
		"(":   -1,
	}
}

func readExpression(expression string) (yard []string) {
	expression = strings.ReplaceAll(expression, " ", "")
	yard = []string{}

	priorities := getOperatorsPriorities()
	var operations []string
	for len(expression) > 0 {
		if !isNextNum(expression) {
			oper := expression[0]
			switch oper {
			case '+':
				fallthrough
			case '-':
				fallthrough
			case '*':
				fallthrough
			case '/':
				fallthrough
			case '^':
				if len(operations) > 0 {
					if priorities[operations[0]] >= priorities[string(oper)] {
						yard = append(yard, operations[0])
						operations = operations[1:]
					}
				}
				operations = append([]string{string(oper)}, operations...)
			case ')':
				i := 0
				//fmt.Printf("%#v", operations)
				for ; operations[i] != "("; i++ {
				}
				yard = append(yard, operations[:i]...)
				operations = operations[i+1:]
				if len(operations) > 0 && isNextFunc(operations[0]) {
					yard = append(yard, operations[0])
					operations = operations[1:]
				}
			case '(':
				operations = append([]string{"("}, operations...)
			case ',':
				// ignore
			default:
				fun, err := readFunc(expression)
				if err != nil {
					log.Fatal(err)
				}
				operations = append([]string{fun}, operations...)
				expression = expression[len(fun)-1:]
			}
			expression = expression[1:]
		} else {
			num := readNum(expression)
			yard = append(yard, num)
			expression = expression[len(num):]
		}
		//log.Printf("%s, %#v, %#v", expression, yard, operations)
	}
	yard = append(yard, operations...)
	return
}

func readNum(expression string) (num string) {
	num = ""
	if strings.HasPrefix(expression, sourceMark) || strings.HasPrefix(expression, targetMark) {
		reg := regexp.MustCompile(sourceTargetNumRegExp)
		res := reg.FindAllString(expression, -1)[0]
		num = res
		return
	}
	for idx := range expression {
		if ('0' <= expression[idx] && expression[idx] <= '9') || expression[idx] == '.' {
			num += string(expression[idx])
		} else {
			if !(expression[idx] == '+' || expression[idx] == '-' ||
				expression[idx] == '*' || expression[idx] == '/' ||
				expression[idx] == '^' || expression[idx] == ')') {
				log.Fatal("invalid symbol")
			}
			break
		}
	}
	return
}

func getNumStartRegExp() string {
	return `^(` + numRegExp + `|` +
		sourceTargetNumRegExp + `)`
}

func isNextNum(expression string) bool {
	matched, err := regexp.Match(getNumStartRegExp(), []byte(expression))
	if err != nil {
		panic(err)
	}
	return matched
}

func getFuncStartRegExp() string {
	return `^` + funcRegExp
}

func isNextFunc(expression string) bool {
	matched, err := regexp.Match(getFuncStartRegExp(), []byte(expression))
	if err != nil {
		panic(err)
	}
	return matched
}

func readFunc(expression string) (fun string, err error) {
	fun = ""
	if isNextFunc(expression) {
		reg := regexp.MustCompile(funcRegExp)
		res := reg.FindAllString(expression, -1)[0]
		fun = res
	} else {
		err = errors.New("invalid Function")
	}
	return
}
