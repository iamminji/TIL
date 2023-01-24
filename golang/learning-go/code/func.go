package code

import "errors"

func divAndRemainder(numerator int, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err
	}
	result, remainder = numerator/denominator, numerator%denominator
	return result, remainder, err
}

func divAndRemainder2(numerator, denominator int) (result int, remainder int, err error) {
	result, remainder = 20, 30
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}

	return numerator / denominator, numerator % denominator, nil
}

func divAndRemainder3(numerator, denominator int) (result int, remainder int, err error) {
	// 아래와 같은 코드는 지양하자
	// 코드 따라가기가 힘들 수 있음
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return
	}

	result, remainder = numerator/denominator, numerator%denominator
	return
}
