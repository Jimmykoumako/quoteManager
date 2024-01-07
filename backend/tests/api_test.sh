#!/bin/bash

# Assuming your API server is running locally on port 8000
API_BASE_URL="http://localhost:8000"

# Function to obtain an access token
function get_access_token {
    echo "Obtaining access token..."
    # Add the command to obtain the access token from your authentication endpoint
    # Example using curl:
    access_token=$(curl -s -X POST "$API_BASE_URL/authenticate" \
        -H "Content-Type: application/json" \
        -d '{"username": "your_username", "password": "your_password"}' | jq -r '.access_token')

    echo "Access Token: $access_token"
    echo
}

# Function to make a cURL request and check the response
function test_api {
    local name="$1"
    local method="$2"
    local endpoint="$3"
    local data="$4"
    local expected_status="$5"
    local expected_body_contains="$6"

    echo "Testing $name..."
    response=$(curl -s -X "$method" "$API_BASE_URL$endpoint" \
        -H "Authorization: Bearer $access_token" \
        -H "Content-Type: application/json" \
        -d "$data")

    status_code=$(echo "$response" | jq -r '.status_code')
    body_contains=$(echo "$response" | jq -e --arg expected "$expected_body_contains" '.body | contains($expected)')

    if [ "$status_code" -eq "$expected_status" ] && [ "$body_contains" = true ]; then
        echo "$name: Passed"
    else
        echo "$name: Failed"
        echo "Response: $response"
        exit 1
    fi
}

# Obtain access token for authentication
get_access_token

# Example tests
test_api "Check Username Availability" "POST" "/api/public/users/check-username" '{"username": "testuser"}' 200 '{"available": true}'
test_api "Register Basic User Info" "POST" "/api/public/users/register/basic-info" '{"username": "testuser", "password": "testpassword"}' 201 '{"message": "Basic user info registered successfully"}'
# Additional tests

# Public route: Login User
test_api "Login User" "POST" "/api/public/users/login" '{"username": "testuser", "password": "testpassword"}' 200 '{"message": "Login successful", "access_token": ""}'

# Public route: Refresh Token
test_api "Refresh Token" "POST" "/api/public/refresh" '{"refresh_token": "your_refresh_token"}' 200 '{"message": "Token refreshed successfully", "access_token": ""}'

# Protected route: Get Quotes
test_api "Get Quotes" "GET" "/api/quotes" "" 200 '{"quotes": []}'

# Protected route: Add Quote
test_api "Add Quote" "POST" "/api/quotes" '{"text": "This is a test quote", "author": "Test Author", "category": "Test Category", "tags": ["tag1", "tag2"]}' 201 '{"message": "Quote added successfully"}'

# Protected route: Update Quote
test_api "Update Quote" "PUT" "/api/quotes/1" '{"text": "Updated quote text"}' 200 '{"message": "Quote updated successfully"}'

# Protected route: Delete Quote
test_api "Delete Quote" "DELETE" "/api/quotes/1" "" 200 '{"message": "Quote deleted successfully"}'

# Protected route: Get User Profile
test_api "Get User Profile" "GET" "/api/users/1/profile" "" 200 '{"user_id": 1, "first_name": "John", "last_name": "Doe", "email": "john.doe@example.com"}'

# Protected route: Update User Profile
test_api "Update User Profile" "PUT" "/api/users/1/profile" '{"first_name": "Updated", "last_name": "Name", "email": "updated.email@example.com"}' 200 '{"message": "User profile updated successfully"}'
