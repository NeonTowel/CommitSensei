# CommitSensei

CommitSensei is a tool that enhances your commit messages with emojis, making them more expressive and fun.

## Installation

To install CommitSensei, clone the repository and build the project:

```bash
git clone https://github.com/neontowel/commitsensei.git
cd commitsensei
go build
```

## Usage

Run the application using the following command:

```bash
./commitsensei
```

## Development

To start developing with CommitSensei, ensure you have the necessary dependencies installed:

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

You can use the following task commands to build and check dependencies:

```bash
# Build the frontend
task dev:build

# Check Go dependencies
task dev:go:check
```

Start the development server:

```bash
wails dev
```

## Building the Frontend

The frontend is built using Svelte and Vite. To build the frontend, navigate to the `frontend` directory and run:

```bash
cd frontend
yarn install
yarn build
```

You can also start the development server for the frontend using:

```bash
yarn dev
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

For any questions or feedback, please contact us at [hello@neontowel.dev](mailto:hello@neontowel.dev). 