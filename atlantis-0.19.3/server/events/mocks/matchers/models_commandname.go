// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"

	"github.com/petergtz/pegomock"
	"github.com/runatlantis/atlantis/server/events/command"
)

func AnyModelsCommandName() command.Name {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(command.Name))(nil)).Elem()))
	var nullValue command.Name
	return nullValue
}

func EqModelsCommandName(value command.Name) command.Name {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue command.Name
	return nullValue
}

func NotEqModelsCommandName(value command.Name) command.Name {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue command.Name
	return nullValue
}

func ModelsCommandNameThat(matcher pegomock.ArgumentMatcher) command.Name {
	pegomock.RegisterMatcher(matcher)
	var nullValue command.Name
	return nullValue
}