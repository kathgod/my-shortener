package logic_test

import (
	"bytes"
	"compress/gzip"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	lgc "urlshortener/internal/logic"

	_ "github.com/lib/pq"
)

func TestRandSeq(t *testing.T) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for _, size := range []int{0, 1, 10, 100} {
		got := lgc.RandSeq(size)
		if len(got) != size {
			t.Errorf("RandSeq(%d) = %s; want length %d", size, got, size)
		}

		for _, c := range got {
			if !strings.ContainsRune(letters, c) {
				t.Errorf("RandSeq(%d) = %s; contains invalid character %c", size, got, c)
			}
		}
	}
}
func TestLogicGetFunc(t *testing.T) {
	handMapGet := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "DELETE",
	}

	tests := []struct {
		name     string
		urlPath  string
		expected int
	}{
		{"Test valid key", "/key1", http.StatusTemporaryRedirect},
		{"Test deleted key", "/key3", http.StatusGone},
		{"Test non-existent key", "/nonexistent", http.StatusBadRequest},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", test.urlPath, nil)
			if err != nil {
				t.Fatal(err)
			}

			status, _ := lgc.LogicGetFunc(req, handMapGet)

			if status != test.expected {
				t.Errorf("Expected status code %d, but got %d", test.expected, status)
			}
		})
	}
}

func TestLogicPostFunc(t *testing.T) {
	handMapPost := map[string]string{}
	handMapGet := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "DELETE",
	}

	tests := []struct {
		name            string
		body            string
		contentEncoding string
		expectedStatus  int
	}{
		{"Test regular body", "test data", "", http.StatusCreated},
		{"Test gzip body", "test gzip data", "gzip", http.StatusCreated},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/", strings.NewReader(test.body))
			if err != nil {
				t.Fatal(err)
			}

			if test.contentEncoding == "gzip" {
				req.Header.Set("Content-Encoding", "gzip")
			}

			rr := httptest.NewRecorder()

			status, _ := lgc.LogicPostFunc(rr, req, handMapPost, handMapGet)

			if status != test.expectedStatus {
				t.Errorf("Expected status code %d, but got %d", test.expectedStatus, status)
			}
		})
	}
}

func TestShortPostFunc(t *testing.T) {
	handMapPost := map[string]string{}
	handMapGet := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "DELETE",
	}

	tests := []struct {
		name       string
		bp         []byte
		cckValue   string
		wantError  int64
		wantPrefix string
	}{
		{"Test valid input", []byte("test data"), "123456", -1, lgc.ResHandParam.BaseURL},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resultPost, sqlError := lgc.ShortPostFunc(handMapPost, handMapGet, test.bp, test.cckValue)

			if sqlError != test.wantError {
				t.Errorf("Expected error code %d, but got %d", test.wantError, sqlError)
			}

			if !strings.HasPrefix(resultPost, test.wantPrefix) {
				t.Errorf("Expected resultPost to start with '%s', but got '%s'", test.wantPrefix, resultPost)
			}
		})
	}
}

func TestLogicPostFuncAPIShorten(t *testing.T) {
	handMapPost := map[string]string{}
	handMapGet := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "DELETE",
	}

	tests := []struct {
		name           string
		rawBsp         []byte
		expectedStatus int
	}{
		{"Test valid input", []byte(`{"url": "https://example.com"}`), http.StatusCreated},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := httptest.NewRequest("POST", "http://localhost:8080/api/shorten", bytes.NewBuffer(test.rawBsp))
			w := httptest.NewRecorder()

			status, _ := lgc.LogicPostFuncAPIShorten(handMapPost, handMapGet, w, r)

			if status != test.expectedStatus {
				t.Errorf("Expected status code %d, but got %d", test.expectedStatus, status)
			}
		})
	}
}

func TestShortPostFuncAPIShorten(t *testing.T) {
	handMapPost := map[string]string{}
	handMapGet := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "DELETE",
	}

	tests := []struct {
		name           string
		rawBsp         []byte
		expectedStatus int64
	}{
		{"Test valid input", []byte(`{"url": "https://example.com"}`), -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, status := lgc.ShortPostFuncAPIShorten(handMapPost, handMapGet, test.rawBsp)

			if status != test.expectedStatus {
				t.Errorf("Expected status code %d, but got %d", test.expectedStatus, status)
			}
		})
	}
}

func TestRecovery(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatal("Failed to create temporary file")
	}
	defer os.Remove(tmpFile.Name())

	testData := "https://example1.com@short1\nhttps://example2.com@short2\n"
	if _, err := tmpFile.Write([]byte(testData)); err != nil {
		t.Fatal("Failed to write test data to temporary file")
	}

	handMapPost := make(map[string]string)
	handMapGet := make(map[string]string)
	lgc.Recovery(handMapPost, handMapGet, tmpFile)

	expectedHandMapPost := map[string]string{
		"https://example1.com": "short1",
		"https://example2.com": "short2",
	}
	expectedHandMapGet := map[string]string{
		"short1": "https://example1.com",
		"short2": "https://example2.com",
	}

	if !reflect.DeepEqual(handMapPost, expectedHandMapPost) {
		t.Errorf("Expected handMapPost to be %v, but got %v", expectedHandMapPost, handMapPost)
	}

	if !reflect.DeepEqual(handMapGet, expectedHandMapGet) {
		t.Errorf("Expected handMapGet to be %v, but got %v", expectedHandMapGet, handMapGet)
	}
}

func TestHandParam(t *testing.T) {
	testCases := []struct {
		name     string
		flg      string
		envValue string
		expected string
	}{
		{"SERVER_ADDRESS", "localhost", "", "localhost"},
		{"SERVER_ADDRESS", "localhost", "otherhost", "otherhost"},
		{"BASE_URL", "https://example.com", "", "https://example.com/"},
		{"BASE_URL", "https://example.com", "https://other.com", "https://other.com/"},
		{"FILE_STORAGE_PATH", "/tmp/storage", "", "/tmp/storage"},
		{"FILE_STORAGE_PATH", "/tmp/storage", "/tmp/other", "/tmp/other"},
		{"DATABASE_DSN", "user:password@/dbname", "", "user:password@/dbname"},
		{"DATABASE_DSN", "user:password@/dbname", "otheruser:otherpassword@/otherdb", "otheruser:otherpassword@/otherdb"},
		{"ENABLE_HTTPS", "false", "", "false"},
		{"ENABLE_HTTPS", "false", "true", "true"},
		{"CONFIG", "<path>", "", "<path>"},
		{"CONFIG", "<path1>", "<path2>", "<path2>"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.envValue != "" {
				os.Setenv(tc.name, tc.envValue)
				defer os.Unsetenv(tc.name)
			}

			flg := tc.flg
			result := lgc.HandParam(tc.name, &flg)
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}

func TestHandConfigParam(t *testing.T) {
	res := "config.json"
	content := []byte(`{
        "PortNumber": "localhost:8000",
        "BaseURL": "http://localhost:8000",
        "FileStoragePath": "/var/www/uploads",
        "DataBaseDSN": "user:password@tcp(127.0.0.1:3306)/dbname",
        "EnableHTTPS": true
    }`)
	expected := lgc.FlagParam{
		PortNumber:      "localhost:8000",
		BaseURL:         "http://localhost:8000",
		FileStoragePath: "/var/www/uploads",
		DataBaseDSN:     "user:password@tcp(127.0.0.1:3306)/dbname",
		EnableHTTPS:     true,
	}

	lgc.ResHandParam = lgc.FlagParam{
		PortNumber:      "localhost:8000",
		BaseURL:         "http://localhost:8000",
		FileStoragePath: "/var/www/uploads",
		DataBaseDSN:     "user:password@tcp(127.0.0.1:3306)/dbname",
		EnableHTTPS:     true,
	}

	file, err := os.Create(res)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		log.Println(err)
	}

	lgc.HandConfigParam(res)

	if !reflect.DeepEqual(lgc.ResHandParam, expected) {
		t.Errorf("Result: %+v, expected: %+v", lgc.ResHandParam, expected)
	}
}

func TestDecompress(t *testing.T) {
	testCases := []struct {
		name        string
		data        []byte
		err0        error
		expected    []byte
		expectedErr error
	}{
		{
			name:        "Error case",
			data:        nil,
			err0:        errors.New("Test error"),
			expected:    nil,
			expectedErr: errors.New("Test error"),
		},
		{
			name:        "Uncompressed data",
			data:        []byte("Hello, world!"),
			err0:        nil,
			expected:    []byte("Hello, world!"),
			expectedErr: nil,
		},
		{
			name: "Compressed data",
			data: func() []byte {
				var buf bytes.Buffer
				gz := gzip.NewWriter(&buf)
				if _, err := gz.Write([]byte("Hello, world!")); err != nil {
					t.Fatal(err)
				}
				if err := gz.Close(); err != nil {
					t.Fatal(err)
				}
				return buf.Bytes()
			}(),
			err0:        nil,
			expected:    []byte("Hello, world!"),
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, _ := lgc.Decompress(tc.data, tc.err0)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected data %v, but got %v", tc.expected, result)
			}
		})
	}
}

func TestCompress(t *testing.T) {
	testCases := []struct {
		name        string
		data        []byte
		expectedErr error
	}{
		{
			name:        "Valid data",
			data:        []byte("Hello, world!"),
			expectedErr: nil,
		},
		{
			name:        "Empty data",
			data:        []byte(""),
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			compressedData, err := lgc.Compress(tc.data)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error %v, but got %v", tc.expectedErr, err)
			}

			if err == nil {
				decompressedData, err := lgc.Decompress(compressedData, nil)
				if err != nil {
					t.Errorf("Failed to decompress compressed data: %v", err)
				} else if !bytes.Equal(decompressedData, tc.data) {
					t.Errorf("Decompressed data does not match the original data: got %v, expected %v", decompressedData, tc.data)
				}
			}
		})
	}
}

func TestCookieCheck(t *testing.T) {
	mockWriter := httptest.NewRecorder()

	testCases := []struct {
		name         string
		cookieName   string
		cookieValue  string
		expectNewCCh bool
	}{
		{
			name:         "No cookie",
			cookieName:   "",
			cookieValue:  "",
			expectNewCCh: true,
		},
		{
			name:         "Invalid cookie",
			cookieName:   "userId",
			cookieValue:  "InvalidValue",
			expectNewCCh: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatal(err)
			}

			if tc.cookieName != "" {
				req.AddCookie(&http.Cookie{Name: tc.cookieName, Value: tc.cookieValue})
			}

			result := lgc.CookieCheck(mockWriter, req)

			if tc.expectNewCCh && (result == tc.cookieValue) {
				t.Errorf("Expected a new cookie, but got the same cookie value: %v", result)
			}
		})
	}
}

func TestMakeNewCookie(t *testing.T) {
	w := httptest.NewRecorder()

	cookieVal := lgc.MakeNewCookie(w)

	if cookieVal == "" {
		t.Errorf("MakeNewCookie() = empty cookie value, want non-empty value")
	}

	cookieHeader := w.Header().Get("Set-Cookie")
	if cookieHeader == "" {
		t.Errorf("MakeNewCookie() did not set cookie header")
	}

	decodedCookieVal, err := hex.DecodeString(cookieVal)
	if err != nil {
		t.Errorf("MakeNewCookie() could not decode cookie value: %v", err)
	}

	if len(decodedCookieVal) != 32 {
		t.Errorf("MakeNewCookie() = len(cookie value) = %d, want 32", len(decodedCookieVal))
	}
}

func TestLogicGetFuncAPIUserUrls(t *testing.T) {
	handMapGet := map[string]string{
		"user1abcdef": "http://example.com/1",
		"user1bbcdef": "http://example.com/2",
		"user2mnopqr": "http://example.com/3",
	}

	mockWriter := httptest.NewRecorder()

	mockRequest := httptest.NewRequest("GET", "/api/user/urls", nil)
	mockRequest.AddCookie(&http.Cookie{Name: "userId", Value: "bcdef"})

	status, data := lgc.LogicGetFuncAPIUserUrls(handMapGet, mockWriter, mockRequest)

	if status != http.StatusOK {
		t.Errorf("Expected status to be http.StatusOK, but got %v", status)
	}

	baseurl := "http://localhost:8080/"

	expectedURLs1 := []lgc.OrShURL{
		{ShortURL: baseurl + "user1abcdef", OriginalURL: "http://example.com/1"},
		{ShortURL: baseurl + "user1bbcdef", OriginalURL: "http://example.com/2"},
	}

	expectedURLs2 := []lgc.OrShURL{
		{ShortURL: baseurl + "user1bbcdef", OriginalURL: "http://example.com/2"},
		{ShortURL: baseurl + "user1abcdef", OriginalURL: "http://example.com/1"},
	}

	var responseURLs []lgc.OrShURL
	err := json.Unmarshal(data, &responseURLs)
	if err != nil {
		t.Errorf("Failed to unmarshal response data: %v", err)
	}

	if (!reflect.DeepEqual(responseURLs, expectedURLs1)) && (!reflect.DeepEqual(responseURLs, expectedURLs2)) {
		t.Errorf("Expected response URLs to be %v, but got %v", expectedURLs1, responseURLs)
		t.Errorf("Expected response URLs to be %v, but got %v", expectedURLs2, responseURLs)
	}
}

func TestLogicGetFuncPing(t *testing.T) {
	tests := []struct {
		name   string
		DBDSN  string
		result int
	}{
		{
			name:   "Negative test",
			DBDSN:  "user=postgres password=password host=invalidhost dbname=test port=5432 sslmode=disable",
			result: http.StatusInternalServerError,
		},
	}
	for _, tc := range tests {
		res := lgc.LogicGetFuncPing(tc.DBDSN)
		if res != tc.result {
			t.Errorf("Expected %v, but got %v", tc.result, res)
		}
	}

}

func TestCreateSQLTable(t *testing.T) {
	tests := []struct {
		name   string
		DBDSN  string
		result *sql.DB
	}{
		{
			name:   "Negative test",
			DBDSN:  "user=postgres password=password host=invalidhost dbname=test port=5432 sslmode=disable",
			result: nil,
		},
	}
	for _, tc := range tests {
		res := lgc.CreateSQLTable(tc.DBDSN)
		if res != tc.result {
			t.Errorf("Expected nil, but got another")
		}
	}
}

func TestAddRecordInTable(t *testing.T) {
	test := []struct {
		name     string
		DBDSN    string
		shortURL string
		longURL  string
		userID   string
		exresult int
	}{
		{
			name:     "Negative Test",
			DBDSN:    "user=postgres password=123 dbname=test host=localhost sslmode=disable",
			shortURL: "abc",
			longURL:  "http://example.com",
			userID:   "123",
			exresult: 0,
		},
	}
	for _, tc := range test {
		result := lgc.AddRecordInTable(tc.DBDSN, tc.shortURL, tc.longURL, tc.userID)
		if result != int64(tc.exresult) {
			t.Errorf("Expected rows affected: %d, but got: %d", tc.exresult, result)
		}
	}
}

func TestLogicPostFuncAPIShortenBatch(t *testing.T) {
	handMapPost := make(map[string]string)
	handMapGet := make(map[string]string)

	testCases := []struct {
		name           string
		inputBody      string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid request",
			inputBody:      `[{"original_url": "https://example.com"}]`,
			expectedStatus: http.StatusCreated,
			expectedBody:   `{"correlation_id":`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/api/shorten/batch", bytes.NewBufferString(tc.inputBody))
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()

			status, body := lgc.LogicPostFuncAPIShortenBatch(handMapPost, handMapGet, rr, req)

			if status != tc.expectedStatus {
				t.Errorf("Expected status %d, got %d", tc.expectedStatus, status)
			}
			if !strings.Contains(string(body), tc.expectedBody) {
				t.Errorf("Expected body to contain %s, got %s", tc.expectedBody, string(body))
			}
		})
	}
}

func TestShortPostAPIShortenBatch(t *testing.T) {
	handMapPost := make(map[string]string)
	handMapGet := make(map[string]string)

	testCases := []struct {
		name         string
		inputJSON    string
		expectedJSON string
	}{
		{
			name:         "Single URL",
			inputJSON:    `[{"correlation_id": "1", "original_url": "https://example.com"}]`,
			expectedJSON: `[{"correlation_id":"1","original_url":"","short_url":"http://localhost:8080/abc123"}]`,
		},
		{
			name:         "Multiple URLs",
			inputJSON:    `[{"correlation_id": "1", "original_url": "https://example.com"}, {"correlation_id": "2", "original_url": "https://example2.com"}]`,
			expectedJSON: `[{"correlation_id":"1","original_url":"","short_url":"http://localhost:8080/abc123"},{"correlation_id":"2","original_url":"","short_url":"http://localhost:8080/def456"}]`,
		},
	}

	rand.New(rand.NewSource(1))

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			inputBytes := []byte(tc.inputJSON)
			expectedBytes := []byte(tc.expectedJSON)

			outputBytes := lgc.ShortPostAPIShortenBatch(handMapPost, handMapGet, inputBytes)

			var expected, output []lgc.LngShrtCrltnID
			if err := json.Unmarshal(expectedBytes, &expected); err != nil {
				t.Fatal(err)
			}
			if err := json.Unmarshal(outputBytes, &output); err != nil {
				t.Fatal(err)
			}

			if len(expected) != len(output) {
				t.Errorf("Expected %d elements, got %d", len(expected), len(output))
			}

			for i := range expected {
				if expected[i].CorrelationID != output[i].CorrelationID {
					t.Errorf("Expected correlation_id %s, got %s", expected[i].CorrelationID, output[i].CorrelationID)
				}
				if len(expected[i].ShortURL) != len(output[i].ShortURL) {
					t.Errorf("Expected short_url %s, got %s", expected[i].ShortURL, output[i].ShortURL)
				}
			}
		})
	}
}

func TestLogicDeleteFuncAPIUserURLs(t *testing.T) {
	handMapPost := map[string]string{
		"https://example.com": "abc123",
	}
	handMapGet := map[string]string{
		"abc123": "https://example.com",
	}

	testCases := []struct {
		name           string
		inputBody      string
		expectedStatus int
	}{
		{
			name:           "Valid request",
			inputBody:      `["abc123"]`,
			expectedStatus: http.StatusAccepted,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("DELETE", "/api/user/urls", bytes.NewBufferString(tc.inputBody))
			if err != nil {
				t.Fatal(err)
			}

			status := lgc.LogicDeleteFuncAPIUserURLs(handMapPost, handMapGet, "", req)

			if status != tc.expectedStatus {
				t.Errorf("Expected status %d, got %d", tc.expectedStatus, status)
			}

			if handMapGet["abc123"] != "DELETE" {
				t.Errorf("Expected handMapGet[\"abc123\"] to be \"DELETE\", got %s", handMapGet["abc123"])
			}
			if _, ok := handMapPost["https://example.com"]; ok {
				t.Error("Expected handMapPost[\"https://example.com\"] to be deleted")
			}
		})
	}
}

func BenchmarkRandSeq(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lgc.RandSeq(6)
	}
}

func BenchmarkPostFunc(b *testing.B) {
	handMapPost := map[string]string{}
	handMapGet := map[string]string{}
	for i := 0; i < b.N; i++ {
		originalURL := lgc.RandSeq(10)
		req, err := http.NewRequest("POST", "http://localhost:8080/", strings.NewReader(originalURL))
		if err != nil {
			log.Println(err)
		}
		nr := httptest.NewRecorder()

		status, _ := lgc.LogicPostFunc(nr, req, handMapPost, handMapGet)
		log.Println(status, handMapGet[originalURL])

	}
}

func BenchmarkGetFunc(b *testing.B) {
	handMapGet := map[string]string{}
	for i := 0; i < b.N; i++ {
		shortURL := lgc.RandSeq(6)
		originalURL := lgc.RandSeq(10)
		handMapGet[shortURL] = originalURL
		reqURL := "http://localhost:8080/" + shortURL
		req1, err := http.NewRequest("GET", reqURL, nil)
		if err != nil {
			log.Println(err)
		}
		status1, _ := lgc.LogicGetFunc(req1, handMapGet)
		fmt.Println(status1)
	}
}
