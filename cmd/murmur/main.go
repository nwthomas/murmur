package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nwthomas/murmur/internal/ai"
	"github.com/nwthomas/murmur/internal/config"
	"github.com/nwthomas/murmur/internal/logger"
)

// Version information
var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

// CLI flags
var (
	showVersion = flag.Bool("version", false, "Show version information")
	debug       = flag.Bool("debug", false, "Enable debug mode")
	prompt      = flag.String("prompt", "", "Poem prompt (if not provided, will use interactive mode)")
)

// Poem generation messages
type poemMsg struct {
	content string
	err     error
}

type model struct {
	prompt     string
	poem       string
	loading    bool
	err        error
	aiClient   *ai.Client
	logger     *logger.Logger
}

// generatePoem generates a poem using the AI client
func (m model) generatePoem() tea.Msg {
	ctx := context.Background()
	poem, err := m.aiClient.GeneratePoem(ctx, m.prompt)
	return poemMsg{content: poem, err: err}
}

func (m model) Init() tea.Cmd {
	if m.prompt == "" {
		// Interactive mode - show prompt input
		return nil
	}
	// Direct mode - generate poem immediately
	m.loading = true
	return m.generatePoem
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case poemMsg:
		m.loading = false
		if msg.err != nil {
			m.err = msg.err
			m.logger.Error("Failed to generate poem", "error", msg.err)
		} else {
			m.poem = msg.content
			m.logger.Info("Poem generated successfully")
		}
		return m, tea.Quit

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEnter:
			if m.prompt == "" {
				// In interactive mode, start generating
				m.loading = true
				return m, m.generatePoem
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("\n❌ Error: %v\n\nPress Ctrl+C to exit.\n", m.err)
	}

	if m.loading {
		return "\n🎭 Generating your poem...\n\nPress Ctrl+C to cancel.\n"
	}

	if m.poem != "" {
		return fmt.Sprintf("\n✨ Your Poem:\n\n%s\n\nPress Ctrl+C to exit.\n", m.poem)
	}

	// Interactive prompt input
	return "\n🎭 Murmur - AI Poetry Generator\n\nEnter your poem prompt: _\n\nPress Enter to generate or Ctrl+C to exit.\n"
}

func main() {
	flag.Parse()

	// Show version information
	if *showVersion {
		fmt.Printf("murmur version %s (commit: %s, built: %s)\n", version, commit, date)
		os.Exit(0)
	}

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load configuration: %v\n", err)
		os.Exit(1)
	}

	// Override debug setting if provided via flag
	if *debug {
		cfg.Debug = true
	}

	// Initialize logger
	log := logger.New(cfg.LogLevel, cfg.Debug)
	log.Info("Starting Murmur", "version", version)

	// Initialize AI client
	aiClient := ai.NewClient(cfg.OpenAIAPIKey, cfg.Model)

	// Determine prompt
	promptText := *prompt
	if promptText == "" {
		// Interactive mode - we'll handle input in the UI
		promptText = ""
	}

	// Create and run the Bubble Tea program
	initialModel := model{
		prompt:   promptText,
		aiClient: aiClient,
		logger:   log,
	}

	p := tea.NewProgram(initialModel, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Error("Program error", "error", err)
		os.Exit(1)
	}
}
