package helper

import (
	"log"
	"reflect"
	"strconv"
)

type Integer interface {
	int | int8 | int16 | int32 | int64
}

type UnsignedInteger interface {
	uint | uint8 | uint16 | uint32 | uint64
}

func Error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ErrorArg[T any](value T, err error) T {
	Error(err)
	return value
}

func ParseInt[T Integer](output *T, value string) error {
	result, err := strconv.ParseInt(value, 10, reflect.TypeOf(*output).Bits())
	*output = T(result)
	return err
}

func ParseUint[T UnsignedInteger](output *T, value string) error {
	result, err := strconv.ParseUint(value, 10, reflect.TypeOf(*output).Bits())
	*output = T(result)
	return err
}

func Min[T Integer | UnsignedInteger](a T, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T Integer | UnsignedInteger](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

func Clamp[T Integer | UnsignedInteger](value T, min T, max T) T {
	return Min(Max(value, min), max)
}
