package gen_test

import (
	"context"
	"errors"
	"testing"

	"github.com/sevein/oapi-codegen-issue/gen"
)

type server struct{}

func (s server) ValidatorValidate(ctx context.Context, request gen.ValidatorValidateRequestObject) (gen.ValidatorValidateResponseObject, error) {
	if request.Validator == "xxx" {
		return gen.ValidatorValidate404JSONResponse{
			Error:   true,
			Message: "validator 'xxx' not found",
		}, nil
	}

	err := validate()

	// TODO: include InvalidDocument in ValidatorValidate400JSONResponse.
	if errors.Is(err, errValidate) {
		return gen.ValidatorValidate400JSONResponse{}, nil
	}

	// TODO: include Genericerror in ValidatorValidate400JSONResponse.
	if err != nil {
		return gen.ValidatorValidate400JSONResponse{}, nil
	}

	return gen.ValidatorValidate200JSONResponse{Valid: true}, nil
}

func TestGen(t *testing.T) {
	srv := &server{}
	gen.NewStrictHandler(srv, []gen.StrictMiddlewareFunc{})
}

var errValidate = errors.New("validation error")

func validate() error {
	return errValidate
}
