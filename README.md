
# DuckDuck

DuckDuck is a comprehensive Smart Sunrise Alarm Clock Application that seamlessly integrates both mobile and IoT functionalities. This capstone project aims to deliver an enhanced wake-up experience through innovative features and technologies.


## Prerequisites

### MQTT Setup:
- **Register on HiveMQ:**
    - Go to [HiveMQ Public MQTT Broker](https://www.hivemq.com/mqtt/public-mqtt-broker/).
    - Sign up for a free account to obtain the MQTT broker details.

- **Get MQTT Broker Details:**
    - After registration, you will receive an email with your MQTT broker details.
    - Note down the MQTT broker address, username, and password.

### MongoDB Atlas Setup:
- **Create MongoDB Atlas Cluster:**
    - Go to [MongoDB Atlas](https://www.mongodb.com/cloud/atlas).
    - Sign up for a free account if you don't have one.
    - Create a new cluster and follow the setup instructions.

- **Get Connection String:**
    - Once the cluster is deployed, go to the cluster dashboard.
    - Click on "Connect" and choose "Connect your application."
    - Copy the connection string. It should look something like:
      ```
      mongodb+srv://<username>:<password>@cluster0.mongodb.net/<database>?retryWrites=true&w=majority&ssl=true
      ```
      Replace `<username>`, `<password>`, and `<database>` with your MongoDB Atlas credentials.

## Installation

Download Go distribution by following this link: https://go.dev/doc/install

Verify that you've installed Go
```
go version
```

## Dependencies

Clone the project

```bash
git clone https://github.com/Panitnun-6243/duckduck-server.git
```

Fetch and install dependencies

```bash
go get ./...
```

Removes any unused dependencies and adds any missing dependencies to the go.mod file

```bash
go mod tidy
```
## Running the App (Local)

Export environment variables

```bash
export MONGO_URI='<YOUR MONGODB URI>&ssl=true'
export DATABASE_NAME=<YOUR DATABASE NAME>
export JWT_SECRET_KEY=<YOUR JWT SECRET>
export MQTT_BROKER=<YOUR MQTT ADDRESS>
export MQTT_CLIENT_ID=<YOUR MQTT CLIENT ID>
export MQTT_USERNAME=<YOUR MQTT USERNAME>
export MQTT_PASSWORD=<YOUR MQTT PASSWORD>
```

Start the application

```bash
go run cmd/server/main.go
```

## Running the App (Docker)

Build container image

```bash
docker build -t duckduck-server .
```

Run container image

```bash
docker run -p 5050:5050 -e MONGO_URI='<YOUR MONGODB URI>&ssl=true' -e DATABASE_NAME=<YOUR DATABASE_NAME> -e JWT_SECRET_KEY=<YOUR JWT SECRET KEY> -e MQTT_BROKER=<YOUR MQTT BROKER ADDRESS> -e MQTT_CLIENT_ID=<YOUR MQTT CLIENT ID> -e MQTT_USERNAME=<YOUR MQTT USERNAME> -e MQTT_PASSWORD='<YOUR MQTT PASSWORD>' duckduck-server
```

## Notes

- **Environment Variables:**
  - Make sure to set the environment variables either by exporting them or by creating a `.env` file.

- **MQTT Client ID:**
  - The MQTT client ID can be anything you choose, for example, `fiber_client_id`.

- **Troubleshooting:**
  - If the application doesn't run correctly, check the MongoDB and MQTT connections before starting the Fiber app.

- **Deployment Note:**
  - Deploy the application on any platform to obtain the URL or IP. Update the connection code on device software to prevent running the server locally on a Raspberry Pi device.

## Authors

### Chutirat Suasombatpattana
- Student ID: 63130500205
- Email: chutirat.earth12@gmail.com
- Phone: 094-197-4851
- Github profile: [@Chutirat](https://github.com/chutirat)
### Anakin Thanainantha
- Student ID: 63130500231
- Email: thanainantha.anakin@gmail.com
- Phone: 084-400-1100
- Github profile: [@Black3800](https://github.com/Black3800)
### Panitnun Suvannabun
- Student ID: 63130500235
- Email: panitnun.6243@gmail.com
- Phone: 084-079-7769
- Github profile: [@Panitnun-6243](https://github.com/Panitnun-6243)

