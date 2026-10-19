# ClientSearchAPIResponseBody


## Supported Types

### ClientSearchResponse

```go
clientSearchAPIResponseBody := operations.CreateClientSearchAPIResponseBodyClientSearchResponse(components.ClientSearchResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch clientSearchAPIResponseBody.Type {
	case operations.ClientSearchAPIResponseBodyTypeClientSearchResponse:
		// clientSearchAPIResponseBody.ClientSearchResponse is populated
	default:
		// Unknown type - use clientSearchAPIResponseBody.GetUnknownRaw() for raw JSON
}
```
