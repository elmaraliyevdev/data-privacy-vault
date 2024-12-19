# Data Privacy Vault

The Data Privacy Vault is a Go-based application that tokenizes sensitive data, securely stores it in memory with AES encryption, and provides the ability to detokenize and retrieve the original data. The service ensures secure handling of sensitive information and includes authentication for added security.

---

## Features

- **Tokenization**: Converts sensitive data into tokens for secure storage.
- **Detokenization**: Retrieves original data using tokens, ensuring authorized access.
- **In-Memory Storage**: Securely stores encrypted data in memory.
- **AES Encryption**: Encrypts sensitive data before storage.
- **Authentication**: Ensures requests are authorized using an `Authorization` header.

---

## Prerequisites

1. Install **Go** (1.19 or higher) on your system.

   - [Go Installation Guide](https://golang.org/doc/install)

2. Install **curl** or a similar HTTP client (e.g., Postman) for testing API endpoints.

---

## How to Run the Application

1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/data-privacy-vault.git
   cd data-privacy-vault
   ```
2. Run the application:
   ```bash
   go run main.go
   ```
3. The server will start on http://localhost:8080.

## Testing the Application

1. Tokenize Data

Use the following curl command to tokenize data:

```bash
curl -X POST http://localhost:8080/tokenize \
-H "Authorization: Bearer valid-token" \
-H "Content-Type: application/json" \
-d '{
  "id": "req-12345",
  "data": {
    "field1": "value1",
    "field2": "value2",
    "field3": "value3"
  }
}'
```

Expected Response:

```json
{
  "id": "req-12345",
  "data": {
    "field1": "mCqK33Lz",
    "field2": "V1f8wMkb",
    "field3": "RVtgCyjg"
  }
}
```

2. Detokenize Data

Take the tokens returned from the /tokenize response and use the following curl command to detokenize them:

```bash
curl -X POST http://localhost:8080/detokenize \
-H "Authorization: Bearer valid-token" \
-H "Content-Type: application/json" \
-d '{
  "id": "req-12345",
  "data": {
    "field1": "mCqK33Lz",
    "field2": "V1f8wMkb",
    "field3": "RVtgCyjg"
  }
}'
```

Expected Response:

```json
{
  "id": "req-12345",
  "data": {
    "field1": {
      "found": true,
      "value": "value1"
    },
    "field2": {
      "found": true,
      "value": "value2"
    },
    "field3": {
      "found": true,
      "value": "value3"
    }
  }
}
```

Error Cases

Missing Authorization Header

If you make a request without the Authorization header:

```bash
curl -X POST http://localhost:8080/tokenize \
-H "Content-Type: application/json" \
-d '{
  "id": "req-12345",
  "data": {
    "field1": "value1"
  }
}'
```

Response:

```json
{
  "error": "Unauthorized: Missing Authorization header"
}
```

Invalid Authorization Token

If you use an invalid token in the Authorization header:

```bash
curl -X POST http://localhost:8080/tokenize \
-H "Authorization: Bearer invalid-token" \
-H "Content-Type: application/json" \
-d '{
  "id": "req-12345",
  "data": {
    "field1": "value1"
  }
}'
```

Response:

```json
{
  "error": "Unauthorized: Invalid token"
}
```

Token Not Found During Detokenization

If you provide a nonexistent or invalid token in the /detokenize request:

```bash
curl -X POST http://localhost:8080/detokenize \
-H "Authorization: Bearer valid-token" \
-H "Content-Type: application/json" \
-d '{
  "id": "req-12345",
  "data": {
    "field1": "invalid-token"
  }
}'
```

Response:

```json
{
  "id": "req-12345",
  "data": {
    "field1": {
      "found": false,
      "value": ""
    }
  }
}
```

Project Structure

```
data-privacy-vault/
├── handlers/       # Contains tokenize and detokenize handler logic
├── middleware/     # Authentication middleware
├── models/         # Data models (e.g., Token)
├── storage/        # In-memory storage
├── utils/          # Encryption utilities
└── main.go         # Entry point for the application
```
