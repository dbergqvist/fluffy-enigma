# Certificate Signing App

This application serves as an integration between Sana Learn API and Scrive API for managing course completion certificates. It allows for submitting certificates for signing and managing the signing process for both students and teachers.

## Features

- Submit certificates for multiple students
- Initiate teacher signing process
- Integration with Scrive API for document management
- Simple web interface for manual operations

## Prerequisites

- Go 1.16 or higher
- Access to Sana Learn API
- Access to Scrive API
- Valid API credentials for both services

## Setup

1. Clone the repository:
```bash
git clone <repository-url>
cd fluffy-enigma
```

2. Install dependencies:
```bash
go mod download
```

3. Configure the application:
   - Copy `config.json.example` to `config.json`
   - Update the configuration with your API credentials and settings

4. Run the application:
```bash
go run main.go
```

The application will be available at `http://localhost:8080`

## Configuration

The `config.json` file should contain the following structure:

```json
{
    "sana_api": {
        "base_url": "https://api.sana.ai",
        "access_token": "your-sana-access-token"
    },
    "scrive_api": {
        "base_url": "https://api.scrive.com",
        "access_token": "your-scrive-access-token"
    },
    "teacher_email": "teacher@example.com"
}
```

## Usage

1. Access the web interface at `http://localhost:8080`
2. Use the "Submit Certificates" form to:
   - Enter student emails (one per line)
   - Provide the certificate JSON data
3. Use the "Teacher Signing" form to:
   - Enter the document ID received from the certificate submission
   - Initiate the teacher signing process

## API Endpoints

- `POST /api/submit-certificates`: Submit certificates for signing
- `POST /api/teacher-sign`: Initiate teacher signing process

## Development

To run the application in development mode:

```bash
go run main.go
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
