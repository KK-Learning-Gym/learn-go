package main

import (
	"fmt"
	"os"
	"testing"
)

// Helper Utilities

func createPage(title string, body []byte) {
	p1 := &Page{Title: title, Body: body}
	p1.save()
	fmt.Printf("HELPER: Created %s.txt\n", title)
}

func deleteFile(filename string) {
	fileLoc := dataLoc + filename + ".txt"

	_, existErr := os.Stat(fileLoc)
	if os.IsNotExist(existErr) {
		return
	}

	removeErr := os.Remove(fileLoc)
	if removeErr != nil {
		fmt.Println(removeErr)
	}
	fmt.Printf("HELPER: %s.txt deleted\n", filename)
}

// Tests

func Test_loadPage(t *testing.T) {
	// Test Config
	title := "test"
	body := []byte("This is a sample page.")
	// Test Setup
	deleteFile(title)
	createPage(title, body)
	// Test Run
	page, err := loadPage(title)
	if err != nil {
		fmt.Println("Page load failed")
		t.FailNow()
	}
	if page.Title != title {
		t.Errorf("Expected page.Title to be: \"%s\" but got: %s", title, page.Title)
	}
	if string(page.Body) != "This is a sample page." {
		t.Errorf("Expected page.Body to be: \"%s\" but got: \"%s\"", body, page.Body)
	}
}
