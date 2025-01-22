# Tree Explorer Library built with Bubbletea and Lipgloss

Highly customizable tree explorer built entirely in Go.

### Example: without directory lines

<img width="374" alt="Screenshot 2025-01-22 at 11 53 06 AM" src="https://github.com/user-attachments/assets/f24c4ba2-5943-4805-98e0-e19b76766104" />

### Example: with directory lines

<img width="374" alt="Screenshot 2025-01-22 at 11 54 30 AM" src="https://github.com/user-attachments/assets/25e054a6-ce5e-442f-a95a-10b6dd0a105c" />

See test.go for usage. There are many exposed properties on the tree object which you can set.

Examples of various customizations are:
- Keybinds
- Styles
- Tree Indentation & Characters
- Folder/File Icons
- Font color/highlight/underline/bold using lipgloss in the Styles attribute
- Tree with directory lines or without them
- Cursor for line selection
- Custom styling for selected line
- Etc.

Reach out to lucas.hancock18@gmail.com for inquiries.

### References
- Bubbletea: https://github.com/charmbracelet/bubbletea
- Lipgloss: https://github.com/charmbracelet/lipgloss

### Note
Lipgloss has a pretty nice tree generator but currently doesn't support selection/navigation. This project is my own code which uses simple iteration and string concatenation to build out the tree. It is probably not that efficient and could've been done better, but it works. If you're looking for a project that integrates with the existing lipgloss tree builder, there is a pull request for a bubble that implements selection and navigation which is yet to be merged. Check it out here: https://github.com/charmbracelet/bubbles/pull/639
