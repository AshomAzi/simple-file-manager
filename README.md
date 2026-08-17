# Simple File Manager

Build a command-line program that examines a directory.

The user should be able to give you a path:
```console
go run main.go C:\Users\ASHOM\Desktop
```
Your program should produce something like:

```console
===== DIRECTORY INFO =====

Path: C:\Users\ASHOM\Desktop

Files: 12
Directories: 4

Contents:

📁 Projects
📁 Downloads
📄 notes.txt
📄 todo.txt
📄 photo.jpg
```

Requirements

Your program should:

- Get the directory path from the command line.
- Check whether the path exists.
- Check whether it's actually a directory.
- Read the directory contents.
- Count the number of files.
- Count the number of directories.
- Print their names.

You'll need to investigate things such as:

- os.Args
- os.Stat()
- os.ReadDir()


Bonus 🔥🔥

For every file, display its size:

- notes.txt       2.4 KB
- photo.jpg       1.8 MB
- video.mp4       523 MB

And allow:
```console
go run main.go
```

with no path to automatically examine the current working directory.