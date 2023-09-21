package functions

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		statusError := errors.New(
			fmt.Sprintf("Error(%s): Invalid HTTP status code is given back", resp.Status),
		)
		return nil, statusError
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}
