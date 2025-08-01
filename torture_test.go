package govaluate

/*
	Courtesy of abrander
	ref: https://gist.github.com/abrander/fa05ae9b181b48ffe7afb12c961b6e90
*/
import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

var (
	hello  = "hello"
	empty  struct{}
	empty2 *string

	values = []interface{}{
		-1,
		0,
		12,
		13,
		"",
		"hello",
		&hello,
		nil,
		"nil",
		empty,
		empty2,
		true,
		false,
		time.Now(),
		rune('r'),
		int64(34),
		time.Duration(0),
		"true",
		"false",
		"\ntrue\n",
		"\nfalse\n",
		"12",
		"nil",
		"arg1",
		"arg2",
		int(12),
		int32(12),
		int64(12),
		complex(1.0, 1.0),
		[]byte{0, 0, 0},
		[]int{0, 0, 0},
		[]string{},
		"[]",
		"{}",
		"\"\"",
		"\"12\"",
		"\"hello\"",
		".*",
		"==",
		"!=",
		">",
		">=",
		"<",
		"<=",
		"=~",
		"!~",
		"in",
		"&&",
		"||",
		"^",
		"&",
		"|",
		">>",
		"<<",
		"+",
		"-",
		"*",
		"/",
		"%",
		"**",
		"-",
		"!",
		"~",
		"?",
		":",
		"??",
		"+",
		"-",
		"*",
		"/",
		"%",
		"**",
		"&",
		"|",
		"^",
		">>",
		"<<",
		",",
		"(",
		")",
		"[",
		"]",
		"\n",
		"\000",
	}

	panics = 0
)

const (
	ITERATIONS = 10000000
	SEED       = 1487873697990155515
)

// Create a local random number generator with fixed seed
var localRand = rand.New(rand.NewSource(SEED))

func TestPanics(test *testing.T) {

	if os.Getenv("GOVALUATE_TORTURE_TEST") == "" {
		test.Logf("'GOVALUATE_TORTURE_TEST' env var not set - skipping torture test.")
		test.Skip()
		return
	}

	fmt.Printf("Running %d torture test cases...\n", ITERATIONS)

	for i := 0; i < ITERATIONS; i++ {

		num := localRand.Intn(3) + 2
		expression := ""

		for n := 0; n < num; n++ {
			expression += fmt.Sprintf(" %s", getRandom(values))
		}

		checkPanic(expression, test)
	}

	test.Logf("Done. %d/%d panics.\n", panics, ITERATIONS)
	if panics > 0 {
		test.Fail()
	}
}

func checkPanic(expression string, test *testing.T) {

	parameters := make(map[string]interface{})

	defer func() {
		if r := recover(); r != nil {
			test.Logf("Panic: \"%s\". Expression: \"%s\". Parameters: %+v\n", r, expression, parameters)
			panics++
		}
	}()

	eval, _ := NewEvaluableExpression(expression)
	if eval == nil {
		return
	}

	vars := eval.Vars()
	for _, v := range vars {
		parameters[v] = getRandom(values)
	}

	_, _ = eval.Evaluate(parameters)
}

func getRandom(haystack []interface{}) interface{} {

	i := localRand.Intn(len(haystack))
	return haystack[i]
}
