package mangodex

import (
	"context"
	"errors"
	"io"
	"net/url"
)

type ResponseType interface {
	GetResult() string
}

type Response struct {
	Result string `json:"result"`
}

func (r *Response) GetResult() string {
	return r.Result
}

// checkErrorAndResult : Helper function to check success of request by error and status code.
func checkErrorAndResult(err error, r ResponseType) error {
	switch {
	case err != nil:
		return err
	case r.GetResult() != "ok":
		return errors.New(r.GetResult())
	default:
		return nil
	}
}

// responseOp : Convenience function for simple operations that return a ResponseType.
func (dc *DexClient) responseOp(ctx context.Context, method, path string, body io.Reader, r ResponseType) error {
	u, _ := url.Parse(BaseAPI)
	u.Path = path

	// Default ResponseType will be a Response struct
	if r == nil {
		res := Response{}
		r = &res
	}

	_, err := dc.RequestAndDecode(ctx, method, u.String(), body, &r)
	return checkErrorAndResult(err, r)
}
