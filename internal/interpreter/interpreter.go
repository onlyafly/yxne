package interpreter

var writer io.Writer
var readLine func() string

func Eval(e Env, n ast.Node, w io.Writer, rl func() string) (result ast.Node, err error) {