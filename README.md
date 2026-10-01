# 🎟 Event Explorer

Event Explorer is a Go-based web application for discovering upcoming events by city.  
It uses Google Places API for location search and Ticketmaster API for event data.

Built with Go, Beego, REST APIs, goroutines, in-memory caching, and server-side templates.

---

# Features

- 🌎 City search using Google Places API
- 🎫 Music and sports event discovery
- ⚡ Concurrent event fetching using goroutines
- 💾 In-memory event caching with manual invalidation
- 🧪 Table-driven unit tests with mocked API servers

---

# Tech Stack

| Technology | Purpose |
|---|---|
| Go | Backend development |
| Beego | Web framework |
| REST APIs | External communication |
| Goroutines | Concurrent processing |
| In-memory Cache | API response optimization |


---

# Project Structure

```
event-explorer/

├── config/          # Configuration
├── controllers/     # HTTP handlers
├── models/          # Data models
├── services/        # Business logic and APIs
├── routers/         # Route definitions
├── views/           # Templates
├── static/          # Frontend assets
├── main.go
└── go.mod
```

---

# Setup

## Requirements

- Go 1.26+
- Google Places API Key
- Ticketmaster API Key

## Environment Variables

Create `.env`:

```env
GOOGLE_API_KEY=your_google_places_api_key
TICKETMASTER_API_KEY=your_ticketmaster_api_key
```

---

# Installation

```bash
git clone https://github.com/gazimaksudur2/w3-a5

cd event-explorer

go mod tidy
```

---

# Running Application

```bash
go run .
```

Application:

```
http://localhost:8080
```

---

# API Routes

Base URL:

```
http://localhost:8080
```

## Home

```http
GET http://localhost:8080/
```

Displays the event search page.

---

## Search Events

```http
GET http://localhost:8080/events
```

Query Parameters:

```
city
countryCode
```

Example:

```
http://localhost:8080/events?city=Toronto&countryCode=CA
```

---

## Event Details

```http
GET http://localhost:8080/events/{eventId}
```

Example:

```
http://localhost:8080/events/G6vYZ9xxx
```

---

## Location Autocomplete

```http
GET http://localhost:8080/api/locations/autocomplete
```

Query Parameters:

```
input
sessionToken
```

Example:

```
http://localhost:8080/api/locations/autocomplete?input=Toronto
```

---

## Location Details

```http
GET http://localhost:8080/api/locations/{placeId}
```

Example:

```
http://localhost:8080/api/locations/ChIJ...
```

---

# Cache Invalidation Routes

The cache can be manually invalidated using:

## Clear Entire Cache

```http
GET http://localhost:8080/cache-invalidate
```

---

## Clear By Category

```http
GET http://localhost:8080/cache-invalidate?category=Music
```

or

```
http://localhost:8080/cache-invalidate?category=Sports
```

---

## Clear By City

```http
GET http://localhost:8080/cache-invalidate?city=Toronto
```

---

## Clear By Country

```http
GET http://localhost:8080/cache-invalidate?country=CA
```

---

## Clear Using Multiple Filters

```http
GET http://localhost:8080/cache-invalidate?city=Toronto&country=CA&category=Music
```

---

# Testing

Run all tests:

```bash
go test ./...
```

Verbose output:

```bash
go test ./... -v
```

Testing includes:

- Table-driven unit tests
- Mock HTTP servers
- Service layer testing

---

# Architecture Overview

```
Browser
   |
Controllers
   |
Services
   |
   +----------------+
   |                |
Google API     Ticketmaster API
   |
Models
   |
Views
```