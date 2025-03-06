# Lis

(just practice project)
Lis is a simple alternative to the `ls` command-line tool, written in Go. It lists files and directories in the current directory along with their sizes and permissions, using a tabular format with color coding.

## Features

- Displays file and directory names
- Shows file sizes in KB
- Lists file permissions
- Uses color coding for better readability
  - **Green**: Regular files
  - **Blue**: Directories

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/lis.git
   cd lis
   ```
2. Build the project:
   ```sh
   go build -o lis
   ```
3. Move the binary to a directory in your `$PATH` for global usage:
   ```sh
   mv lis /usr/local/bin/
   ```

## Usage

Simply run:
```sh
lis
```
This will list files and directories in the current directory with their size and permissions.

## Example Output
```
Names           Size(KB)   Permissions  
--------------------------------------
file1.txt       1.20       -rw-r--r--  
directory1      4.00       drwxr-xr-x  
file2.go        2.45       -rw-r--r--  
```

## Dependencies

This project uses:
- `os` package for reading directory contents
- `fmt` for formatted output
- `log` for error handling
- `text/tabwriter` for aligned tabular output


