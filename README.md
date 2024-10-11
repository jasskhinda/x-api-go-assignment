# X.com API Project using Go

## Introduction
This project demonstrates how to interact with X.com's API using Go. It allows users to create and delete a post (tweet) through the X.com API, focusing on OAuth authentication and API response handling.

## Setup Instructions

1. **Create an X.com Developer Account**: Visit [X.com Developer Platform](https://developer.twitter.com/), apply for a developer account, and create an app. Generate API keys (API Key, Secret Key, Bearer Token, Access Token, and Secret).

2. **Configure OAuth**: Set the Callback URL to `http://localhost:8080/callback`. Ensure **Read and Write** permissions are enabled.

3. **Run the Application**:
   - Clone the repository:
     ```bash
     git clone https://github.com/jasskhinda/x-api-go-assignment.git
     ```
   - Add your credentials to a `.env` file:
     ```plaintext
     X_CLIENT_ID=your_client_id
     X_CLIENT_SECRET=your_client_secret
     ```
   - Run the application:
     ```bash
     go run main.go
     ```
   - Open `http://localhost:8080` in the browser.

## Student Notes: Issues Faced

### Summary:
I encountered an issue with the OAuth process while working with X.com's API. I followed the assignment guidelines but was unable to get past the OAuth authentication phase.

### Steps Taken:

1. **Created an X.com Developer Account** and set up the app with all necessary API keys and permissions.
2. **Implemented OAuth Flow in Go**, where the app correctly generates the OAuth URL and redirects to X.com.
3. **Tested in multiple browsers** but kept encountering the error: **"Something went wrong. You weren’t able to give access to the App."**

### Request for Assistance:
I have followed all steps but am still facing issues with the OAuth flow. I kindly request guidance on how to resolve this issue.

Best regards,  
**Jaspal Singh**  
**Student ID: 500237233**
