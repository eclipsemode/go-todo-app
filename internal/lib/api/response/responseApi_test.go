package responseApi

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
	"testing"
)

type testValidationErrReq struct {
	Title string `json:"title" validate:"required,max=4,len=3"`
}

func TestResponseApi_Ok(t *testing.T) {
	res := Ok()

	if res.Status != StatusOK {
		t.Errorf("expected status %s, got %s", StatusOK, res.Status)
	}
}

func TestResponseApi_Error(t *testing.T) {
	errText := "something bad happened"

	res := Error(errText)

	if res.Status != StatusError {
		t.Errorf("expected status %s, got %s", StatusError, res.Status)
	}

	if res.Error != errText {
		t.Errorf("expected error %s, got %s", errText, res.Error)
	}
}

func TestResponseApi_ValidationError(t *testing.T) {
	cases := []struct {
		name string
		req  testValidationErrReq
	}{
		{
			name: "error_required",
			req:  testValidationErrReq{},
		},
		{
			name: "error_max_length",
			req: testValidationErrReq{
				Title: "Lorem",
			},
		},
		{
			name: "error_length",
			req: testValidationErrReq{
				Title: "Lo",
			},
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			err := validator.New().Struct(tCase.req)
			var validateErr validator.ValidationErrors
			errors.As(err, &validateErr)

			ve := ValidationError(validateErr)

			require.Equal(t, ve.Status, StatusError)

			reqErr := validateErr[0]

			require.Equal(t, reqErr.Field(), "Title")

			switch tCase.name {
			case "error_required":
				require.Equal(t, reqErr.ActualTag(), "required")
			case "error_max_length":
				require.Equal(t, reqErr.ActualTag(), "max")
			default:
				require.Equal(t, reqErr.ActualTag(), reqErr.ActualTag())

			}

		})
	}
}
