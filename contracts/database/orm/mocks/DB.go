// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	orm "github.com/goravel/framework/contracts/database/orm"
	mock "github.com/stretchr/testify/mock"
)

// DB is an autogenerated mock type for the DB type
type DB struct {
	mock.Mock
}

// Begin provides a mock function with given fields:
func (_m *DB) Begin() (orm.Transaction, error) {
	ret := _m.Called()

	var r0 orm.Transaction
	if rf, ok := ret.Get(0).(func() orm.Transaction); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Count provides a mock function with given fields: count
func (_m *DB) Count(count *int64) error {
	ret := _m.Called(count)

	var r0 error
	if rf, ok := ret.Get(0).(func(*int64) error); ok {
		r0 = rf(count)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: value
func (_m *DB) Create(value interface{}) error {
	ret := _m.Called(value)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: value, conds
func (_m *DB) Delete(value interface{}, conds ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, value)
	_ca = append(_ca, conds...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) error); ok {
		r0 = rf(value, conds...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exec provides a mock function with given fields: sql, values
func (_m *DB) Exec(sql string, values ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, sql)
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ...interface{}) error); ok {
		r0 = rf(sql, values...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: dest, conds
func (_m *DB) Find(dest interface{}, conds ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, dest)
	_ca = append(_ca, conds...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) error); ok {
		r0 = rf(dest, conds...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// First provides a mock function with given fields: dest
func (_m *DB) First(dest interface{}) error {
	ret := _m.Called(dest)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(dest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FirstOrCreate provides a mock function with given fields: dest, conds
func (_m *DB) FirstOrCreate(dest interface{}, conds ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, dest)
	_ca = append(_ca, conds...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) error); ok {
		r0 = rf(dest, conds...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ForceDelete provides a mock function with given fields: value, conds
func (_m *DB) ForceDelete(value interface{}, conds ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, value)
	_ca = append(_ca, conds...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) error); ok {
		r0 = rf(value, conds...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: dest
func (_m *DB) Get(dest interface{}) error {
	ret := _m.Called(dest)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(dest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Group provides a mock function with given fields: name
func (_m *DB) Group(name string) orm.Query {
	ret := _m.Called(name)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(string) orm.Query); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Having provides a mock function with given fields: query, args
func (_m *DB) Having(query interface{}, args ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) orm.Query); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Join provides a mock function with given fields: query, args
func (_m *DB) Join(query string, args ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(string, ...interface{}) orm.Query); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Limit provides a mock function with given fields: limit
func (_m *DB) Limit(limit int) orm.Query {
	ret := _m.Called(limit)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(int) orm.Query); ok {
		r0 = rf(limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Model provides a mock function with given fields: value
func (_m *DB) Model(value interface{}) orm.Query {
	ret := _m.Called(value)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(interface{}) orm.Query); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Offset provides a mock function with given fields: offset
func (_m *DB) Offset(offset int) orm.Query {
	ret := _m.Called(offset)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(int) orm.Query); ok {
		r0 = rf(offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// OrWhere provides a mock function with given fields: query, args
func (_m *DB) OrWhere(query interface{}, args ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) orm.Query); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Order provides a mock function with given fields: value
func (_m *DB) Order(value interface{}) orm.Query {
	ret := _m.Called(value)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(interface{}) orm.Query); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Pluck provides a mock function with given fields: column, dest
func (_m *DB) Pluck(column string, dest interface{}) error {
	ret := _m.Called(column, dest)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(column, dest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Raw provides a mock function with given fields: sql, values
func (_m *DB) Raw(sql string, values ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, sql)
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(string, ...interface{}) orm.Query); ok {
		r0 = rf(sql, values...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Save provides a mock function with given fields: value
func (_m *DB) Save(value interface{}) error {
	ret := _m.Called(value)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Scan provides a mock function with given fields: dest
func (_m *DB) Scan(dest interface{}) error {
	ret := _m.Called(dest)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(dest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Scopes provides a mock function with given fields: funcs
func (_m *DB) Scopes(funcs ...func(orm.Query) orm.Query) orm.Query {
	_va := make([]interface{}, len(funcs))
	for _i := range funcs {
		_va[_i] = funcs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(...func(orm.Query) orm.Query) orm.Query); ok {
		r0 = rf(funcs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Select provides a mock function with given fields: query, args
func (_m *DB) Select(query interface{}, args ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) orm.Query); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Table provides a mock function with given fields: name, args
func (_m *DB) Table(name string, args ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, name)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(string, ...interface{}) orm.Query); ok {
		r0 = rf(name, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Update provides a mock function with given fields: column, value
func (_m *DB) Update(column string, value interface{}) error {
	ret := _m.Called(column, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(column, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Updates provides a mock function with given fields: values
func (_m *DB) Updates(values interface{}) error {
	ret := _m.Called(values)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(values)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Where provides a mock function with given fields: query, args
func (_m *DB) Where(query interface{}, args ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) orm.Query); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// WithTrashed provides a mock function with given fields:
func (_m *DB) WithTrashed() orm.Query {
	ret := _m.Called()

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func() orm.Query); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

type mockConstructorTestingTNewDB interface {
	mock.TestingT
	Cleanup(func())
}

// NewDB creates a new instance of DB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDB(t mockConstructorTestingTNewDB) *DB {
	mock := &DB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
