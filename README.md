# Murmur

> "The words of the poets are the murmurs of the stream, the sighing of the wind, the whisper of leaves."
>
> — Henry Wadsworth Longfellow

A beautiful Go-based CLI tool for generating poetry using AI. Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) for an elegant terminal experience.

## ✨ Features

- 🎭 **Interactive Mode**: Beautiful terminal UI for composing poetry prompts
- 🚀 **Command Line Mode**: Generate poems directly from the command line
- 🎨 **Elegant Design**: Built with Bubble Tea for a smooth terminal experience
- 🔧 **Configurable**: Environment-based configuration with sensible defaults
- 🐳 **Docker Support**: Run anywhere with Docker
- 📦 **Cross-Platform**: Builds for Linux, macOS, and Windows
- 🧪 **Well Tested**: Comprehensive test coverage and CI/CD pipeline

## 🚀 Quick Start

### Prerequisites

- Go 1.22.1 or later
- OpenAI API key

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/nwthomas/murmur.git
   cd murmur
   ```

2. **Set up environment:**
   ```bash
   cp .env.example .env
   # Edit .env and add your OpenAI API key
   ```

3. **Install dependencies:**
   ```bash
   make setup
   ```

4. **Run the application:**
   ```bash
   make run
   ```

### Using Docker

```bash
# Build and run with Docker
docker build -t murmur .
docker run -it --rm -e OPENAI_API_KEY="your-key" murmur
```

## 📖 Usage

### Interactive Mode

Start Murmur without arguments for an interactive experience:

```bash
./murmur
```

### Command Line Mode

Generate poems directly:

```bash
./murmur -prompt="Write a haiku about the ocean"
```

### Available Flags

- `-prompt`: Specify a poem prompt (enables direct mode)
- `-debug`: Enable debug logging
- `-version`: Show version information

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `OPENAI_API_KEY` | Your OpenAI API key | **Required** |
| `OPENAI_MODEL` | OpenAI model to use | `gpt-3.5-turbo` |
| `DEBUG` | Enable debug mode | `false` |
| `LOG_LEVEL` | Logging level | `info` |
| `THEME` | UI theme | `default` |
| `MAX_RETRIES` | Maximum API retries | `3` |
| `TIMEOUT` | Request timeout (seconds) | `30` |

## 🛠️ Development

### Project Structure

```
├── cmd/murmur/           # Main application entry point
├── internal/
│   ├── ai/              # AI client for OpenAI integration
│   ├── config/          # Configuration management
│   └── logger/          # Structured logging
├── examples/            # Usage examples
├── .github/workflows/   # CI/CD pipelines
├── Dockerfile           # Container configuration
├── docker-compose.yml   # Development environment
└── Makefile            # Build and development tasks
```

### Available Make Commands

```bash
make help              # Show all available commands
make run               # Run in development mode
make build             # Build the binary
make test              # Run tests
make test-coverage     # Run tests with coverage
make fmt               # Format code
make vet               # Run go vet
make lint              # Run linter (requires golangci-lint)
make docker-build      # Build Docker image
make docker-run        # Run Docker container
make clean             # Clean build artifacts
```

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run tests with race detection
make test-race
```

### Building for Different Platforms

```bash
# Build for current platform
make build

# Build for all platforms
make build-all

# Build for specific platform
make build-linux
make build-windows
make build-darwin
```

## 🐳 Docker

### Development with Docker

```bash
# Run development container with live reload
make docker-dev

# Or use docker-compose directly
docker-compose --profile dev up murmur-dev
```

### Production Docker

```bash
# Build production image
make docker-build

# Run production container
make docker-run
```

## 🚀 CI/CD

This project includes GitHub Actions workflows for:

- **Continuous Integration**: Automated testing, linting, and building
- **Multi-platform Builds**: Automatic builds for Linux, macOS, and Windows
- **Docker Publishing**: Automatic Docker image builds and publishing
- **Code Quality**: Format checking, vetting, and test coverage

## 📝 Examples

See the [examples directory](examples/) for detailed usage examples and common patterns.

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Setup

1. Run `make setup` to install dependencies
2. Copy `.env.example` to `.env` and configure
3. Run `make test` to ensure everything works
4. Make your changes and test them
5. Run `make check` before committing

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgements

- To my parents for always encouraging creativity, curiosity, and persistence
- The [Bubble Tea](https://github.com/charmbracelet/bubbletea) team for the amazing TUI framework
- The Go community for excellent tooling and practices