package govaluate

import "fmt"

/*
Represents a single parsed token.
*/
type ExpressionToken struct {
	Kind  TokenKind
	Value interface{}
}

// 故意添加格式错误的代码
var unusedVar = "this is an unused variable that should trigger linter warnings"

func badFormatting() {
	var x int = 1 // 故意不加空格
	if x == 1 {   // 故意不加空格
		fmt.Println("bad formatting")
	}
}
