package heuristics

import (
	"errors"
	"log"
	"math"
	"strconv"
	"strings"
)

const ( // iota is reset to 0
	source = iota // source == 0
	target = iota // target == 1
)

type Params map[string]interface{}

type stage func(operands *[]float64, sourceParams Params, targetParams Params) float64

func yardToStages(yard []string) []stage {
	stages := make([]stage, 0)

	for _, value := range yard {
		switch value {
		case "+":
			stages = append(stages, plus)
		case "-":
			stages = append(stages, minus)
		case "*":
			stages = append(stages, multiply)
		case "/":
			stages = append(stages, divide)
		case "^":
			stages = append(stages, pow)
		case "u_-":
			stages = append(stages, unar_minus)
		case "abs":
			stages = append(stages, abs)
		case "max":
			stages = append(stages, max)
		case "min":
			stages = append(stages, min)
		default:
			if strings.HasPrefix(value, sourceMark) {
				param := value[len(sourceMark)+1:]
				stages = append(stages, getParameterStage(param, source))
			} else if strings.HasPrefix(value, targetMark) {
				param := value[len(targetMark)+1:]
				stages = append(stages, getParameterStage(param, target))
			} else {
				stages = append(stages, getOpearandStage(value))
			}
		}
	}

	return stages
}

func getOpearandStage(operand string) stage {
	num, err := strconv.ParseFloat(operand, 64)
	if err != nil {
		log.Panic(err)
	}

	return func(operands *[]float64, sourceParams, targetParams Params) float64 {
		return num
	}
}

func getParameterStage(parameter string, sourceOrTarget uint) stage {
	if sourceOrTarget == source {

		return func(operands *[]float64, sourceParams, targetParams Params) float64 {
			num, err := getFloat(sourceParams[parameter])
			if err != nil {
				log.Panic(err)
			}
			return num
		}

	} else if sourceOrTarget == target {

		return func(operands *[]float64, sourceParams, targetParams Params) float64 {
			num, err := getFloat(targetParams[parameter])
			if err != nil {
				log.Panic(err)
			}
			return num
		}

	} else {
		panic("Bad sourceOrTarget: " + strconv.FormatUint(uint64(sourceOrTarget), 10))
	}
}

//
// Operations
//

func plus(operands *[]float64, sourceParams, targetParams Params) float64 {
	operand1 := (*operands)[0]
	operand2 := (*operands)[1]
	*operands = (*operands)[2:]

	result := operand2 + operand1
	return result
}

func minus(operands *[]float64, sourceParams, targetParams Params) float64 {
	operand1 := (*operands)[0]
	operand2 := (*operands)[1]
	*operands = (*operands)[2:]

	result := operand2 - operand1
	return result
}

func multiply(operands *[]float64, sourceParams, targetParams Params) float64 {
	operand1 := (*operands)[0]
	operand2 := (*operands)[1]
	*operands = (*operands)[2:]

	result := operand2 * operand1
	return result
}

func divide(operands *[]float64, sourceParams, targetParams Params) float64 {
	operand1 := (*operands)[0]
	operand2 := (*operands)[1]
	*operands = (*operands)[2:]

	result := operand2 / operand1
	return result
}

func pow(operands *[]float64, sourceParams, targetParams Params) float64 {
	operand1 := (*operands)[0]
	operand2 := (*operands)[1]
	*operands = (*operands)[2:]

	result := math.Pow(operand2, operand1)
	return result
}

func unar_minus(operands *[]float64, sourceParams, targetParams Params) float64 {
	operand1 := (*operands)[0]
	*operands = (*operands)[1:]

	result := -operand1
	return result
}

func abs(operands *[]float64, sourceParams, targetParams Params) float64 {
	operand1 := (*operands)[0]
	*operands = (*operands)[1:]

	result := math.Abs(operand1)
	return result
}

func max(operands *[]float64, sourceParams, targetParams Params) float64 {
	operand1 := (*operands)[0]
	operand2 := (*operands)[1]
	*operands = (*operands)[2:]

	result := math.Max(operand2, operand1)
	return result
}

func min(operands *[]float64, sourceParams, targetParams Params) float64 {
	operand1 := (*operands)[0]
	operand2 := (*operands)[1]
	*operands = (*operands)[2:]

	result := math.Min(operand2, operand1)
	return result
}

//
// general
//

var errUnexpectedType = errors.New("Non-numeric type could not be converted to float")

func getFloat(unk interface{}) (float64, error) {
	switch i := unk.(type) {
	case float64:
		return i, nil
	case float32:
		return float64(i), nil
	case int64:
		return float64(i), nil
	case int32:
		return float64(i), nil
	case int:
		return float64(i), nil
	case uint64:
		return float64(i), nil
	case uint32:
		return float64(i), nil
	case uint:
		return float64(i), nil
	default:
		return math.NaN(), errUnexpectedType
	}
}
