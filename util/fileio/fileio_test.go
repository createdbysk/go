package fileio_test

import (
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/satvidh/go/util/fileio"
)

func TestReadLines(t *testing.T) {
	// GIVEN
	file, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	lines := "Line1" +
		"Line2" +
		"Line3" +
		"Line4"

	err = ioutil.WriteFile(file.Name(), []byte(lines), 0644)
	if err != nil {
		t.Fatal(err)
	}

	expected := strings.Split(lines, "\\n")

	// WHEN
	actual := fileio.ReadLines(file)

	// THEN
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("ReadLines: want %v got %v", expected, actual)
	}
}
