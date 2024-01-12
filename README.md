# A simple application to practice Structs and Packages in Go

- Navigate into the application folder on the command line and run the application using `go run .`
  - It asks for the title and content of the note.
  - Writes the information into a JSON file
  - Reads the same file and presents the JSON data on the command line
- The application has 3 packages.
  - The default `main` package
  - `note` package for all implementation related to `Note` struct
  - `file` package to read and write to file
