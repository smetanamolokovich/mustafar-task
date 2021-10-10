# Mustafar task

## API endpoints:

| Method | URL Patter     |     Handler     |             Action |
| ------ | -------------- | :-------------: | -----------------: |
| GET    | /v1/value/:key | getValueHandler |   Get value by key |
| POST   | /v1/value      | setValueHandler | Create a new value |

## App structure:

```bash
├── cmd
│   └── api
│       ├── errors.go
│       ├── handlers.go
│       ├── helpers.go
│       ├── main.go
│       └── routes.go
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── Makefile
├── pkg
│   ├── data
│   │   └── data.go
│   ├── kvstore
│   │   └── kvstore.go
│   └── validator
│       └── validator.go
└── README.md

6 directories, 14 files
```

## Start app:

```bash
 make build && make run
```

## Examples:

Create dummy data:

```bash
KV_TEST_DATA='{"key":"lorem","value":"TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdC4gTnVsbGEgZ3JhdmlkYSBlZ2V0IGR1aSB2ZWwgY3Vyc3VzLiBTdXNwZW5kaXNzZSBwb3RlbnRpLiBTdXNwZW5kaXNzZSBldSBhcmN1IG5vbiBlcm9zIG9ybmFyZSBkaWN0dW0u","expires":"2021-02-02T15:04:05Z"}'
```

Send the request:

```bash
curl -i -d "$KV_TEST_DATA" localhost:8000/v1/value
```

Response:

```bash
HTTP/1.1 201
CreatedContent-Type: application/json
Date: Sun, 10 Oct 2021 06:45:37 GMT
Content-Length: 295

{
        "data": {
                "key": "lorem",
                "value": "TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdC4gTnVsbGEgZ3JhdmlkYSBlZ2V0IGR1aSB2ZWwgY3Vyc3VzLiBTdXNwZW5kaXNzZSBwb3RlbnRpLiBTdXNwZW5kaXNzZSBldSBhcmN1IG5vbiBlcm9zIG9ybmFyZSBkaWN0dW0u",
                "expires": "2021-02-02T15:04:05Z"
        }
}
```
