package errwrap_test

import (
	"testing"

	"github.com/flew1x/errwrap"
	"github.com/stretchr/testify/require"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
)

func TestConfigure_WithDomain(t *testing.T) {
	errwrap.Configure(errwrap.WithDomain("test-service"))

	err := errwrap.Wrap("Init", errwrap.CodeInternal, dummyError("fail"), nil)
	st := errwrap.ToGRPCStatus(err)

	details, ok := extractErrorInfo(st)
	require.True(t, ok)
	require.Equal(t, "test-service", details.Domain)
}

func TestConfigure_Override(t *testing.T) {
	errwrap.Configure(errwrap.WithDomain("another-service"))

	err := errwrap.Wrap("Load", errwrap.CodeInternal, dummyError("boom"), nil)
	st := errwrap.ToGRPCStatus(err)

	details, ok := extractErrorInfo(st)
	require.True(t, ok)
	require.Equal(t, "another-service", details.Domain)
}

func extractErrorInfo(err error) (*errdetails.ErrorInfo, bool) {
	st, ok := status.FromError(err)
	if !ok {
		return nil, false
	}

	for _, d := range st.Details() {
		if ei, ok := d.(*errdetails.ErrorInfo); ok {
			return ei, true
		}
	}

	return nil, false
}

type dummyError string

func (e dummyError) Error() string { return string(e) }
