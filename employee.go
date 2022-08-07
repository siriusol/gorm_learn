package main

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Employee struct {
	Id       int64         `json:"id"`
	Deptno   int32         `json:"deptno"`
	Hiredate time.Time     `json:"hiredate"`
	Extra    EmployeeExtra `json:"extra" gorm:"type:json"`
}

type EmployeeExtra map[string]interface{}

func (e EmployeeExtra) Value() (driver.Value, error) {
	return json.Marshal(e)
}

func (e EmployeeExtra) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), &e)
}

func (e *Employee) TableName() string {
	return "employee"
}

func GetEmployees(ctx context.Context) ([]*Employee, error) {
	var employees []*Employee
	err := DB.WithContext(ctx).Debug().Find(&employees).Error
	return employees, err
}

func SaveEmployees(ctx context.Context, es []*Employee) error {
	err := DB.WithContext(ctx).Debug().Save(es).Error
	return err
}

func SaveEmployee(ctx context.Context, e *Employee) error {
	err := DB.WithContext(ctx).Debug().Save(e).Error
	return err
}
