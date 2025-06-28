package govaluate

import (
	"errors"
	"fmt"
	"strings"
)

/*
Struct used to test "parameter calls".
*/
type dummyParameter struct {
	String    string
	Int       int
	BoolFalse bool
	Nil       interface{}
	Nested    dummyNestedParameter
	Map       map[string]interface{}
}

func (dummyParameter) Func() string {
	return "funk"
}

func (dummyParameter) Func2() (string, error) {
	return "frink", nil
}

func (*dummyParameter) Func3() string {
	return "fronk"
}

func (dummyParameter) FuncArgStr(arg1 string) string {
	return arg1
}

func (dummyParameter) TestArgs(str string, ui uint, ui8 uint8, ui16 uint16, ui32 uint32, ui64 uint64, i int, i8 int8, i16 int16, i32 int32, i64 int64, f32 float32, f64 float64, b bool) string {

	var sum float64

	sum = float64(ui) + float64(ui8) + float64(ui16) + float64(ui32) + float64(ui64)
	sum += float64(i) + float64(i8) + float64(i16) + float64(i32) + float64(i64)
	sum += float64(f32)

	if b {
		sum += f64
	}

	return fmt.Sprintf("%v: %v", str, sum)
}

func (dummyParameter) AlwaysFail() (interface{}, error) {
	return nil, errors.New("function should always fail")
}

type dummyNestedParameter struct {
	Funk string
}

func (dummyNestedParameter) Dunk(arg1 string) string {
	return arg1 + "dunk"
}

var dummyParameterInstance = dummyParameter{
	String:    "string!",
	Int:       101,
	BoolFalse: false,
	Nil:       nil,
	Nested: dummyNestedParameter{
		Funk: "funkalicious",
	},
	Map: map[string]interface{}{
		"String":        "string!",
		"Int":           101,
		"StringCompare": strings.Compare,
		"IntArray":      []interface{}{1, 2, 3},
		"StringArray":   []interface{}{"foo", "bar", "baz"},
	},
}

var fooParameter = EvaluationParameter{
	Name:  "foo",
	Value: dummyParameterInstance,
}

var fooPtrParameter = EvaluationParameter{
	Name:  "fooptr",
	Value: &dummyParameterInstance,
}

var fooFailureParameters = map[string]interface{}{
	"foo":    fooParameter.Value,
	"fooptr": &fooPtrParameter.Value,
}
