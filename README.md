# 🎟 Event Explorer

Event Explorer is a web application that helps users discover upcoming events in their selected city. Users can search for cities using Google Places autocomplete, select a location, and explore available music and sports events powered by the Ticketmaster API.

The application is built with Go and Beego, with a lightweight Tailwind CSS-based frontend.

---

# Features

## 🌎 City Search with Google Places API

- Real-time city autocomplete suggestions
- Uses Google Places API (New)
- Retrieves selected location details using place ID
- Automatically extracts:
  - City name
  - Two-letter country code

Flow:

```
User enters city
        |
        ↓
Google Places Autocomplete API
        |
        ↓
Select location
        |
        ↓
Google Place Details API
        |
        ↓
City + Country Code
```

---

## 🎫 Event Discovery

Users can search events based on:

- City
- Country code

The application retrieves:

- Music events
- Sports events

from the Ticketmaster Discovery API.

---

## ⚡ Concurrent Event Fetching

Music and sports events are fetched concurrently using Go goroutines.

Architecture:

```
                User Request
                     |
                     |
              Event Controller
                     |
             Event Service Layer
              /              \
             /                \
        Music API          Sports API
             \                /
              \              /
              Event Response
```

This reduces waiting time compared with sequential API calls.

---

## 🚀 Event Details

Users can view:

- Event name
- Event image
- Date
- Venue
- Description
- Ticket purchase link

---

## 💾 Event Caching

The application includes an in-memory cache system.

Cached data:

- City
- Country
- Category
- Events

Cache duration:

```
5 minutes
```

This reduces unnecessary Ticketmaster API requests.

---

# Tech Stack

## Backend

| Technology | Purpose |
|---|---|
| Go | Backend programming language |
| Beego | Web framework |
| REST APIs | External service communication |
| Goroutines | Concurrent API calls |

---

## Frontend

| Technology | Purpose |
|---|---|
| HTML Templates | Server-side rendering |
| Tailwind CSS | UI styling |
| JavaScript | Autocomplete interaction |

---

## External APIs

### Google Places API (New)

Used for:

- City autocomplete
- Place details lookup

Endpoints:

```
POST /v1/places:autocomplete

GET /v1/places/{placeId}
```

---

### Ticketmaster Discovery API

Used for:

- Music event search
- Sports event search
- Event details retrieval

---

# Project Architecture

```
event-explorer/

│
├── config/
│   └── config.go
│
├── controllers/
│   ├── api.go
│   ├── event.go
│   ├── details.go
│   └── home.go
│
├── models/
│   ├── event.go
│   ├── location.go
│   └── api.go
│
├── services/
│   ├── google_places.go
│   ├── ticketmaster.go
│   ├── event_service.go
│   ├── event_details.go
│   ├── cache.go
│   └── http_client.go
│
├── routers/
│   └── router.go
│
├── static/
│   ├── css/
│   └── js/
│
├── views/
│   ├── home.tpl
│   ├── listing.tpl
│   ├── details.tpl
│   └── partials/
│
├── main.go
└── go.mod
```

---

# Application Flow

```
                Browser

                   |
                   |

              Home Page

                   |
                   |

        Google Places Autocomplete

                   |
                   |

          Selected Place ID

                   |
                   |

        Google Place Details API

                   |
                   |

        City + Country Code

                   |
                   |

          Event Controller

                   |
                   |

          Ticketmaster API

                   |
                   |

        Music + Sports Events

                   |
                   |

             Event Listing

                   |
                   |

             Event Details
```

---

# Installation and Setup

## Requirements

Install:

- Go 1.25+
- Git


---

## Clone Repository

```bash
git clone https://github.com/gazimaksudur2/w3-a5.git

cd event-explorer
```

---

## Environment Configuration

Create a `.env` file in the project root:

```
GOOGLE_API_KEY=your_google_places_api_key

TICKETMASTER_API_KEY=your_ticketmaster_api_key
```

---

# Running the Application

Install dependencies:

```bash
go mod tidy
```

Run:

```bash
go run .
```

Application will start at:

```
http://localhost:8080
```

---

# API Routes

## Web Routes

### Home

```
GET /
```

Displays the city search page.

---

### Search Events

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

### Event Details

```
GET /events/{eventId}
```

---

## Internal API Routes

### City Autocomplete

```
GET /api/locations/autocomplete
```

Parameters:

```
input
sessionToken
```

---

### Place Details

```
GET /api/locations/{placeId}
```

Parameters:

```
sessionToken
```

Returns:

```json
{
    "city":"Toronto",
    "countryCode":"CA"
}
```

---

# Environment Variables

| Variable | Description |
|-|-|
| GOOGLE_API_KEY | Google Places API key |
| TICKETMASTER_API_KEY | Ticketmaster API key |

---

# Current Development Status

Completed:

✅ Beego backend setup  
✅ Ticketmaster event integration  
✅ Google Places city autocomplete  
✅ Place details lookup  
✅ City and country extraction  
✅ Concurrent event fetching  
✅ Event caching  
✅ Event details page  
✅ Tailwind UI integration  


---

# Future Improvements

Planned improvements:

- Better event card grid layout
- Loading animations
- Improved error handling
- Persistent database storage
- User accounts and favorites
- Advanced event filtering
- Deployment configuration

---

# License

This project is developed for educational and demonstration purposes.