package repository

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestListFiles(t *testing.T) {
	testDir := createTestDirectory()
	defer removeTestDirectory(testDir)

	createTestFile(testDir, "file1.txt")
	createTestFile(testDir, "file2.txt")
	createTestFile(testDir, "file1.pdf")
	createTestFile(testDir, "file2.pdf")

	testCases := []struct {
		extensions []string
		expected   []string
	}{
		{
			extensions: []string{"txt"},
			expected: []string{
				filepath.Join(testDir, "file1.txt"),
				filepath.Join(testDir, "file2.txt"),
			},
		},
		{
			extensions: []string{"txt"},
			expected: []string{
				filepath.Join(testDir, "file1.txt"),
				filepath.Join(testDir, "file2.txt"),
			},
		},
		{
			extensions: []string{"zip"},
			expected:   []string{},
		},
	}

	repo := NewDirectoryRepository()
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Extensions: %v", testCase.extensions), func(t *testing.T) {
			result, err := repo.ListFiles(testDir, testCase.extensions)

			if err != nil {
				t.Errorf("Error: %v", err)
			}

			assertStringSlicesEqual(t, result.Files, testCase.expected)
		})
	}
}

func assertStringSlicesEqual(t *testing.T, result, expected []string) {
	if len(result) != len(expected) {
		t.Errorf("Result and expected slices have different lengths. Result: %v, Expected: %v", result, expected)
		return
	}
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Mismatch at index %d. Result: %s, Expected: %s", i, result[i], expected[i])
		}
	}
}

func createTestFile(dir, filename string) {
	filePath := dir + "/" + filename
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Error creating test file: %v", err)
	}
	defer file.Close()
}

func createTestDirectory() string {
	dir, err := ioutil.TempDir("", "testdir")
	if err != nil {
		log.Fatalf("Error creating test directory: %v", err)
	}
	return dir
}

func removeTestDirectory(dir string) {
	err := os.RemoveAll(dir)
	if err != nil {
		log.Fatalf("Error removing test directory: %v", err)
	}
}
