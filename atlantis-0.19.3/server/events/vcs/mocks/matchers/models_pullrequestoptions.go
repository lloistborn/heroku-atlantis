// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"

	"github.com/petergtz/pegomock"

	models "github.com/runatlantis/atlantis/server/events/models"
)

func AnyModelsPullRequestOptions() models.PullRequestOptions {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(models.PullRequestOptions))(nil)).Elem()))
	var nullValue models.PullRequestOptions
	return nullValue
}

func EqModelsPullRequestOptions(value models.PullRequestOptions) models.PullRequestOptions {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue models.PullRequestOptions
	return nullValue
}

func NotEqModelsPullRequestOptions(value models.PullRequestOptions) models.PullRequestOptions {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue models.PullRequestOptions
	return nullValue
}

func ModelsPullRequestOptionsThat(matcher pegomock.ArgumentMatcher) models.PullRequestOptions {
	pegomock.RegisterMatcher(matcher)
	var nullValue models.PullRequestOptions
	return nullValue
}
