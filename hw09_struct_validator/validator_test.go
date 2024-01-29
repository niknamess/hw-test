package hw09structvalidator

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.

type (
	WrongCase struct {
		V0       int    `validate:"min:5|max:10"`
		NotImpl  int    `validate:"notimpl:notimpl"`
		WrongCmd string `validate:"len11"`
	}

	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
	}

	App struct {
		Version string `validate:"len:5"`
		User    User
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name          string
		in            interface{}
		expectedErrIs error
		expectedErrAs interface{}
	}{
		{
			name: "min",
			in: User{
				Age: 15,
			},
			expectedErrIs: ErrValidationMin,
		},
		{
			name: "max",
			in: User{
				Age: 62,
			},
			expectedErrIs: ErrValidationMax,
		},
		{
			name: "len",
			in: User{
				ID: "UUID",
			},
			expectedErrIs: ErrValidationLen,
		},
		{
			name: "in",
			in: User{
				Role: "admins",
			},
			expectedErrIs: ErrValidationIn,
		},
		{
			name: "regexp",
			in: User{
				Email: "test1@gmailcom",
			},
			expectedErrIs: ErrValidationRegExp,
		},
		{
			name: "not implemented",
			in: WrongCase{
				NotImpl: 1,
			},
			expectedErrIs: ErrValidationNImpl,
		},
		{
			name: "parse failure",
			in: WrongCase{
				WrongCmd: "wrong",
			},
			expectedErrIs: ErrValidationParse,
		},

		{
			name: "zero value",
			in: WrongCase{
				V0: 0,
			},
			expectedErrIs: ErrValidationMin,
		},

		{
			name: "all positive",
			in: User{
				Age:    40,
				ID:     "LEMZ",
				Role:   "admin",
				Email:  "test1@gmail.com",
				Phones: []string{"79999999999", "79999999999"},
				meta:   json.RawMessage(`{"name": value}`),
			},
			expectedErrIs: nil,
			expectedErrAs: nil,
		},
		{
			name:          "struct no tags",
			in:            UserRole("userRole"),
			expectedErrIs: nil,
			expectedErrAs: nil,
		},
		{
			name: "substructs",
			in: App{
				Version: "0.0.1",
				User: User{
					Age:   17,
					ID:    "LEMZ",
					Role:  "admin",
					Email: "test2@gmail.com",
				},
			},
		},
		{
			name: "substructs negative",
			in: App{
				Version: "0.0.1",
				User: User{
					Age: 1,
				},
			},
			expectedErrIs: ErrValidationMin,
		},
		{
			name: "errors.As",
			in: Response{
				Code: 300,
				Body: "{\"name\": \"value\"}",
			},
			expectedErrAs: &ValidationError{Field: "Code", Err: ErrValidationIn},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("case %s", tt.name), func(t *testing.T) {
			tt := tt
			t.Parallel()
			err := Validate(tt.in)
			if tt.expectedErrAs == nil && tt.expectedErrIs == nil {
				require.Truef(t, errors.Is(err, nil), "expected:\"%v\" got:\"%v\"", nil, err)
			}
			if tt.expectedErrIs != nil {
				require.Truef(t, errors.Is(err, tt.expectedErrIs), "expected:\"%v\" got:\"%v\"", tt.expectedErrIs, err)
			}
			if tt.expectedErrAs != nil {
				require.Truef(t, errors.As(err, tt.expectedErrAs), "expected:\"%v\" got:\"%v\"", tt.expectedErrAs, err)
			}
			_ = tt
		})
	}
}
