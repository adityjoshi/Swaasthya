<div align="center">
  <img src="https://github.com/user-attachments/assets/3a8732ac-01b1-4432-9fdb-090fcf7a06ae" alt="logo">
</div>


[Swaasthya](http://github.com/adityjoshi/Swaasthya) is a comprehensive healthcare management platform designed to improve the efficiency of hospitals. It provides solutions for managing patient flow, bed availability, patient admissions, and medical inventory in real-time, with future expansions planned for enhanced data streaming and real-time analytics.

## Table of Contents

- [Features](#features)
- [Project Architecture](#project-architecture)
- [Technology Stack](#technology-stack)
- [Installation](#installation)
- [Running the Application](#running-the-application)
- [Upcoming Features](#upcoming-features)

## Features

- **Patient Flow Management**: Optimizes appointment scheduling and reduces wait times.
- **Real-Time Bed Tracking**: Monitors bed availability across departments.
- **Smart Inventory Management**: Tracks and updates consumable stock levels.
- **Emergency Preparedness Dashboard**: Displays real-time updates on patient loads, available resources, and emergency protocols.
- **2-Factor Authentication**: Secure login with OTP verification after multiple attempts.
- **Multilingual Support**: Supports multiple languages for diverse hospital environments.
- **Real-Time Notifications**: Alerts for bed availability, appointment reminders, and emergencies.
- **Offline Functionality**: Operates in low-bandwidth environments, syncing data when connectivity is restored.

## Project Architecture

Swaasthya follows a scalable microservices architecture, ensuring efficient handling of healthcare operations.

- **Backend**: Golang manages the core services and APIs.
- **Frontend**: React.js provides a responsive user interface for hospital staff and administrators.
- **Data Storage**: PostgreSQL stores patient data, inventory levels, and appointment details.
- **OTP & Notification Services**: Redis is used for caching, OTP generation, and real-time notifications.

### Components:
1. **Authentication**: User login, session management, and two-factor authentication (2FA).
2. **OPD Management**: Manages patient appointments and OPD scheduling.
3. **Bed Management**: Real-time tracking of bed availability and patient admissions.
4. **Inventory Management**: Tracks and manages hospital consumables and medicines.
5. **Emergency Dashboard**: Provides real-time monitoring of emergency situations and resources.

## Technology Stack

- **Backend**: Golang
- **Frontend**: React.js
- **Database**: PostgreSQL
- **Cache & OTP**: Redis
- **Real-Time Notifications**: Redis

> **Note**: Kafka will be integrated in the future for real-time data streaming and enhanced event processing.

## Installation

### Prerequisites

Make sure you have the following installed:

- [Go](https://golang.org/doc/install) (v1.18+)
- [Node.js](https://nodejs.org/en/download/) (v14+)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Redis](https://redis.io/download)

### Steps

1. **Clone the Repository**
    ```bash
    git clone https://github.com/adityjoshi/Swaasthya.git
    cd Swaasthya
    ```

2. **Backend Setup**

    - Navigate to the backend folder:
    ```bash
    cd backend
    ```

    - Install dependencies:
    ```bash
    go mod tidy
    ```

    - Set up your `.env` file with the necessary configurations (PostgreSQL, Redis):
    ```bash
    cp .env.example .env
    ```

    - Run the database migrations:
    ```bash
    go run migrations/migrate.go
    ```

    - Start the backend server:
    ```bash
    go run main.go
    ```

    Backend server will run at `http://localhost:2426`.

3. **Frontend Setup**

    - Navigate to the frontend directory:
    ```bash
    cd ../frontend
    ```

    - Install dependencies:
    ```bash
    npm install
    ```

    - Start the React application:
    ```bash
    npm start
    ```

    Frontend will run at `http://localhost:3000`.

## Running the Application

Ensure both backend and frontend services are running:

- **Backend**: `http://localhost:2426`
- **Frontend**: `http://localhost:3000`

Make sure PostgreSQL and Redis are running before starting the backend service.

## Upcoming Features

- **Kafka Integration**: In the future, Apache Kafka will be integrated to handle real-time data streaming, notifications, and event-driven architecture. This will improve scalability, analytics, and real-time reporting within the system.
