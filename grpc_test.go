package errwrap_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/flew1x/errwrap"
	"github.com/stretchr/testify/require"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestToGRPCStatus_WithDetails(t *testing.T) {
	errwrap.Configure(errwrap.WithDomain("test-service"))

	appErr := errwrap.Wrap("TestCreating", errwrap.CodeInvalidArgument, errors.New("invalid input"), map[string]any{
		"field": "email",
		"value": "invalid@",
	})

	grpcErr := errwrap.ToGRPCStatus(appErr)

	st, ok := status.FromError(grpcErr)
	require.True(t, ok)
	require.Equal(t, codes.InvalidArgument, st.Code())
	require.Contains(t, st.Message(), "TestCreating")

	fmt.Println("gRPC Code:", st.Code())
	fmt.Println("Message:", st.Message())

	details := st.Details()
	require.Len(t, details, 1)

	detail, ok := details[0].(*errdetails.ErrorInfo)
	require.True(t, ok)

	fmt.Println("Reason:", detail.Reason)
	fmt.Println("Domain:", detail.Domain)
	for k, v := range detail.Metadata {
		fmt.Printf("Meta[%s]: %s\n", k, v)
	}

	require.Equal(t, "INVALID_ARGUMENT", detail.Reason)
	require.Equal(t, "test-service", detail.Domain)
	require.Equal(t, "TestCreating", detail.Metadata["operataion"])
	require.Equal(t, "email", detail.Metadata["field"])
	require.Equal(t, "invalid@", detail.Metadata["value"])

	fmt.Println("GRPC Error:", grpcErr)
	fmt.Println("App Error:", appErr)
}
