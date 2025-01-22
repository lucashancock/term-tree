package tree

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	CursorRowStyle    lipgloss.Style
	FileStyle         lipgloss.Style
	FolderStyle       lipgloss.Style
	TreeStyle         lipgloss.Style
	TitleStyle        lipgloss.Style
	CursorStyle       lipgloss.Style
	SelectedFileStyle lipgloss.Style

	// Styles for the line tree specifically
	LineTreeTStyle        lipgloss.Style
	LineTreeStraightStyle lipgloss.Style
	LineTreeLStyle        lipgloss.Style
	LineTreeBlankStyle    lipgloss.Style
}

// DefaultStyles defines the default styling for the file picker.
func DefaultStyles() Styles {
	return DefaultStylesWithRenderer(lipgloss.DefaultRenderer())
}

// DefaultStylesWithRenderer defines the default styling for the file picker,
// with a given Lip Gloss renderer.
func DefaultStylesWithRenderer(r *lipgloss.Renderer) Styles {
	return Styles{
		CursorRowStyle:    r.NewStyle().Foreground(lipgloss.Color("99")).Underline(true),
		FileStyle:         r.NewStyle().Foreground(lipgloss.Color("243")),
		FolderStyle:       r.NewStyle().Foreground(lipgloss.Color("243")).Bold(true),
		TreeStyle:         r.NewStyle().PaddingTop(1).PaddingLeft(1),
		TitleStyle:        r.NewStyle().PaddingLeft(1).PaddingRight(1).MarginBottom(1).Background(lipgloss.Color("99")),
		CursorStyle:       r.NewStyle().Foreground(lipgloss.Color("99")),
		SelectedFileStyle: r.NewStyle().Foreground(lipgloss.Color("105")),

		LineTreeTStyle:        r.NewStyle().Foreground(lipgloss.Color("246")),
		LineTreeStraightStyle: r.NewStyle().Foreground(lipgloss.Color("246")),
		LineTreeLStyle:        r.NewStyle().Foreground(lipgloss.Color("246")),
		LineTreeBlankStyle:    r.NewStyle().Foreground(lipgloss.Color("246")),
	}
}
