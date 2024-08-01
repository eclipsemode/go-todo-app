package todos

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateValidate(t *testing.T) {
	cases := []struct {
		name string
		req  createTodoReq
	}{
		{
			name: "default",
			req: createTodoReq{
				Title:       "Lorem ipsum",
				Description: "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam",
			},
		},
		{
			name: "no_description",
			req: createTodoReq{
				Title: "Lorem ipsum",
			},
		},
		{
			name: "max_title_length_20",
			req: createTodoReq{
				Title: "Lorem ipsum dolor si",
			},
		},
		{
			name: "max_description_length_100",
			req: createTodoReq{
				Title:       "Lorem ipsum",
				Description: "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut l",
			},
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			err := validator.New().Struct(tCase.req)
			require.NoError(t, err)
		})
	}

}

func TestCreateValidateError(t *testing.T) {
	cases := []struct {
		name string
		req  createTodoReq
	}{
		{
			name: "bad_no_title",
			req: createTodoReq{
				Description: "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam",
			},
		},
		{
			name: "bad_title_length_21",
			req: createTodoReq{
				Title: "Lorem ipsum dolor sit",
			},
		},
		{
			name: "bad_description_length_101",
			req: createTodoReq{
				Title:       "Lorem ipsum",
				Description: "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut la",
			},
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			err := validator.New().Struct(tCase.req)
			require.Error(t, err)
		})
	}

}
