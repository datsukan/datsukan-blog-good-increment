package main

import (
	"github.com/datsukan/datsukan-blog-good-lambda-frame/frame"

	"github.com/datsukan/datsukan-blog-good-core/usecase"
)

func main() {
	frame.Exec(increment)
}

func increment(articleID string) (int, error) {
	return usecase.Increment(articleID)
}
