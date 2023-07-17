package lambda_router

import (
	"encoding/base64"
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/jgroeneveld/trial/assert"
)

type customStruct struct {
	StructKey string `json:"structKey"`
}

func TestCustomRes(t *testing.T) {
	os.Setenv(CORSHeadersEnvKey, "headers-header-val")
	defer os.Unsetenv(CORSHeadersEnvKey)
	os.Setenv(CORSMethodsEnvKey, "methods-header-val")
	defer os.Unsetenv(CORSMethodsEnvKey)
	os.Setenv(CORSOriginEnvKey, "origin-header-val")
	defer os.Unsetenv(CORSOriginEnvKey)

	httpStatus := http.StatusTeapot
	headers := map[string]string{
		"key": "value",
	}

	structValue := customStruct{
		StructKey: "structValue",
	}

	res, err := CustomRes(httpStatus, headers, structValue)
	assert.Nil(t, err)

	t.Run("verify CustomRes returns the struct in the response body", func(t *testing.T) {
		var returnedStruct customStruct
		err = UnmarshalRes(res, &returnedStruct)
		assert.Nil(t, err)

		assert.Equal(t, structValue, returnedStruct)
	})
	t.Run("verify CustomRes returns the key value pair in the response headers", func(t *testing.T) {
		assert.Equal(t, res.Headers["key"], headers["key"])
	})
	t.Run("verify CustomRes returns the correct status code", func(t *testing.T) {
		assert.Equal(t, httpStatus, res.StatusCode)
	})
	t.Run("verify CustomRes returns CORS headers", func(t *testing.T) {
		assert.Equal(t, res.Headers[CORSHeadersHeaderKey], "headers-header-val")
		assert.Equal(t, res.Headers[CORSMethodsHeaderKey], "methods-header-val")
		assert.Equal(t, res.Headers[CORSOriginHeaderKey], "origin-header-val")
	})
}

func TestEmptyRes(t *testing.T) {
	res, err := EmptyRes()
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Nil(t, err)
	assert.Equal(t, "{}", res.Body)

	t.Run("verify EmptyRes returns the correct status code", func(t *testing.T) {
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})
	t.Run("verify EmptyRes returns CORS headers", func(t *testing.T) {
		assert.Equal(t, res.Headers[CORSHeadersEnvKey], "")
		assert.Equal(t, res.Headers[CORSMethodsEnvKey], "")
		assert.Equal(t, res.Headers[CORSOriginEnvKey], "")
	})
}

func TestErrorRes(t *testing.T) {
	os.Setenv(CORSHeadersEnvKey, "*")
	defer os.Unsetenv(CORSHeadersEnvKey)
	os.Setenv(CORSMethodsEnvKey, "*")
	defer os.Unsetenv(CORSMethodsEnvKey)
	os.Setenv(CORSOriginEnvKey, "*")
	defer os.Unsetenv(CORSOriginEnvKey)

	t.Run("Handle an HTTPError ErrorRes without ExposeServerErrors set and verify CORS", func(t *testing.T) {
		res, _ := ErrorRes(HTTPError{
			Status:  http.StatusBadRequest,
			Message: "Invalid input",
		})
		assert.Equal(t, http.StatusBadRequest, res.StatusCode, "status status must be correct")
		assert.Equal(t, `{"status":400,"message":"Invalid input"}`, res.Body, "body must be correct")
		t.Run("verify ErrorRes returns CORS headers", func(t *testing.T) {
			assert.Equal(t, res.Headers[CORSHeadersHeaderKey], "*")
			assert.Equal(t, res.Headers[CORSMethodsHeaderKey], "*")
			assert.Equal(t, res.Headers[CORSOriginHeaderKey], "*")
		})
	})
	t.Run("Handle an HTTPError for ErrorRes when ExposeServerErrors is true", func(t *testing.T) {
		ExposeServerErrors = true
		res, _ := ErrorRes(HTTPError{
			Status:  http.StatusInternalServerError,
			Message: "database down",
		})
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "status must be correct")
		assert.Equal(t, `{"status":500,"message":"database down"}`, res.Body, "body must be correct")
	})
	t.Run("Handle an HTTPError for ErrorRes when ExposeServerErrors is false", func(t *testing.T) {
		ExposeServerErrors = false
		res, _ := ErrorRes(HTTPError{
			Status:  http.StatusInternalServerError,
			Message: "database down",
		})
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "status must be correct")
		assert.Equal(t, `{"status":500,"message":"Internal Server Error"}`, res.Body, "body must be correct")
	})
	t.Run("Handle a general error for ErrorRes when ExposeServerErrors is true", func(t *testing.T) {
		ExposeServerErrors = true
		res, _ := ErrorRes(errors.New("database down"))
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "status must be correct")
		assert.Equal(t, `{"status":500,"message":"database down"}`, res.Body, "body must be correct")
	})
	t.Run("Handle a general error for ErrorRes when ExposeServerErrors is false", func(t *testing.T) {
		ExposeServerErrors = false
		res, _ := ErrorRes(errors.New("database down"))
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "status must be correct")
		assert.Equal(t, `{"status":500,"message":"Internal Server Error"}`, res.Body, "body must be correct")
	})
}

func TestFileRes(t *testing.T) {
	os.Setenv(CORSHeadersEnvKey, "*")
	defer os.Unsetenv(CORSHeadersEnvKey)
	os.Setenv(CORSMethodsEnvKey, "*")
	defer os.Unsetenv(CORSMethodsEnvKey)
	os.Setenv(CORSOriginEnvKey, "*")
	defer os.Unsetenv(CORSOriginEnvKey)

	csvContent := `
header1, header2
value1, value2
`
	res, err := FileRes("text/csv", map[string]string{"key": "value"}, []byte(csvContent))
	assert.Nil(t, err)

	t.Run("verify FileRes returns the correct status code", func(t *testing.T) {
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})
	t.Run("verify FileRes marks the response as NOT base64 encoded", func(t *testing.T) {
		assert.False(t, res.IsBase64Encoded)
	})
	t.Run("verify FileRes embeds the bytes correctly in the response object as a string", func(t *testing.T) {
		assert.Equal(t, csvContent, res.Body)
	})
	t.Run("verify FileRes preserves the original header values", func(t *testing.T) {
		assert.Equal(t, "value", res.Headers["key"])
	})
	t.Run("verify FileRes returns CORS headers", func(t *testing.T) {
		assert.Equal(t, res.Headers[CORSHeadersHeaderKey], "*")
		assert.Equal(t, res.Headers[CORSMethodsHeaderKey], "*")
		assert.Equal(t, res.Headers[CORSOriginHeaderKey], "*")
	})
}

func TestFileB64Res(t *testing.T) {
	os.Setenv(CORSHeadersEnvKey, "*")
	defer os.Unsetenv(CORSHeadersEnvKey)
	os.Setenv(CORSMethodsEnvKey, "*")
	defer os.Unsetenv(CORSMethodsEnvKey)
	os.Setenv(CORSOriginEnvKey, "*")
	defer os.Unsetenv(CORSOriginEnvKey)

	csvContent := `
header1, header2
value1, value2
`
	res, err := FileB64Res("text/csv", map[string]string{"key": "value"}, []byte(csvContent))
	assert.Nil(t, err)

	t.Run("verify FileB64Res returns the correct status code", func(t *testing.T) {
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})
	t.Run("verify FileB64Res marks the response as base64 encoded", func(t *testing.T) {
		assert.True(t, res.IsBase64Encoded)
	})
	t.Run("verify FileB64Res embeds the bytes correctly in the response object as a byte64 encoded string", func(t *testing.T) {
		b64CSVContent := res.Body

		decodedCSVContent, decodeErr := base64.StdEncoding.DecodeString(b64CSVContent)
		assert.Nil(t, decodeErr)
		assert.Equal(t, csvContent, string(decodedCSVContent))
	})
	t.Run("verify FileRes preserves the original header values", func(t *testing.T) {
		assert.Equal(t, "value", res.Headers["key"])
	})
	t.Run("verify FileB64Res returns CORS headers", func(t *testing.T) {
		assert.Equal(t, res.Headers[CORSHeadersHeaderKey], "*")
		assert.Equal(t, res.Headers[CORSMethodsHeaderKey], "*")
		assert.Equal(t, res.Headers[CORSOriginHeaderKey], "*")
	})
}

func TestStatusAndErrorRes(t *testing.T) {
	os.Setenv(CORSHeadersEnvKey, "*")
	defer os.Unsetenv(CORSHeadersEnvKey)
	os.Setenv(CORSMethodsEnvKey, "*")
	defer os.Unsetenv(CORSMethodsEnvKey)
	os.Setenv(CORSOriginEnvKey, "*")
	defer os.Unsetenv(CORSOriginEnvKey)

	newErr := errors.New("hello there")
	res, err := StatusAndErrorRes(http.StatusTeapot, newErr)
	assert.Nil(t, err)

	t.Run("verify StatusAndErrorRes returns the correct status code", func(t *testing.T) {
		assert.Equal(t, http.StatusTeapot, res.StatusCode)
	})
	t.Run("verify StatusAndErrorRes returns CORS headers", func(t *testing.T) {
		assert.Equal(t, res.Headers[CORSHeadersHeaderKey], "*")
		assert.Equal(t, res.Headers[CORSMethodsHeaderKey], "*")
		assert.Equal(t, res.Headers[CORSOriginHeaderKey], "*")
	})
}

func TestSuccessRes(t *testing.T) {
	os.Setenv(CORSHeadersEnvKey, "*")
	defer os.Unsetenv(CORSHeadersEnvKey)
	os.Setenv(CORSMethodsEnvKey, "*")
	defer os.Unsetenv(CORSMethodsEnvKey)
	os.Setenv(CORSOriginEnvKey, "*")
	defer os.Unsetenv(CORSOriginEnvKey)

	cs := customStruct{StructKey: "hello there"}
	res, err := SuccessRes(cs)
	assert.Nil(t, err)
	t.Run("verify SuccessRes returns the correct status code", func(t *testing.T) {
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})
	t.Run("verify SuccessRes returns the struct in the response body", func(t *testing.T) {
		var returnedStruct customStruct
		unmarshalErr := UnmarshalRes(res, &returnedStruct)
		assert.Nil(t, unmarshalErr)
		assert.Equal(t, cs, returnedStruct)
	})
	t.Run("verify SuccessRes returns CORS headers", func(t *testing.T) {
		assert.Equal(t, res.Headers[CORSHeadersHeaderKey], "*")
		assert.Equal(t, res.Headers[CORSMethodsHeaderKey], "*")
		assert.Equal(t, res.Headers[CORSOriginHeaderKey], "*")
	})
}
