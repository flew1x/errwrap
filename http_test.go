package errwrap_test

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/flew1x/errwrap"
	"github.com/stretchr/testify/require"
)

func TestWriteHTTPError_AppError(t *testing.T) {
	tests := []struct {
		name         string
		err          error
		expectedCode int
		expectedBody errwrap.ErrorResponse
	}{
		{
			name: "not found",
			err: errwrap.Wrap("GetUser", errwrap.CodeNotFound, errors.New("user not found"), map[string]any{
				"id": "123",
			}),
			expectedCode: http.StatusNotFound,
			expectedBody: errwrap.ErrorResponse{
				Code:      errwrap.CodeNotFound,
				Message:   "user not found",
				Operation: "GetUser",
				Meta: map[string]any{
					"id": "123",
				},
			},
		},
		{
			name:         "unauthenticated",
			err:          errwrap.Wrap("AuthCheck", errwrap.CodeUnauthenticated, errors.New("token missing"), nil),
			expectedCode: http.StatusUnauthorized,
			expectedBody: errwrap.ErrorResponse{
				Code:      errwrap.CodeUnauthenticated,
				Message:   "token missing",
				Operation: "AuthCheck",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			errwrap.WriteHTTPError(rec, tt.err)

			res := rec.Result()
			defer res.Body.Close()

			require.Equal(t, tt.expectedCode, res.StatusCode)

			body, err := io.ReadAll(res.Body)
			require.NoError(t, err)

			var resp errwrap.ErrorResponse
			err = json.Unmarshal(body, &resp)
			require.NoError(t, err)

			require.Equal(t, tt.expectedBody.Code, resp.Code)
			require.Equal(t, tt.expectedBody.Message, resp.Message)
			require.Equal(t, tt.expectedBody.Operation, resp.Operation)
			require.Equal(t, tt.expectedBody.Meta, resp.Meta)
		})
	}
}

func TestWriteHTTPError_GenericError(t *testing.T) {
	rec := httptest.NewRecorder()

	err := errors.New("unexpected error")
	errwrap.WriteHTTPError(rec, err)

	res := rec.Result()
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	require.Equal(t, http.StatusInternalServerError, res.StatusCode)
	require.Equal(t, "Internal Server Error\n", string(body)) // стандартный ответ от http.Error
}
