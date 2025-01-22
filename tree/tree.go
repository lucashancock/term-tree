package tree

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// Model represents the tree structure, starting with a root folder
type Model struct {
	Root                      Folder
	Cursor                    int
	SelectedRow               File
	IndentWidth               int
	IndentChar                string
	IndentAfterFolderFileIcon string
	FileIcon                  string
	ClosedFolderIcon          string
	OpenFolderIcon            string
	CursorIcon                string
	KeyMap                    KeyMap
	Styles                    Styles
	Title                     string
	TitleDisabled             bool
	LineStraight              string
	LineT                     string
	LineL                     string
	LineBlank                 string
	LineTree                  bool
	LineTreeCursor            string
}

// Folder represents a folder in the tree, containing files and subfolders
type Folder struct {
	Name     string
	Expanded bool
	Files    []File
	Folders  []Folder
}

// File represents a file inside a folder
type File struct {
	Name  string
	Index int
}

type Padding int

const (
	Blank Padding = iota
	Bar
)

var NumRows int

// New initializes a new Model with the specified root folder name
func New(rootFolderName string) *Model {
	return &Model{
		Root: Folder{
			Name:     rootFolderName,
			Expanded: true,
			Files:    []File{},
			Folders:  []Folder{},
		},

		// whether or not to display the tree with directory lines or the basic one
		LineTree: true,

		// State
		Cursor:      0,
		SelectedRow: File{},
		Title:       "tree",

		// config for indent and spacing params for tree no lines
		IndentWidth:               2,
		IndentChar:                " ",
		IndentAfterFolderFileIcon: " ",

		// config for icons in the trees
		ClosedFolderIcon: "+", // ▶
		OpenFolderIcon:   "-", // ▼
		FileIcon:         "~",
		CursorIcon:       "",

		// config for line tree
		LineStraight:   "│   ",
		LineT:          "├── ",
		LineL:          "└── ",
		LineBlank:      "    ",
		LineTreeCursor: "",

		KeyMap: DefaultKeyMap(),
		Styles: DefaultStyles(),
	}
}

func (m *Model) SetTitle(title string) {
	m.Title = title
}

func (m *Model) SetTitleHidden(b bool) {
	m.TitleDisabled = b
}

// Add multiple folders at once
func (f *Folder) AddFolders(folders ...*Folder) *Folder {
	for _, folder := range folders {
		f.Folders = append(f.Folders, *folder)
	}
	return f
}

// Add multiple files at once
func (f *Folder) AddFiles(files ...File) *Folder {
	f.Files = append(f.Files, files...)
	return f
}

// Does nothing
func (m Model) Init() tea.Cmd {
	return nil
}

// CreateFolder allows the creation of a folder in a chainable fashion.
func CreateFolder(name string) *Folder {
	return &Folder{
		Name:     name,
		Expanded: true,
		Files:    []File{},
		Folders:  []Folder{},
	}
}

// CreateFile allows the creation of a file in a chainable fashion.
func CreateFile(name string, index int) File {
	return File{
		Name:  name,
		Index: index,
	}
}

// Update handles the key events and updates the model accordingly
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.KeyMap.CursorUp):
			if m.Cursor > 0 {
				m.Cursor--
			}
		case key.Matches(msg, m.KeyMap.CursorDown):
			if m.Cursor < NumRows-1 {
				m.Cursor++
			}
		case key.Matches(msg, m.KeyMap.CursorSelect):
			// m.SelectedRow = m.Cursor
			helperRowSelect(&m, &m.Root, m.Cursor)
		}
	}
	return m, nil
}

// View generates a textual representation of the entire tree structure
func (m Model) View() string {
	var row int
	var builder strings.Builder
	var titleString string

	renderFolder(m, m.Root, 0, &row, &builder)
	NumRows = row
	// whether or not to display title
	if !m.TitleDisabled {
		titleString += m.Styles.TitleStyle.Render(m.Title)
	}

	// Conditionally render based on line tree bool
	if m.LineTree {
		return titleString + m.Styles.TreeStyle.Render(helperPrintTree(m, &m.Root))
	} else {
		return titleString + m.Styles.TreeStyle.Render(builder.String())
	}
}

///////////////////////////////////////////////////////

func helperRowSelect(m *Model, folder *Folder, cursor int) {
	row := 0
	selectRow(m, folder, cursor, &row)
}

func selectRow(m *Model, folder *Folder, cursor int, row *int) {
	if cursor == *row {
		folder.Expanded = !folder.Expanded
	}
	if !m.LineTree {

		if folder.Expanded {
			for i := range folder.Folders {
				*row++
				selectRow(m, &folder.Folders[i], cursor, row)
			}
			for _, file := range folder.Files {
				*row++
				if cursor == *row {
					// view has been selected, handle it here
					m.SelectedRow = file
				}
			}

		}
	} else {
		if folder.Expanded {
			for _, file := range folder.Files {
				*row++
				if cursor == *row {
					// view has been selected, handle it here
					m.SelectedRow = file
				}
			}
			for i := range folder.Folders {
				*row++
				selectRow(m, &folder.Folders[i], cursor, row)
			}
		}
	}
}

// renderFolder renders the folder and its contents as a string
// this function is a mess but it works i think
func renderFolder(m Model, folder Folder, level int, row *int, builder *strings.Builder) {
	indent := strings.Repeat(m.IndentChar, m.IndentWidth*level)
	tab := strings.Repeat(m.IndentChar, m.IndentWidth)
	cursorTab := strings.Repeat(m.IndentChar, len(m.CursorIcon))

	if folder.Expanded {
		if m.Cursor == *row {
			builder.WriteString(m.Styles.CursorStyle.Render(m.CursorIcon) + m.IndentChar + indent + m.Styles.CursorRowStyle.Render(m.OpenFolderIcon+m.IndentAfterFolderFileIcon+folder.Name) + "\n")
		} else {
			builder.WriteString(cursorTab + m.IndentChar + indent + m.Styles.FolderStyle.Render(m.OpenFolderIcon+m.IndentAfterFolderFileIcon+folder.Name) + "\n")
		}
		*row++

		// Render subfolders and files
		for _, subFolder := range folder.Folders {
			renderFolder(m, subFolder, level+1, row, builder)
		}

		// Render files
		for _, subFile := range folder.Files {
			if m.Cursor == *row {
				builder.WriteString(m.Styles.CursorStyle.Render(m.CursorIcon) + m.IndentChar + indent + tab + m.Styles.CursorRowStyle.Render(m.FileIcon+m.IndentAfterFolderFileIcon+subFile.Name) + "\n")
			} else if m.SelectedRow == subFile {
				builder.WriteString(cursorTab + m.IndentChar + indent + tab + m.Styles.SelectedFileStyle.Render(m.FileIcon+m.IndentAfterFolderFileIcon+subFile.Name) + "\n")
			} else {
				builder.WriteString(cursorTab + m.IndentChar + indent + tab + m.Styles.FileStyle.Render(m.FileIcon+m.IndentAfterFolderFileIcon+subFile.Name) + "\n")
			}
			*row++
		}
	} else {
		if m.Cursor == *row {
			builder.WriteString(m.Styles.CursorStyle.Render(m.CursorIcon) + m.IndentChar + indent + m.Styles.CursorRowStyle.Render(m.ClosedFolderIcon+m.IndentAfterFolderFileIcon+folder.Name) + "\n")
		} else {
			builder.WriteString(cursorTab + m.IndentChar + indent + m.Styles.FolderStyle.Render(m.ClosedFolderIcon+m.IndentAfterFolderFileIcon+folder.Name) + "\n")
		}
		*row++
	}
}

func helperPrintTree(m Model, tree *Folder) string {
	row := 0
	return printTree(m, tree, &row)
}

// printTree recursively generates the tree structure as a string
// this function is even more of a mess but it also works... i think
func printTree(m Model, tree *Folder, row *int) string {
	var sb strings.Builder

	// Precompute the styles
	lineTreeBlank := m.Styles.LineTreeBlankStyle.Render(m.LineBlank)
	lineTreeStraight := m.Styles.LineTreeStraightStyle.Render(m.LineStraight)
	lineTreeL := m.Styles.LineTreeLStyle.Render(m.LineL)
	lineTreeT := m.Styles.LineTreeTStyle.Render(m.LineT)
	cursorRowStyle := m.Styles.CursorRowStyle
	cursorStyle := m.Styles.CursorStyle.Render(m.LineTreeCursor)
	folderStyle := m.Styles.FolderStyle
	fileStyle := m.Styles.FileStyle

	// Helper function for recursive printing
	var rec func(tree *Folder, prev []Padding, last bool)

	rec = func(tree *Folder, prev []Padding, last bool) {
		// Add padding based on previous state
		if len(prev) > 0 {
			for i := 0; i < len(prev)-1; i++ {
				if prev[i] == Blank {
					sb.WriteString(lineTreeBlank)
				} else if prev[i] == Bar {
					sb.WriteString(lineTreeStraight)
				}
			}
			if last {
				sb.WriteString(lineTreeL)
			} else {
				sb.WriteString(lineTreeT)
			}
		}

		// Add the current folder/file name
		if tree.Expanded {
			if m.Cursor == *row {
				sb.WriteString(cursorRowStyle.Render(m.OpenFolderIcon+m.IndentAfterFolderFileIcon+tree.Name) + cursorStyle + "\n")
			} else {
				sb.WriteString(folderStyle.Render(m.OpenFolderIcon+m.IndentAfterFolderFileIcon+tree.Name) + "\n")
			}
		} else {
			if m.Cursor == *row {
				sb.WriteString(cursorRowStyle.Render(m.ClosedFolderIcon+m.IndentAfterFolderFileIcon+tree.Name) + cursorStyle + "\n")
			} else {
				sb.WriteString(folderStyle.Render(m.ClosedFolderIcon+m.IndentAfterFolderFileIcon+tree.Name) + "\n")
			}
			*row++
			return
		}
		*row++

		// Print the files inside the folder
		for i := 0; i < len(tree.Files)-1; i++ {
			// Add the padding before printing the file name
			for i := 0; i < len(prev); i++ {
				if prev[i] == Bar {
					sb.WriteString(lineTreeStraight)
				} else {
					sb.WriteString(lineTreeBlank)
				}
			}

			if m.Cursor == *row {
				sb.WriteString(lineTreeT + cursorRowStyle.Render(m.FileIcon+m.IndentAfterFolderFileIcon+tree.Files[i].Name) + cursorStyle + "\n")
			} else if m.SelectedRow == tree.Files[i] {
				sb.WriteString(lineTreeT + m.Styles.SelectedFileStyle.Render(m.FileIcon+m.IndentAfterFolderFileIcon+tree.Files[i].Name) + "\n")
			} else {
				sb.WriteString(lineTreeT + fileStyle.Render(m.FileIcon+m.IndentAfterFolderFileIcon+tree.Files[i].Name) + "\n")
			}
			*row++
		}
		if len(tree.Files) > 0 {
			for i := 0; i < len(prev); i++ {
				if prev[i] == Bar {
					sb.WriteString(lineTreeStraight)
				} else {
					sb.WriteString(lineTreeBlank)
				}
			}
			if len(tree.Folders) != 0 {
				if m.Cursor == *row {
					sb.WriteString(lineTreeT + cursorRowStyle.Render(m.FileIcon+m.IndentAfterFolderFileIcon+tree.Files[len(tree.Files)-1].Name) + cursorStyle + "\n")
				} else if m.SelectedRow == tree.Files[len(tree.Files)-1] {
					sb.WriteString(lineTreeT + m.Styles.SelectedFileStyle.Render(m.FileIcon+m.IndentAfterFolderFileIcon+tree.Files[len(tree.Files)-1].Name) + "\n")
				} else {
					sb.WriteString(lineTreeT + fileStyle.Render(m.FileIcon+m.IndentAfterFolderFileIcon+tree.Files[len(tree.Files)-1].Name) + "\n")
				}
				*row++
			} else {
				if m.Cursor == *row {
					sb.WriteString(lineTreeL + cursorRowStyle.Render(tree.Files[len(tree.Files)-1].Name) + cursorStyle + "\n")
				} else if m.SelectedRow == tree.Files[len(tree.Files)-1] {
					sb.WriteString(lineTreeL + m.Styles.SelectedFileStyle.Render(tree.Files[len(tree.Files)-1].Name) + "\n")
				} else {
					sb.WriteString(lineTreeL + fileStyle.Render(tree.Files[len(tree.Files)-1].Name) + "\n")
				}
				*row++
			}
		}

		// If this is a directory, recursively print its subfolders
		for i, child := range tree.Folders {
			nextLast := i == len(tree.Folders)-1
			prev = append(prev, func() Padding {
				if nextLast {
					return Blank
				}
				return Bar
			}())
			rec(&child, prev, nextLast)
			prev = prev[:len(prev)-1]
		}
	}

	// Start the recursive printing
	rec(tree, nil, true)

	// Return the accumulated string
	return sb.String()
}
