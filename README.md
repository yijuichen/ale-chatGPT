# ale-chatGPT
This is a Command Line Interface (CLI) tool that interacts with OpenAI's API to provide various functionalities.

## Features

- **Configurable via command-line flags**: Set the configuration file path and token via CLI flags.
- **Persistent Flags**: These flags can be used across all commands.

## Installation

1. **Clone the repository**:
   ```bash
      git clone https://github.com/your-username/your-repo-name.git
      cd your-repo-name
   ```
2.  **Build the application**:
   ```bash
      go build -o gpt-cli
   ```

## Usage
1. **Export your API_KEY**
   ```bash
     export API_KEY=YOUR_API_KEY
   ```
2. **Basic Command**
   ```bash
      ./gpt-cli -t "Give me a new perspective"
   ```

## Contributing
1. Fork this repository.
2. Create a new branch: git checkout -b feature/your-feature.
3. Commit your changes: git commit -m 'Add some feature'.
4. Push to the branch: git push origin feature/your-feature.
5. Submit a pull request.