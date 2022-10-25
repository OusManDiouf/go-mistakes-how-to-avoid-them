package main

import (
	"errors"
	"log"
	"strings"
)

type MultiError struct {
	errs []string
}

func (m *MultiError) Add(err error) {
	m.errs = append(m.errs, err.Error())
}
func (m *MultiError) Error() string {
	return strings.Join(m.errs, ";")
}

type Customer struct {
	Age  int
	Name string
}

func (c Customer) Validate() error {
	var m *MultiError
	if c.Age < 12 {
		m = &MultiError{}
		m.Add(errors.New("age is under the legal age"))
	}
	if c.Name == "" {
		if m == nil {
			m = &MultiError{}
		}
		m.Add(errors.New("name is empty"))
	}

	// here the returned value isn't nil directly but a nil pointer.
	// because a nil pointer is a valide receiver,
	// converting the result into an interface won't yield a nil value !!
	// error interface is a dispatch wrapper
	// here the wrappee is nil (m) but the wrapper interface is not.
	// that's why the customer validation always return non-nil value.
	// return m

	// return m only if there was at least one error
	if m != nil {
		return m
	}

	// otherwise return nil explicitly
	// Hence, in the case of a valid Customer, we return a nil interface,
	// not a nil receiver converted into a non-nil interface.
	return nil
}

func main() {
	c := Customer{
		Age:  55,
		Name: "Augustus",
	}
	err := c.Validate()
	if err != nil {
		log.Fatalf("Customer is invalid: %v", err)
	}
}
