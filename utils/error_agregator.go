package utils

import (
	"errors"
	"strings"
)

type AggregateError struct {
	errs []string
}

func NewAggregateError() *AggregateError {
	return &AggregateError{
		errs: make([]string, 0),
	}
}

func (ae *AggregateError) Add(err error) {
	if err == nil {
		return
	}
	ae.errs = append(ae.errs, err.Error())
}

func (ae *AggregateError) GetError() error {
	if len(ae.errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(ae.errs, "\n"))
}
