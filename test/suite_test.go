package test

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/onlyafly/yxne/internal/ast"
	"github.com/onlyafly/yxne/internal/interpreter"
	"github.com/onlyafly/yxne/internal/parser"
	"github.com/onlyafly/yxne/internal/util"
)

const (
	testsuiteDir   = "test/testdata"
	baseDir        = ".."
	fileExtPattern = "*.yaml"
)

// TestFullSuite runs the entire language test suite
func TestFullSuite(t *testing.T) {

	// We set the base directory so that the test cases can use paths that make
	// sense. If this is not set, the current working directory while the tests
	// run will be "test"
	err := os.Chdir(baseDir)
	if err != nil {
		t.Errorf("Error changing directory to <" + baseDir + ">: " + err.Error())
	}

	err = filepath.Walk(testsuiteDir, func(fp string, fi os.FileInfo, err error) error {
		if err != nil {
			return nil // Can't visit this node, but continue walking elsewhere
		}
		if fi.IsDir() {
			return nil // Not a file, ignore.
		}

		name := fi.Name()
		matched, err := filepath.Match(fileExtPattern, name)
		if err != nil {
			return err // malformed pattern
		}

		if matched {
			testInputFile(fp, t)
		}

		return nil
	})

	if err != nil {
		t.Errorf("Error walking test suite directory <" + testsuiteDir + ">: " + err.Error())
	}
}

func testInputFile(sourceFilePath string, t *testing.T) {
	sourceDirPart, sourceFileNamePart := filepath.Split(sourceFilePath)
	parts := strings.Split(sourceFileNamePart, ".")
	testName := parts[0]

	outputFilePath := sourceDirPart + testName + ".out"

	input, errIn := util.ReadFile(sourceFilePath)
	if errIn != nil {
		t.Errorf("Error reading file <" + sourceFilePath + ">: " + errIn.Error())
		return
	}

	expectedRaw, errOut := util.ReadFile(outputFilePath)
	if errOut != nil {
		t.Errorf("Error reading file <" + outputFilePath + ">: " + errOut.Error())
		return
	}

	// Remove any carriage return line endings from .out file
	expectedWithUntrimmed := strings.Replace(expectedRaw, "\r", "", -1)
	expected := strings.TrimSpace(expectedWithUntrimmed)

	nodes, errors := parser.Parse(input, sourceFilePath)
	if errors.Len() != 0 {
		verify(t, sourceFilePath, input, expected, errors.String())
	} else {
		e := interpreter.NewTopLevelMapEnv()

		var outputBuffer bytes.Buffer

		dummyReadLine := func() string {
			return "text from dummy read line"
		}

		var result ast.Node
		var evalError error
		for _, n := range nodes {
			result, evalError = interpreter.Eval(e, n, &outputBuffer, dummyReadLine)
			if evalError != nil {
				break
			}
		}

		actual := (&outputBuffer).String()

		if evalError == nil {
			//DEBUG fmt.Printf("RESULT(%v): %v\n", sourceFilePath, result)
			if result != nil {
				actual = actual + result.String()
			}
		} else {
			actual = actual + evalError.Error()
		}
		verify(t, sourceFilePath, input, expected, actual)
	}
}

func verify(t *testing.T, testCaseName, input, expected, actual string) {
	if expected != actual {
		t.Errorf(
			"\n===== TEST SUITE CASE FAILED: %s\n"+
				"===== INPUT\n%v\n"+
				"===== EXPECTED\n%v\n"+
				"===== ACTUAL\n%v\n"+
				"===== END\n",
			testCaseName,
			input,
			expected,
			actual)
	}
}
