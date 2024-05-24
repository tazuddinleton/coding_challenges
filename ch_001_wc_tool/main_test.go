package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestByteCount(t *testing.T) {
	files := []string{"emptyfile.txt", "testfile1.txt", "testfile2.txt"}
	contents := [][]byte{
		[]byte(""),
		[]byte("Hello world"),
		[]byte("Hello world"),
	}

	err := createTestFiles(files, contents)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	os.Stdout = w
	countBytes(files)
	w.Close()
	os.Stdout = old

	out, err := io.ReadAll(r)

	expected := fmt.Sprintf(
		"%d %s\n%d %s\n%d %s\n%d total\n",
		len(contents[0]),
		files[0],
		len(contents[1]),
		files[1],
		len(contents[2]),
		files[2],
		len(contents[0])+len(contents[1])+len(contents[2]),
	)
	if string(out) != expected {
		t.Errorf("expected: %s\n\nactual: %s", expected, string(out))
		t.Fail()
	}

	defer removeTestFiles(files)
}

func TestLineCount(t *testing.T) {
	files := []string{"emptyfile.txt", "testfile1.txt", "testfile2.txt"}
	contents := [][]byte{
		[]byte(""),
		[]byte("Hello\nWorld"),
		[]byte("Hello\nWorld\nAnother\nLine"),
	}

	err := createTestFiles(files, contents)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	os.Stdout = w
	countLines(files)
	w.Close()
	os.Stdout = old

	out, err := io.ReadAll(r)

	expected := fmt.Sprintf(
		"%d %s\n%d %s\n%d %s\n%d total\n",
		0,
		files[0],
		2,
		files[1],
		4,
		files[2],
		6,
	)
	if string(out) != expected {
		t.Errorf("expected: %s\n\nactual: %s", expected, string(out))
		t.Fail()
	}

	defer removeTestFiles(files)
}

func removeTestFiles(files []string) {
	for _, f := range files {
		os.Remove(f)
	}
}

func createTestFiles(files []string, contents [][]byte) error {
	if len(files) != len(contents) {
		return fmt.Errorf("Number of files and contents do not match")
	}

	for i, f := range files {
		err := ioutil.WriteFile(f, contents[i], 0644)
		if err != nil {
			fmt.Println("Can not create file")
			return err
		}
	}
	return nil
}
