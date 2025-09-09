# Basic Usage Examples

## Interactive Mode

Run Murmur without any arguments to enter interactive mode:

```bash
./murmur
```

This will start an interactive session where you can enter your poem prompt.

## Command Line Mode

Generate a poem directly from the command line:

```bash
./murmur -prompt="Write a haiku about the ocean"
```

## Environment Variables

Set your OpenAI API key:

```bash
export OPENAI_API_KEY="your-api-key-here"
./murmur -prompt="A poem about coding"
```

Or use a `.env` file:

```bash
# Copy the example file
cp .env.example .env

# Edit .env with your API key
# Then run
./murmur
```

## Debug Mode

Enable debug logging:

```bash
./murmur -debug -prompt="A poem about debugging"
```

## Version Information

Check the version:

```bash
./murmur -version
```

## Docker Usage

Build and run with Docker:

```bash
# Build the image
docker build -t murmur .

# Run with environment variable
docker run -it --rm -e OPENAI_API_KEY="your-key" murmur

# Run with a specific prompt
docker run -it --rm -e OPENAI_API_KEY="your-key" murmur -prompt="A poem about containers"
```

## Development Mode

For development with live reload:

```bash
# Using docker-compose
docker-compose --profile dev up murmur-dev

# Or run directly
make run
```
