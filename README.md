# Train Booking Go

## Overview
Train Booking Go is a web application that allows users to book train tickets easily. This project is built using Go and utilizes PostgreSQL for the database and Redis for caching.

## Features
- User authentication
- Train schedule management
- Ticket booking and cancellation
- Email notifications for bookings

## Prerequisites
- Go 1.16 or later
- PostgreSQL
- Redis
- A Gmail account for SMTP (for email notifications)

## Setup Instructions

1. **Clone the repository**
   ```
   git clone https://github.com/yourusername/train-booking-go.git
   cd train-booking-go
   ```

2. **Create a `.env` file**
   Copy the `.env.example` to `.env` and fill in the required values.
   ```
   cp .env.example .env
   ```

3. **Install dependencies**
   Make sure you have Go modules enabled and run:
   ```
   go mod tidy
   ```

4. **Run the application**
   Start the application using:
   ```
   go run main.go
   ```

5. **Access the application**
   Open your browser and go to `http://localhost:3000` to access the application.

## Usage
Follow the on-screen instructions to register, log in, and start booking your train tickets.

## Contributing
Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## License
This project is licensed under the MIT License. See the LICENSE file for details.