package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	// Save original environment
	originalAPIKey := os.Getenv("OPENAI_API_KEY")
	originalModel := os.Getenv("OPENAI_MODEL")
	originalDebug := os.Getenv("DEBUG")
	originalLogLevel := os.Getenv("LOG_LEVEL")

	// Clean up after test
	defer func() {
		if originalAPIKey != "" {
			os.Setenv("OPENAI_API_KEY", originalAPIKey)
		} else {
			os.Unsetenv("OPENAI_API_KEY")
		}
		if originalModel != "" {
			os.Setenv("OPENAI_MODEL", originalModel)
		} else {
			os.Unsetenv("OPENAI_MODEL")
		}
		if originalDebug != "" {
			os.Setenv("DEBUG", originalDebug)
		} else {
			os.Unsetenv("DEBUG")
		}
		if originalLogLevel != "" {
			os.Setenv("LOG_LEVEL", originalLogLevel)
		} else {
			os.Unsetenv("LOG_LEVEL")
		}
	}()

	tests := []struct {
		name        string
		envVars     map[string]string
		expectError bool
		expected    *Config
	}{
		{
			name: "valid configuration",
			envVars: map[string]string{
				"OPENAI_API_KEY": "test-key",
				"OPENAI_MODEL":   "gpt-4",
				"DEBUG":          "true",
				"LOG_LEVEL":      "debug",
			},
			expectError: false,
			expected: &Config{
				OpenAIAPIKey: "test-key",
				Model:        "gpt-4",
				Debug:        true,
				LogLevel:     "debug",
				Theme:        "default",
				MaxRetries:   3,
				Timeout:      30,
			},
		},
		{
			name: "missing API key",
			envVars: map[string]string{
				"OPENAI_MODEL": "gpt-3.5-turbo",
			},
			expectError: true,
		},
		{
			name: "default values",
			envVars: map[string]string{
				"OPENAI_API_KEY": "test-key",
			},
			expectError: false,
			expected: &Config{
				OpenAIAPIKey: "test-key",
				Model:        "gpt-3.5-turbo",
				Debug:        false,
				LogLevel:     "info",
				Theme:        "default",
				MaxRetries:   3,
				Timeout:      30,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clear environment
			os.Unsetenv("OPENAI_API_KEY")
			os.Unsetenv("OPENAI_MODEL")
			os.Unsetenv("DEBUG")
			os.Unsetenv("LOG_LEVEL")

			// Set test environment variables
			for key, value := range tt.envVars {
				os.Setenv(key, value)
			}

			config, err := Load()

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if config.OpenAIAPIKey != tt.expected.OpenAIAPIKey {
				t.Errorf("Expected OpenAIAPIKey %s, got %s", tt.expected.OpenAIAPIKey, config.OpenAIAPIKey)
			}

			if config.Model != tt.expected.Model {
				t.Errorf("Expected Model %s, got %s", tt.expected.Model, config.Model)
			}

			if config.Debug != tt.expected.Debug {
				t.Errorf("Expected Debug %v, got %v", tt.expected.Debug, config.Debug)
			}

			if config.LogLevel != tt.expected.LogLevel {
				t.Errorf("Expected LogLevel %s, got %s", tt.expected.LogLevel, config.LogLevel)
			}
		})
	}
}

func TestGetEnv(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		fallback string
		envValue string
		expected string
	}{
		{
			name:     "environment variable exists",
			key:      "TEST_VAR",
			fallback: "fallback",
			envValue: "env-value",
			expected: "env-value",
		},
		{
			name:     "environment variable does not exist",
			key:      "NONEXISTENT_VAR",
			fallback: "fallback",
			envValue: "",
			expected: "fallback",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment variable if provided
			if tt.envValue != "" {
				os.Setenv(tt.key, tt.envValue)
				defer os.Unsetenv(tt.key)
			}

			result := getEnv(tt.key, tt.fallback)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}
