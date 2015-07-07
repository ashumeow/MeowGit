package main

import (
	"strings"
)

type x struct {
	Errors []error
}

func new_x(err ...error) x {
	return x{Errors: err}
}

func (m x) Error() string {
	errs := make([]string,len(m.Errors))
	for i,err:=range m.Errors {
		errs[i]=err.Error()
	}
	return strings.Join(errs,"\n")
}
