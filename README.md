# Prayers4Deaf

Prayers4Deaf is a Go project designed to provide prayer times for deaf individuals, enabling them to stay informed about prayer schedules through a potential embedded solution. This application aims to enhance accessibility for the deaf community by delivering timely notifications of prayer times.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [API Documentation](#api-documentation)
- [Contributing](#contributing)
- [License](#license)

## Features

- Fetches accurate prayer times based on geographical location.
- Supports configuration for daylight saving time adjustments.
- Returns upcoming and current prayer times.
- **Embedded System Compatibility**: This feature is an idea for future implementation to provide real-time updates through visual or tactile signals.

## Getting Started

### Prerequisites

- Go version 1.23.1 or later.
- An API key from [ipgeolocation](https://ipgeolocation.io/) for accessing geographical location data.
- An API key from [Aladhan](https://aladhan.com/) for accessing prayer time data.

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/mahmoudalnkeeb/prayers4deaf.git
   cd prayers4deaf
   ```

2. Set up your Go environment:

   ```bash
   go mod tidy
   ```

3. Create a `.env` file or set environment variables for your API keys:

   ```bash
   export IPGEO_API_KEY="your_api_key_here" && export X7X_API_KEY="your_api_key_here" && go run .
   ```

### Running the Application

To run the application, use the following command:

```bash
go run .
```

Or if you are using `run.sh`:

```bash
./path/to/run.sh
```

## Usage

The application fetches prayer times based on the user's geographical location. It can be configured for different methods of calculating prayer times.

### Functions Overview

- `GetGeoLocation()`: Retrieves the geographical location of the user.
- `GetPrayers()`: Fetches prayer times from the external API.
- `GetNextPrayer()`: Determines the next upcoming prayer time.
- `GetCurrentPrayer()`: Identifies if the current time falls within a prayer time window.

## API Documentation

For detailed API endpoints and request parameters, refer to the [Aladhan API Documentation](https://aladhan.com/prayer-times-api) & [Ipgeolocation API Documentation](https://ipgeolocation.io/documentation.html).

## Contributing

Contributions are welcome! If you have suggestions for improvements or features, please open an issue or submit a pull request.

1. Fork the project.
2. Create your feature branch:

   ```bash
   git checkout -b feature/YourFeature
   ```

3. Commit your changes:

   ```bash
   git commit -m "Add some feature"
   ```

4. Push to the branch:

   ```bash
   git push origin feature/YourFeature
   ```

5. Open a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.
