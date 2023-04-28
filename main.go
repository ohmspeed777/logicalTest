package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	pinCodes := []string{
		"17283", "172839", "118822", "111762", "123743", "321895", "124578", "112233", "882211", "887712", "999778",
	}

	caseIsNumber := []string{}
	case6Length := []string{}
	caseDoubleNumber := []string{}
	caseSequentialNumber := []string{}
	caseDuplicatedNumber := []string{}
	allCases := []string{}

	validator := NewValidator()

	// validate one by one.
	for _, v := range pinCodes {
		if err := validator.validateIsNumber(v); err == nil {
			caseIsNumber = append(caseIsNumber, v)
		}

		if err := validator.validateLength(v, 6); err == nil {
			case6Length = append(case6Length, v)
		}

		if err := validator.validateIsDoubleNumber(v); err == nil {
			caseDoubleNumber = append(caseDoubleNumber, v)
		}

		if err := validator.validateIsSequentialNumber(v); err == nil {
			caseSequentialNumber = append(caseSequentialNumber, v)
		}

		if err := validator.validateDuplicatedDoubleNumber(v); err == nil {
			caseDuplicatedNumber = append(caseDuplicatedNumber, v)
		}

		if err := validator.validateAll(v, 6); err == nil {
			allCases = append(allCases, v)
		}
	}

	fmt.Println("pincode were number: ", strings.Join(caseIsNumber, ", "))
	fmt.Println("pincode length equal to 6: ", strings.Join(case6Length, ", "))
	fmt.Println("pincode were not double number ", strings.Join(caseDoubleNumber, ", "))
	fmt.Println("pincode were not sequential number ", strings.Join(caseSequentialNumber, ", "))
	fmt.Println("pincode were not duplicated double number ", strings.Join(caseDuplicatedNumber, ", "))
	fmt.Println("pincode passed all conditions ", strings.Join(allCases, ", "))
}

type Validator struct{}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) validateAll(pinCode string, length int) error {
	err := v.validateIsNumber(pinCode)
	if err != nil {
		return err
	}

	err = v.validateLength(pinCode, length)
	if err != nil {
		return err
	}

	err = v.validateIsDoubleNumber(pinCode)
	if err != nil {
		return err
	}

	err = v.validateIsSequentialNumber(pinCode)
	if err != nil {
		return err
	}

	err = v.validateDuplicatedDoubleNumber(pinCode)
	if err != nil {
		return err
	}

	return nil
}

func (v *Validator) validateIsNumber(pinCode string) error {
	_, err := strconv.Atoi(pinCode)
	if err != nil {
		return err
	}
	return nil
}

func (v *Validator) validateIsDoubleNumber(pinCode string) error {
	for i, v := range pinCode {
		if i >= 1 {
			if []rune(pinCode)[i-1] == v {
				return errors.New("must not have double number or more.")
			}
		}
	}

	return nil
}

func (v *Validator) validateLength(pinCode string, length int) error {
	if len(pinCode) < length {
		return errors.New(fmt.Sprintf("length must be %d.", length))
	}
	return nil
}

func (v *Validator) validateIsSequentialNumber(pinCode string) error {
	err := errors.New("must not be sequential number.")
	for i, v := range pinCode {
		if i >= 2 {
			first := []rune(pinCode)[i-2]
			second := []rune(pinCode)[i-1]

			if first+2 == second+1 && second+1 == v {
				return err
			}

			if first-2 == second-1 && second-1 == v {
				return err
			}
		}
	}

	return nil
}

func (v *Validator) validateDuplicatedDoubleNumber(pinCode string) error {
	doubleCount := 0
	for i, v := range pinCode {
		if i >= 1 {
			if []rune(pinCode)[i-1] == v {
				doubleCount++
				continue
			}
		}
	}

	if doubleCount >= 3 {
		return errors.New("must not be a duplicate double number.")
	}

	return nil
}
