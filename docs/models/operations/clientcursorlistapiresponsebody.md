# ClientCursorListAPIResponseBody


## Supported Types

### ClientCursorListResponse

```go
clientCursorListAPIResponseBody := operations.CreateClientCursorListAPIResponseBodyClientCursorListResponse(components.ClientCursorListResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch clientCursorListAPIResponseBody.Type {
	case operations.ClientCursorListAPIResponseBodyTypeClientCursorListResponse:
		// clientCursorListAPIResponseBody.ClientCursorListResponse is populated
	default:
		// Unknown type - use clientCursorListAPIResponseBody.GetUnknownRaw() for raw JSON
}
```
