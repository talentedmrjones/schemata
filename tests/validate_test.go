package schemata_test

import (
	"errors"
	"testing"

	"github.com/cue-sh/schemata"
)

func TestValidate(t *testing.T) {
	testCases := []struct {
		expectation string
		schema      string
		value       interface{}
		err         error
	}{
		{
			expectation: "Valid user struct should validate.",
			schema:      "./user_schema.cue",
			value: struct {
				FirstName string
				LastName  string
			}{"Richard", "Jones"},
		},
		{
			expectation: "Disallowed field should NOT validate.",
			schema:      "./user_schema.cue",
			value: struct {
				FirstName  string
				MiddleName string
				LastName   string
			}{"Richard", "A", "Jones"},
			err: errors.New(`field not allowed: MiddleName`),
		},
		{
			expectation: "Empty field should NOT validate.",
			schema:      "./user_schema.cue",
			value: struct {
				FirstName string
				LastName  string
			}{"", "Jones"},
			err: errors.New(`FirstName: invalid value "" (out of bound =~"^[A-Z]{1}[a-zA-Z]{1,}")`),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.expectation, func(t *testing.T) {
			err := schemata.Validate(tc.schema, tc.value)
			if !cmpError(err, tc.err) {
				t.Errorf("err: got %v; want %v", err, tc.err)
			}
		})
	}
}
