// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import mock "github.com/stretchr/testify/mock"
import unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

// LabelsSvc is an autogenerated mock type for the LabelsSvc type
type LabelsSvc struct {
	mock.Mock
}

// EnsureLabelsAreApplied provides a mock function with given fields: res, labels
func (_m *LabelsSvc) EnsureLabelsAreApplied(res *unstructured.Unstructured, labels map[string]string) error {
	ret := _m.Called(res, labels)

	var r0 error
	if rf, ok := ret.Get(0).(func(*unstructured.Unstructured, map[string]string) error); ok {
		r0 = rf(res, labels)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnsureLabelsAreDeleted provides a mock function with given fields: res, labels
func (_m *LabelsSvc) EnsureLabelsAreDeleted(res *unstructured.Unstructured, labels map[string]string) error {
	ret := _m.Called(res, labels)

	var r0 error
	if rf, ok := ret.Get(0).(func(*unstructured.Unstructured, map[string]string) error); ok {
		r0 = rf(res, labels)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
