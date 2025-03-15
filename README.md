# SMS Gateway API

## Overview
This project provides a simple HTTP API for sending SMS messages using a GSM modem connected via a serial port. It listens for HTTP `POST` requests and forwards SMS messages to the configured modem.

## Features
- RESTful API endpoint for sending SMS
- Serial communication with a GSM modem
- Configurable port and baud rate
- Logs sent messages for tracking

## Prerequisites
- A GSM modem connected via USB (e.g., `/dev/ttyUSB0`)
- Go installed on your system
- The `github.com/tarm/serial` package

## Installation
1. Clone this repository:
   ```sh
   git clone https://github.com/your-repo/sms-gateway.git
   cd sms-gateway
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Set up environment variables (optional):
   ```sh
   export PORT=8080
   ```

## Usage
### Start the Server
Run the application using:
```sh
 go run main.go
```
The server will start on the default port `8080` unless overridden by the `PORT` environment variable.

### Send an SMS
Make a `POST` request to the `/send-sms` endpoint with a JSON payload:
```json
{
  "to": "+1234567890",
  "message": "Hello, this is a test SMS!"
}
```
#### Example using `curl`:
```sh
curl -X POST http://localhost:8080/send-sms \
     -H "Content-Type: application/json" \
     -d '{"to": "+1234567890", "message": "Hello, this is a test SMS!"}'
```

## Configuration
The serial port configuration can be adjusted in the `sendSMS` function:
```go
c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}
```
Modify `Name` to match your device's port (e.g., `/dev/ttyS0` or `COM3` on Windows).

## Error Handling
- If the serial port is not accessible, the application logs an error and fails to send the SMS.
- The API returns appropriate HTTP status codes:
  - `200 OK`: SMS sent successfully
  - `400 Bad Request`: Invalid JSON payload
  - `405 Method Not Allowed`: Non-POST request
  - `500 Internal Server Error`: Issues with the GSM modem or serial port

## License
This project is licensed under the MIT License.

## Author
[barayiti] - [mpholouischauke@outlook.com]

