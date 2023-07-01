
# chatbot in your terminal 

Binary in go to use the openai api to create a chatbot in your terminal

## Function: New()
The `New()` returns a new Chatbot 

### Loading environment variables
The function first attempts to load environment variables from a `.env` file using the `godotenv` package. If the file cannot be loaded, the script will log a fatal error and exit. The expected environment variables are:
- `TOKEN`: The OpenAI API token
- `SYSTEM_PROMPT`: The prompt to be used for the system
- `MAX_PRIOR_MESSAGES`: The maximum number of prior messages to be considered for context

### User input
If the `TOKEN` environment variable is not provided, the function prompts the user to enter it manually.


## Usage
just build it and run it, use the .env.example to create your own .env file with your openai token
