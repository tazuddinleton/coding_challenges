package main

import (
	"fmt"
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
