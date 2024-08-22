# WebSocket Project

## Overview

This project includes a WebSocket server and a client application. The WebSocket server, written in Go, reverses incoming messages and maintains a history of the last 5 messages. The client, built with HTML and CSS, allows users to interact with the server by sending messages, viewing responses, and retrieving message history.

## Features

-   **WebSocket Server:**

    -   Reverses incoming messages and returns the reversed string.
    -   Maintains a history of the last 5 messages.
    -   Responds to a "history" command by returning the last 5 messages.

-   **Client Application:**
    -   Send messages to the WebSocket server.
    -   View sent and received messages in a table with columns for sent data, received data, timestamp, and status.
    -   Retrieve and display the message history in a separate section below the table.
    -   Show WebSocket connection status and handle errors.

## Folder Structure

```

websocket-project/
│
├── server/
│ ├── main.go # Go WebSocket server code
│ ├── go.mod # Go module file
│ ├── go.sum # Go checksum file
│
├── client/
│ ├── index.html # Main HTML file for the client
│ ├── styles.css # CSS file for styling
│
├── .gitignore # Git ignore file
├── README.md # Project overview and instructions
└── Procfile # Heroku process file (optional)

```

## Requirements

-   [Go](https://golang.org/doc/install) for running the WebSocket server.
-   A modern web browser for accessing the client application.

## Setup and Run Locally

### 1. Set Up the Server

1. **Navigate to the server directory:**
    ```bash
    cd server
    ```

````

2. **Install Go modules:**

    ```bash
    go mod tidy
    ```

3. **Run the WebSocket server:**
    ```bash
    go run main.go
    ```

### 2. Set Up the Client

1. **Navigate to the client directory:**

    ```bash
    cd client
    ```

2. **Serve the client files:**
    - Use Python’s built-in HTTP server to serve the files locally:
        ```bash
        python3 -m http.server
        ```
    - Access the client application at `http://localhost:8000` in your web browser.

## Usage

1. Open the client application in your web browser.
2. Enter a message in the input box and click "Send" to send it to the WebSocket server.
3. View the sent and received messages in the table with columns for sent data, received data, timestamp, and status.
4. Click "Get History" to retrieve and display the last 5 messages separately below the table.
5. Use the "Clear Chat" button to clear the message table.
````
