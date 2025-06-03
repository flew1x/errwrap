# 📦 errwrap — Structured Error Handling for Go

`errwrap` is a lightweight, extensible library for structured error handling in Go microservices. It provides consistent formatting, metadata support, and easy integration with both gRPC and HTTP protocols.

---

## 🚀 Installation

```bash
go get github.com/flew1x/errwrap
```

---

## ✨ Features

- Typed error codes (`ErrorCode`)
- Rich context with `operation` and `meta` fields
- Easy wrapping of internal errors
- Native support for gRPC error mapping (`status.Status`)
- JSON error output for HTTP APIs
- Domain scoping via configuration

---

## 🔧 Usage

### Create a new error

```go
err := errwrap.Wrap("CreateUser", errwrap.CodeConflict, errors.New("user already exists"), map[string]any{
  "email": "user@example.com",
})
```

### Access error code

```go
code := errwrap.CodeOf(err)
if code == errwrap.CodeConflict {
  // handle conflict case
}
```

---

## 📡 gRPC Integration

```go
return nil, errwrap.ToGRPCStatus(
  errwrap.Wrap("GetUser", errwrap.CodeNotFound, errors.New("not found"), map[string]any{
    "id": "123",
  }),
)
```

This will attach `errdetails.ErrorInfo` to the status and preserve error metadata.

---

## 🌐 HTTP Integration

```go
errwrap.WriteHTTPError(w, errwrap.Wrap("Validate", errwrap.CodeInvalidArgument, errors.New("missing field"), nil))
```

Responds with structured JSON and proper HTTP status:

```json
{
  "code": "INVALID_ARGUMENT",
  "message": "missing field",
  "operation": "Validate"
}
```

---

## ⚙️ Configuration

Set domain (for tracing, metrics, etc):

```go
errwrap.Configure(
  errwrap.WithDomain("billing-service"),
)
```

---

## 📋 Error Codes

| Code                   | Meaning                              |
| ---------------------- | ------------------------------------ |
| `CodeUnknown`          | Default/fallback error code          |
| `CodeNotFound`         | Resource does not exist              |
| `CodeInvalidArgument`  | Invalid input data                   |
| `CodeConflict`         | Conflict (duplicate, already exists) |
| `CodeUnauthenticated`  | Missing or invalid credentials       |
| `CodePermissionDenied` | Access denied                        |

You can extend this list as needed in your service.

---

## 🧪 Testing

```bash
go test -v ./...
```

Includes coverage for:

- `Wrap` and `CodeOf`
- HTTP error writing
- gRPC error status generation
- Domain configuration

---

## 📁 Project Structure

```
errwrap/
├── errwrap.go        // Core error type and wrapping
├── http.go          // HTTP helpers
├── grpc.go          // gRPC integration
├── config.go        // Global configuration
├── error_code.go    // ErrorCode constants
├── *_test.go        // Tests
```

---

## 🧑‍💻 Contributing

Issues and merge requests are welcome via GitLab. Follow standard Go practices and format code with `go fmt`.

---

## 📄 License

MIT
