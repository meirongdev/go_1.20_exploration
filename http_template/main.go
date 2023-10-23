package main

import (
	"bytes"
	"fmt"

	// "github.com/hundt/test/template"
	"html/template"
)

type TestCase struct {
	Template       string
	Vars           map[string]interface{}
	ExpectedOutput string
}

var testCases = []*TestCase{
	// Regular quoted string after template literal
	// ok
	{
		Template:       "<script>var tpl = `<div></div>`;var path = \"/{{.V}}\";</script>",
		Vars:           map[string]interface{}{"V": ";alert(1);"},
		ExpectedOutput: "<script>var tpl = `<div></div>`;var path = \"/;alert(1);\";</script>",
	},
	// End template literal
	// panic: {{.X}} appears in a JS template literal
	{
		Template:       "<script>var x = `{{.X}}`;</script>",
		Vars:           map[string]interface{}{"X": "`+alert(1);`"},
		ExpectedOutput: "<script>var x = `\\x60\\x2balert(1);\\x60`;</script>",
	},
	// Inject template literal into script
	// ok
	{
		Template:       "<script>var x = {{.X}};</script>",
		Vars:           map[string]interface{}{"X": "`${alert(1)}`"},
		ExpectedOutput: "<script>var x = \"`${alert(1)}`\";</script>",
	},
	// Inject ${} into template literal
	// panic: {{.V}} appears in a JS template literal
	{
		Template:       "<script>var v = `{{.V}}`;</script>",
		Vars:           map[string]interface{}{"V": "${alert(1)}"},
		ExpectedOutput: "<script>var v = `\\x24{alert(1)}`;</script>",
	},
	// // Inject code into interpolation that looks closed
	// panic: {{.V}} appears in a JS template literal
	{
		Template:       "<script>var v = `${function(x){return x+1}({{.V}})}`;</script>",
		Vars:           map[string]interface{}{"V": "alert(1)"},
		ExpectedOutput: "<script>var v = `${function(x){return x+1}(\"alert(1)\")}`;</script>",
	},
}

func main() {
	for i, tc := range testCases {
		t := template.Must(template.New("test").Parse(tc.Template))
		buf := new(bytes.Buffer)
		err := t.Execute(buf, tc.Vars)
		if err != nil {
			panic(err)
		}
		if buf.String() != tc.ExpectedOutput {
			fmt.Printf("Test case %d failed: expected:\n  %s\nbut got:\n  %s\n\n", i+1, tc.ExpectedOutput, buf.String())
		}
	}
}

// Go1.19.1
// go run main.go
// Test case 1 failed: expected:
//   <script>var tpl = `<div></div>`;var path = "/;alert(1);";</script>
// but got:
//   <script>var tpl = `<div></div>`;var path = "/";alert(1);"";</script>

// Test case 2 failed: expected:
//   <script>var x = `\x60\x2balert(1);\x60`;</script>
// but got:
//   <script>var x = `"`+alert(1);`"`;</script>

// Test case 4 failed: expected:
//   <script>var v = `\x24{alert(1)}`;</script>
// but got:
//   <script>var v = `"${alert(1)}"`;</script>

// Go1.20.3以及之后的版本
// go run main.go
// panic: html/template:test:1:19: {{.X}} appears in a JS template literal
// goroutine 1 [running]:
// main.main()
//         /home/daomin/projects/meirongdev/go_1.20_exploration/http_template/main.go:56 +0x2cb
// exit status 2

//  GODEBUG=jstmpllitinterp=1 go run main.go                                                                      [22:21:30]
// Test case 2 failed: expected:
//   <script>var x = `\x60\x2balert(1);\x60`;</script>
// but got:
//   <script>var x = `\u0060\u002balert(1);\u0060`;</script>

// Test case 4 failed: expected:
//   <script>var v = `\x24{alert(1)}`;</script>
// but got:
//   <script>var v = `${alert(1)}`;</script>

// Test case 5 failed: expected:
//   <script>var v = `${function(x){return x+1}("alert(1)")}`;</script>
// but got:
//   <script>var v = `${function(x){return x+1}(alert(1))}`;</script>
