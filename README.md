# WeatherTrack

WeatherTrack is a Go-based application that tracks severe storm warnings from the National Weather Service (NWS). It allows authenticated users to subscribe to alerts in their area based on their county or local NWS office. The application fetches real-time weather warnings and provides timely notifications to keep users informed and safe.

# Features

Real-time Weather Warnings: Fetches and displays severe storm warnings from the National Weather Service.
User Authentication: Secure authentication system for users to create accounts and manage their subscriptions.
Area-based Subscriptions: Users can subscribe to alerts specific to their county or local NWS office.
Notification System: Sends notifications to users via email or SMS when there are severe storm warnings in their subscribed areas.

# Installation

## Prerequisites

- Go 1.18 or higher
- A running PostgreSQL database instance
- An SMTP server for sending email notifications (optional, if using email notifications)
- A Twilio account for SMS notifications (optional, if using SMS notifications)

# Steps

## Clone the repository:

`git clone https://github.com/clinto-bean/weathertrack.git`
`cd weathertrack`

## Set up the PostgreSQL database:

` CREATE DATABASE weathertrack;`

## Create a .env file in the root directory and add your configuration:

`DB_HOST=your_db_host`
`DB_PORT=your_db_port`
`DB_USER=your_db_user`
`DB_PASSWORD=your_db_password`
`DB_NAME=weathertrack`

`SMTP_SERVER=your_smtp_server`
`SMTP_PORT=your_smtp_port`
`SMTP_USER=your_smtp_user`
`SMTP_PASSWORD=your_smtp_password`

## Install dependencies:

`go mod tidy`

## Run database migrations:

`go run cmd/migrate/main.go`

## Start the application:

`go run cmd/weathertrack/main.go`

# Usage

## Register and Authenticate

Register a new user account by sending a POST request to /register with the following JSON payload:

`{`
`    "username": "your_username",`
`    "password": "your_password",`
`    "email": "your_email"`
`}`

Authenticate by sending a POST request to /login with the following JSON payload:

`{`
`    "username": "your_username",`
`    "password": "your_password"`
`}`

The response will include a token which you need to authenticate subsequent requests.

## Subscribe to Alerts

Subscribe to alerts by sending a POST request to /subscribe with the following JSON payload:

`{`
`    "token": "your_auth_token",`
`    "county": "your_county_name",`
`    "nws_office": "your_local_nws_office"`
`}`

You will receive notifications via email or SMS based on the provided configuration in the .env file.

# Contributing

## Contributions to WeatherTrack are always welcome! Please fork the repository and submit pull requests. For major changes, please open an issue first to discuss what you would like to change.

# License

WeatherTrack is licensed under the MIT License. See the LICENSE file for more information.

# Contact

For questions or support, please open an issue on the GitHub repository or contact the maintainer at clintobean95@gmail.com.
