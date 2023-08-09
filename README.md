# My App Documentation

## Overview

This application serves as a real-time chat platform. It's constructed with React on the frontend, Go on the backend, and Postgres for database management. The sections below delve into the project's structure and significant implementation facets.

## Getting Started

1. A GitHub repository was initiated and subsequently cloned onto the local machine.
2. Two primary directories were crafted:
   - `frontend/`: Spawned with `npm install vite@latest`.
   - `backend/`: Started off with a rudimentary `main.go` file which harnesses the `net/http` package to run a server.

## Frontend Details

### Framework and Routes

- The frontend benefits from `react-router v6` for navigation. This led to the creation of several components representing pages such as:
   - Login
   - Signup
   - Rooms
   - Chat Room

### Authentication Mechanics

- The Login and Signup components showcase a form. Upon form submission:
   - A POST request journeys to the backend API.
   - This API request is processed, JSON data interpreted, passwords are either hashed (signup) or matched (login), and relevant details are stored in the Postgres database.
   - A JWT (JSON Web Token) is crafted based on the user's credentials and forwarded to the frontend.
   - Upon its reception at the frontend, this token is tucked into local storage, followed by redirecting the user to the root (`"/"`) route.
   - Middleware on the frontend inspects the JWT in the local storage, deciding either to grant access or reroute to `/login`.

### Chat Dynamics

- Post authorization, users can:
   - Craft chat rooms.
   - Engage in chat rooms.
   - View and dispatch messages.
   
   Fetching data is achieved through GET requests, with JWTs playing a pivotal role in authorization. This token is stored in a `jotai` atom.

### WebSocket-based Real-Time Interaction 

- For instantaneous features like room creation and message dispatching, the app forgoes traditional POST requests in favor of WebSockets. A snippet showcasing the WebSocket dynamics in `ChatRoom` and `Rooms` components is as follows:

```javascript
// WebSocket handlers 
const ws = new WebSocket(`${socket_url}/ws?token=${JWT}`);
```

## Backend Dynamics

### WebSocket Handling

The Go-based WebSocket handler in the backend utilizes a switch-case approach to handle various data types (`newRoom`, `newMessage`). It ensures the broadcasting of messages to all connected clients via channels.

### Styling

The application's visual appeal is attributed to CSS Modules.

### Backend Nuances

With the integration of the `chi` router, the backend API incorporates a custom JWT middleware that:

- Assesses URLs.
- Extracts and confirms the JWT from incoming request headers with a secret key.
- Provides functionalities like fetching rooms, users, and messages.

## Deployment & Security

The deployment of this application was carried out across multiple platforms:

- The Postgres database is stationed on Neon.
- The Go-centric API is hosted through Railway.
- The React frontend is served via Firebase.

For an added layer of security, environment variables are actively used both in the frontend and backend sectors.


