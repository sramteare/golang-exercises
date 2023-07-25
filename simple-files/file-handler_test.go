package simplefiles

import "testing"

func TestSimpleFileOprations(t *testing.T) {
	txt, err := simpleOpenWriteRead()
	if err != nil {
		t.Errorf("Errored %v", err)
	}
	if txt != `Hello world!!` {
		t.Errorf("should return expected text from the file found\n%s", txt)
	}
}
