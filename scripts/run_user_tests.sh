#!/bin/bash

# Define the path to the Ginkgo test files
TEST_FILES="./backend/tests/user_test.go"

# Run the tests using Ginkgo
go test -v $TEST_FILES
