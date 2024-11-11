package eye

import (
	"errors"
)

func ClientError(msg string) error {
	// sentry.Client.CaptureException(errors.New(res))
	return errors.New(msg)
}
