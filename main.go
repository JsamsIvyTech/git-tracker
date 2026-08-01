package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// repo track holds real data for discovered git folder
type repository struct {
	displayText string
	folderPath  string
}

type model struct {
	repositories []repository
	cursor       int
	err          error
}

func scanWorkplace(rootPath string) ([]repository, error) {
	var found []repository

	//filepath.WalkDir automatically fixes slashes (\ vs /) for Windows 11
	err := filepath.WalkDir(rootPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() && d.Name() == ".git" {
			repoDir := filepath.Dir(path) //get folder containing git
			repoName := filepath.Base(repoDir)

			//running native windows git command
			cmd := exec.Command("git", "status", "--porcelain")
			cmd.Dir = repoDir
			output, execErr := cmd.Output()

			isDirty := "CLEAN"
			if execErr == nil && len(strings.TrimSpace(string(output))) > 0 {
				isDirty = "DIRTY"
			}

			textLine := fmt.Sprintf("%-7s | %s", isDirty, repoName)

			found = append(found, repository{
				displayText: textLine,
				folderPath:  repoDir,
			})

			return filepath.SkipDir //doesnt look inside .git itself
		}
		return nil
	})
	return found, err
}

func initialModel() model {
	//targets your windows document directory
	targetDir := `C:\Users\jacks\Documents`

	realRepos, err := scanWorkplace(targetDir)

	return model{
		repositories: realRepos,
		cursor:       0,
		err:          err,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.repositories)-1 {
				m.cursor++
			}
		case "enter":
			if len(m.repositories) == 0 {
				return m, nil
			}
			//Get the target folder path of the highlighted repository
			targetFolder := m.repositories[m.cursor].folderPath
			//Execute the Windows native "explorer" command to open the folder window
			cmd := exec.Command("explorer", targetFolder)
			_ = cmd.Run()

			return m, nil
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v\nPress q to quit.", m.err)
	}

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Padding(0, 1).
		MarginBottom(0)

	cursorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#EE6FF8")).Bold(true)
	selectedStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#EE6FF8")).Underline(true)

	s := titleStyle.Render("==== Git Status Workspace Tracker ====\n\n") + "\n\n"

	if len(m.repositories) == 0 {
		s += " [ ] No Git repositories found.\n"
	}

	for i, repo := range m.repositories {
		if m.cursor == i {
			s += fmt.Sprintf("%s [x] %s\n", cursorStyle.Render(">"), selectedStyle.Render(repo.displayText))
		} else {
			s += fmt.Sprintf("   [ ] %s\n", repo.displayText)
		}
	}
	s += "\nUse 'up/down' or 'j/k' to navigate. Press 'Enter' to open folder. Press 'ctrl+c' or 'q' to quit.\n"
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
