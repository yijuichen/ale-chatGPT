#!/bin/bash

# Define environment variables (if needed)
# export API_KEY="your-api-key"

# Prompt the user to input a multi-line string
echo "Please enter the string to pass to the command. Press Ctrl+D to end input:"

# Read the multi-line string from the user
MULTILINE_STRING=$(cat)

# Call the Go program with the user input
./ale-chatGPT gpt -t "$MULTILINE_STRING"

