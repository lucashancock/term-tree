package main

import (
	"fmt"
	"tree/tree"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Tree tree.Model
}

func main() {
	// Create and populate the tree.
	t := tree.New("Root")
	t.LineTree = true
	// TEST CASE 1
	t.Root.AddFolders(
		tree.CreateFolder("empty"),
		// Create Root's folders
		tree.CreateFolder("rootfolder1").AddFolders(
			// Create rootfolder1's folders
			tree.CreateFolder("sub1").AddFiles(tree.CreateFile("subfile1", 0)),
		).AddFiles(
			// Create rootfolder1's files
			tree.CreateFile("rootsubfile1", 1),
			tree.CreateFile("rootsubfile2", 2),
		),
		tree.CreateFolder("rootfolder2").AddFolders(
			// Create rootfolder2's folders
			tree.CreateFolder("sub2").AddFiles(
				// Create sub2's files
				tree.CreateFile("sub2file1", 3),
			),
		).AddFiles(
			// Create rootfolder2's files
			tree.CreateFile("rf2file1", 4),
			tree.CreateFile("rf2file2", 5),
		),
	).AddFiles(
		// Create Root's files
		tree.CreateFile("rootfile1", 6),
		tree.CreateFile("rootfile2", 7),
	)

	// TEST CASE 2
	// t := tree.New("Root")
	// Chainable folder and file creation
	// t.Root.AddFolders(
	// 	tree.CreateFolder("empty"),
	// 	// Create Root's folders
	// 	tree.CreateFolder("projects").AddFolders(
	// 		// Create projects' folders
	// 		tree.CreateFolder("webdev").AddFiles(
	// 			// Create webdev's files
	// 			tree.CreateFile("index.html", 2),
	// 			tree.CreateFile("style.css", 3),
	// 		),
	// 	).AddFiles(
	// 		// Create projects' files
	// 		tree.CreateFile("readme.txt", 1),
	// 		tree.CreateFile("LICENSE", 0),
	// 	),
	// 	tree.CreateFolder("media").AddFolders(
	// 		// Create media's folders
	// 		tree.CreateFolder("images").AddFiles(
	// 			// Create images' files
	// 			tree.CreateFile("logo.png", 5),
	// 			tree.CreateFile("banner.jpg", 8),
	// 		),
	// 	).AddFiles(
	// 		// Create media's files
	// 		tree.CreateFile("background.mp4", 100),
	// 		tree.CreateFile("song.mp3", 4),
	// 	),
	// ).AddFiles(
	// 	// Create Root's files
	// 	tree.CreateFile("rootfile1.txt", 6),
	// 	tree.CreateFile("rootfile2.docx", 9),
	// )

	// TEST CASE 3
	// t.Root.AddFolders(
	// 	tree.CreateFolder("docs").AddFolders(
	// 		// Create docs' folders
	// 		tree.CreateFolder("guides").AddFiles(
	// 			// Create guides' files
	// 			tree.CreateFile("installation.pdf", 2),
	// 			tree.CreateFile("setup_guide.pdf", 3),
	// 		),
	// 	).AddFiles(
	// 		// Create docs' files
	// 		tree.CreateFile("manual.pdf", 7),
	// 		tree.CreateFile("terms_of_service.txt", 0),
	// 	),
	// 	tree.CreateFolder("projects").AddFolders(
	// 		// Create projects' folders
	// 		tree.CreateFolder("webdev").AddFiles(
	// 			// Create webdev's files
	// 			tree.CreateFile("index.html", 2),
	// 			tree.CreateFile("style.css", 3),
	// 		),
	// 	),
	// )

	// TEST CASE 4
	// t.Root.AddFolders(
	// 	tree.CreateFolder("Folder4").AddFiles(
	// 		tree.CreateFile("file6.java", 13),
	// 		tree.CreateFile("file7.xml", 14),
	// 	).AddFolders(
	// 		tree.CreateFolder("SubfolderE").AddFiles(
	// 			tree.CreateFile("subfileE1.md", 15),
	// 		).AddFolders(
	// 			tree.CreateFolder("SubSubfolderF").AddFiles(
	// 				tree.CreateFile("subSubfileF1.yml", 16),
	// 			),
	// 		),
	// 	),
	// 	tree.CreateFolder("Folder5").AddFolders(
	// 		tree.CreateFolder("SubfolderG").AddFiles(
	// 			tree.CreateFile("subfileG1.sql", 17),
	// 			tree.CreateFile("subfileG2.psd", 18),
	// 		),
	// 	),
	// )

	model := Model{
		Tree: *t,
	}

	// p := tea.NewProgram(model, tea.WithAltScreen())
	p := tea.NewProgram(model)

	if _, err := p.Run(); err != nil {
		fmt.Println("error:", err)
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			// Exit the program if "q" is pressed
			return m, tea.Quit
		}

		// delegate not handled msg types to update the tree.
		// enabled navigation in the tree
		var cmd tea.Cmd
		m.Tree, cmd = m.Tree.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m Model) View() string {
	// Show which File was selected
	return m.Tree.View() + fmt.Sprintf("\nSelected: %s", m.Tree.SelectedRow.Name)
}
