package demo

import (
	"context"
	"io"
	"net/http"
)

func Sum(a, b int) int {
	return a + b
}

func GetDummyJSON(ctx context.Context) (any, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
