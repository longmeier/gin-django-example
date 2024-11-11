package exception

import (
	"errors"
	"gin-django-example/pkg/sentry"
)

type Block struct {
	Try     func()
	Catch   func(interface{})
	Finally func()
}

func (t Block) Do() {
	if t.Finally != nil {
		defer t.Finally()
	}
	if t.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				t.Catch(r)
			}
		}()
	}
	t.Try()
}

type MyError struct {
	Str string
}

func (e *MyError) Error() string {
	return e.Str
}
func SentryError(content string) {
	sentry.Client.CaptureException(errors.New(content))
}

/*
例子:
	exception.Block{
		Try: func() {

			return
		},
		Catch: func(e interface{}) {

		},
	}.Do()
*/
