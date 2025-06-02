# 📦 apperr — Structured Error Handling for Go

`apperr` is a lightweight, extensible library for structured error handling in Go microservices. It provides consistent formatting, metadata support, and easy integration with both gRPC and HTTP protocols.

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
err := apperr.Wrap("CreateUser", apperr.CodeConflict, errors.New("user already exists"), map[string]any{
  "email": "user@example.com",
})
```

### Access error code

```go
code := apperr.CodeOf(err)
if code == apperr.CodeConflict {
  // handle conflict case
}
```

---

## 📡 gRPC Integration

```go
return nil, apperr.ToGRPCStatus(
  apperr.Wrap("GetUser", apperr.CodeNotFound, errors.New("not found"), map[string]any{
    "id": "123",
  }),
)
```

This will attach `errdetails.ErrorInfo` to the status and preserve error metadata.

---

## 🌐 HTTP Integration

```go
apperr.WriteHTTPError(w, apperr.Wrap("Validate", apperr.CodeInvalidArgument, errors.New("missing field"), nil))
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
apperr.Configure(
  apperr.WithDomain("billing-service"),
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
