package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// Helper Utilities

func createPage(title string, body []byte) {
	p1 := &Page{Title: title, Body: body}
	p1.save()
	fmt.Printf("+++ HELPER: created %s.txt\n", title)
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
	fmt.Printf("+++ HELPER: %s.txt deleted\n", filename)
}

func runServer(fn func(http.ResponseWriter, *http.Request, string)) *httptest.Server {
	return httptest.NewServer(makeHandler(fn))
}

// Tests

func Test__PageLoad(t *testing.T) {
	// Test Config
	title := "test"
	body := []byte("This is a sample page.")
	// Test Setup
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

func Test__HandlerStatusWhenFileExists(t *testing.T) {
	// Test Config
	title := "test"
	body := []byte("This is a sample page.")
	// Test Setup
	createPage(title, body)
	handlers := map[string]func(http.ResponseWriter, *http.Request, string){"view": viewHandler, "edit": editHandler}
	// Test Run
	for handlerName, handler := range handlers {
		ts := runServer(handler)
		url := fmt.Sprintf("%s/%s/%s", ts.URL, handlerName, title)
		res, getErr := http.Get(url)
		if getErr != nil {
			fmt.Println(getErr)
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("Expected StatusCode: 200, Received StatusCode: %v", res.StatusCode)
		}

		ts.Close()
	}
}

func Test_SuccessfulRedirectOnRoot(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(frontHandler))

	res, _ := http.Get(ts.URL)
	if res.StatusCode != http.StatusFound {
		t.Errorf("Expected StatusCode: 302, Received StatusCode: %v", res.StatusCode)
	}

	ts.Close()
}
