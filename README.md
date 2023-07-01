
# ChatBot GPT-3.5-Turbo 

This README file provides an overview of a ChatBot script written in Go. The script attempts to keep some context on prior api calls, the length of context messages, can be changed in the `.env` file The script includes a single `setup` function that initializes a new `ChatBot` instance, using the OpenAI API to communicate with the GPT-3.5-Turbo model.

## Function: setup()
The `setup()` function initializes a new ChatBot instance with the necessary configurations, such as the API token, system context, and chat context role. It returns a pointer to the newly created ChatBot.

### Loading environment variables

The function first attempts to load environment variables from a `.env` file using the `godotenv` package. If the file cannot be loaded, the script will log a fatal error and exit. The expected environment variables are:

- `TOKEN`: The OpenAI API token
- `MAX_PRIOR_MESSAGES`: The maximum number of prior messages to be considered for context

### User input

If the `TOKEN` environment variable is not provided, the function prompts the user to enter it manually.

### Validating the API token

The function checks the validity of the API token using the `validateToken` function. If the token is invalid, the script will panic and exit.

### Creating a new ChatBot

The function creates a new instance of the `openai.Client` with the provided token and initializes a `ChatBot` instance with the necessary fields. The `ChatBot` structure contains:

- `apiToken`: The OpenAI API token
- `chatContext`: A `ChatContext` struct that includes the chatbot role and the maximum number of prior messages
- `client`: An instance of the `openai.Client`

Finally, the function returns a pointer to the newly created `ChatBot`.

## Usage

To use this script, ensure you have the necessary dependencies installed and the required environment variables set in a `.env` file. Then, call the `setup()` function to initialize a new ChatBot instance. This instance can then be used to interact with the GPT-4 model via the OpenAI API.

