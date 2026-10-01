# 🎟 Event Explorer

Event Explorer is a Go-based web application for discovering upcoming events by city. It uses Google Places for location search and Ticketmaster API for event data.

Built with Go, Beego, REST APIs, and server-side templates.

---

## Features

- 🌎 City search with Google Places API
- 🎫 Music and sports event discovery
- ⚡ Concurrent event fetching with goroutines
- 💾 In-memory caching for API responses
- 🧪 Table-driven unit tests with mocked API servers

---

# Tech Stack

## Backend

| Technology | Purpose |
|---|---|
| Go | Backend development |
| Beego | Web framework |
| REST APIs | External communication |
| Goroutines | Concurrent processing |
| In-memory Cache | API optimization |

## External Services

| Service | Purpose |
|---|---|
| Google Places API | Location search |
| Ticketmaster API | Event data |

# Project Structure

```
event-explorer/

├── config/          # Environment configuration
├── controllers/     # HTTP request handlers
├── models/          # Data structures
├── services/        # Business logic and API integrations
├── routers/         # Application routes
├── views/           # HTML templates
├── static/          # CSS and JavaScript files
├── main.go
└── go.mod
```

---

# Setup

## Requirements

- Go 1.26+
- Google Places API key
- Ticketmaster API key

---

## Environment Configuration

Create a `.env` file in the project root:

```env
GOOGLE_API_KEY=your_google_places_api_key
TICKETMASTER_API_KEY=your_ticketmaster_api_key
```

---

# Installation

Clone the repository:

```bash
git clone https://github.com/gazimaksudur2/w3-a5

cd event-explorer
```

Install dependencies:

```bash
go mod tidy
```

---

# Running the Application

Start the server:

```bash
go run .
```

Application runs on:

```
http://localhost:8080
```

---

# Routes

## Home

```
GET /
```

Displays the event search page.

---

## Search Events

```
GET /events
```

Parameters:

```
city
countryCode
```

Example:

```
/events?city=Toronto&countryCode=CA
```

---

## Event Details

```
GET /events/{eventId}
```

Displays detailed event information.

---

## Location APIs

Autocomplete:

```
GET /api/locations/autocomplete
```

Parameters:

```
input
sessionToken
```

Location details:

```
GET /api/locations/{placeId}
```

---

# Running Tests

Run all tests:

```bash
go test ./...
```

Run tests with detailed output:

```bash
go test ./... -v
```

The project uses:
- Table-driven unit tests
- Mock HTTP servers for API testing
- Service-level testing with interfaces

---

# Architecture Overview

```
Browser
   |
   |
Controllers
   |
   |
Services
   |
   +------------+
   |            |
Google API   Ticketmaster API
   |
Models
   |
Views
```
