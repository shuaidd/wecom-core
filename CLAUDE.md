# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

wecom-core is a WeCom (Enterprise WeChat) Go SDK that provides a clean, easy-to-use interface for integrating with WeCom APIs. The SDK is designed to handle 30+ business modules with automatic token management, intelligent retry logic, and extensible architecture.

## Build and Development Commands

### Running Tests
```bash
# Run all tests
go test ./...

# Run tests for a specific package
go test ./internal/auth
go test ./services/contact

# Run tests with coverage
go test -cover ./...

# Run a specific test function
go test -run TestTokenManager ./internal/auth
```

### Building
```bash
# Build the project
go build

# Verify dependencies
go mod tidy
go mod verify

# Run go vet for static analysis
go vet ./...
```

### Running Examples
```bash
# Run the basic example
go run examples/basic/main.go

# Run the contact service example
go run examples/contact/main.go
```

## Architecture Overview

### Core Design Principles

**Layered Architecture:**
- `wecom.go` - Main entry point that initializes and exposes all services
- `internal/` - Internal infrastructure (HTTP client, auth, retry, errors) - NOT exposed to users
- `pkg/` - Public interfaces (logger, cache) - Users can implement these
- `types/` - Data structures organized by business domain
- `services/` - Business logic organized by WeCom API modules

**Key Components:**

1. **Token Management (`internal/auth/token.go`)**: Automatically handles access_token lifecycle:
   - Auto-fetches tokens on first use
   - Caches tokens with 5-minute early refresh to avoid expiration
   - Thread-safe to prevent duplicate token requests
   - Auto-refreshes and retries when token expires (errcode 40014, 42001)

2. **HTTP Client (`internal/client/`)**: Unified request/response handling:
   - Integrates with TokenManager to auto-inject access_token
   - Integrates with retry.Executor for intelligent retries
   - Supports trace ID propagation via context
   - Debug mode for detailed request/response logging

3. **Retry Logic (`internal/retry/`)**: Smart retry with exponential backoff:
   - Retries on token expiration, rate limiting (errcode 45009), system busy (errcode 10001)
   - Configurable max retries and backoff parameters
   - Prevents retry storms with exponential backoff

4. **Service Pattern**: Each service (contact, oauth, security, etc.) follows the same structure:
   - `Service` struct that wraps `internal/client.Client`
   - Methods organized by functionality (e.g., user.go, department.go, tag.go)
   - Request/response types defined in `types/` package

### Request Flow

```
User calls client.Contact.GetUser()
  ↓
contact.Service.GetUser() prepares request
  ↓
internal/client.Client.Do() executes with retry logic
  ↓
retry.Executor wraps execution with retry policy
  ↓
TokenManager injects access_token (fetches if needed)
  ↓
HTTP request sent to WeCom API
  ↓
Response parsed and errors handled
  ↓
Auto-retry if token expired or rate limited
```

## Adding New Service Modules

When implementing new WeCom API modules:

1. **Create type definitions** in `types/[module-name]/`:
   - Request structs with json tags
   - Response structs matching WeCom API responses
   - Common types used across multiple endpoints

2. **Create service package** in `services/[module-name]/`:
   - `[module-name].go` - Service struct and constructor
   - `[feature].go` - Group related API methods (e.g., user.go, department.go)
   - Each method should accept `context.Context` as first parameter
   - **IMPORTANT**: Use `client.PostAndUnmarshal[T]()` or `client.GetAndUnmarshal[T]()` for API calls
     - ✅ **Correct**: `return client.PostAndUnmarshal[ResponseType](s.client, ctx, url, req)`
     - ❌ **Wrong**: `s.client.Post(ctx, url, req, &resp)` - This is the old API and will cause compilation errors
     - For endpoints with no response data (only errcode/errmsg), use `client.CommonResponse` as the type
   - Example:
     ```go
     // For endpoints returning data
     func (s *Service) GetSomething(ctx context.Context, req *SomeRequest) (*SomeResponse, error) {
         return client.PostAndUnmarshal[SomeResponse](s.client, ctx, "/api/path", req)
     }

     // For endpoints returning only errcode/errmsg
     func (s *Service) DeleteSomething(ctx context.Context, req *DeleteRequest) error {
         _, err := client.PostAndUnmarshal[client.CommonResponse](s.client, ctx, "/api/path", req)
         return err
     }
     ```

3. **Register service** in `wecom.go`:
   - Add field to `Client` struct
   - Initialize in `New()` function

4. **Follow existing patterns**:
   - Check `services/contact/` or `services/wedoc/` for reference implementation
   - Error handling: return errors directly, don't wrap unnecessarily
   - Logging: client handles logging, don't add redundant logs
   - Token management: client handles automatically, don't manually manage
   - **Always use `client.PostAndUnmarshal` or `client.GetAndUnmarshal`** - never call `s.client.Post` or `s.client.Get` directly

5. **Update documentation**:
   - **IMPORTANT**: After completing API integration, MUST update `README.md` to document the new service and its usage
   - Add the new module to the feature list
   - Provide usage examples for the key APIs
   - Update the "Current Implementation Status" section in this file (CLAUDE.md)

## Configuration and Extensibility

### Custom Logger Implementation

Implement the `pkg/logger/Logger` interface:
```go
type Logger interface {
    Debug(msg string, fields ...Field)
    Info(msg string, fields ...Field)
    Warn(msg string, fields ...Field)
    Error(msg string, fields ...Field)
}
```

### Custom Cache Implementation

Implement the `pkg/cache/Cache` interface for distributed token storage (e.g., Redis):
```go
type Cache interface {
    Get(ctx context.Context, key string) (token string, expireAt time.Time, err error)
    Set(ctx context.Context, key string, token string, expireAt time.Time) error
    Delete(ctx context.Context, key string) error
}
```

Default is in-memory cache (`internal/auth/memory_cache.go`) suitable for single-instance deployments.

## Important Context

- **Go Version**: 1.25.0
- **WeCom API**: All endpoints use HTTPS + JSON, require access_token, return unified error format
- **Token Lifetime**: 7200 seconds (2 hours), SDK refreshes 5 minutes early
- **Error Codes**: errcode=0 is success, non-zero indicates error (see `internal/errors/`)
- **Rate Limiting**: WeCom enforces API rate limits, SDK auto-retries with backoff on errcode 45009
- **Module Organization**: Services are split by functionality - large modules like contact have multiple files (user.go, department.go, tag.go, etc.)

## Current Implementation Status

Completed:
- Core framework (client, auth, retry, errors)
- Contact service (users, departments, tags, batch operations, ID conversion)
- OAuth service (authentication)
- QRCode service (enterprise invite QR codes)
- IP service (get WeCom IP ranges)
- UpDown service (upstream/downstream enterprise features)
- CorpGroup service (enterprise interconnection)
- Security service (advanced features, audit logs, device management, file DLP)
- Agent service (application management, menu management, workbench customization)
- Media service (upload/download images, temporary media, JSSDK media, async upload)
- Invoice service (query and update e-invoice status)
- External contact service (customer management, tags, group chats, moments, statistics, messaging)
- KF service (customer service - account management, contact links, servicer management, session allocation, message sending, customer basic info, upgrade service configuration)
- Email service (send emails, public mailbox management, email group management)
- Wedoc service (document management, form management, form data collection)
- Meeting service (create, update, cancel, get info, get user meeting IDs)
- ReserveMeeting service (advanced meeting management - guest management, attendee management, enrollment management, health check, waiting room management, custom links)

In Progress:
- Message service (send messages, receive callbacks, template cards)

Planned:
- 20+ additional business modules per WeCom API documentation
