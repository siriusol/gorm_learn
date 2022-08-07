package main

import (
	"context"
	"testing"
)

func TestGetEmployees(t *testing.T) {
	ctx := context.Background()
	Connect()
	employees, err := GetEmployees(ctx)
	if err != nil {
		panic(err)
	}
	t.Logf(JsonString(employees))

	for _, e := range employees {
		if e == nil {
			continue
		}
		if e.Extra == nil {
			e.Extra = EmployeeExtra{}
		}
		e.Extra["inner_id"] = e.Id
	}
	err = SaveEmployees(ctx, employees)
	if err != nil {
		panic(err)
	}
}
