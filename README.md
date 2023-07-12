# Go Phrase Finder

This Go script takes in a file path and a search phrase as parameters. It reads the file and writes the parts of the file containing the search phrase (bounded by blank lines) into a new file.

## Requirements

- Go programming language installed (version 1.12 or later is recommended).

## Building

### Windows

To build an executable for Windows, open your command prompt, navigate to the directory of your script and run:

```shell
set GOOS=windows GOARCH=amd64 
go build -o script.exe script.go
```

### Linux
To build an executable for Linux, open your terminal, navigate to the directory of your script and run:

```shell
GOOS=linux GOARCH=amd64 go build -o script script.go
```

### macOS
To build an executable for macOS, open your terminal, navigate to the directory of your script and run:

```shell
GOOS=darwin GOARCH=amd64 go build -o script script.go
```

### Running the script
Replace script with the name of the executable you just built (script.exe for Windows).

### Windows
In the command prompt:

```shell
script.exe input.txt "my phrase" output.txt
```

### Linux
In the terminal:

```shell
./script input.txt "my phrase" output.txt
```

### macOS
In the terminal:

```shell
./script input.txt "my phrase" output.txt
```

### Examples
Let's say we have a text file input.txt with the following content:

```
дата, час, частота
Lorem ipsum dolor sit amet consectetur adipiscing elit montes potenti, mollis pulvinar nostra nullam aenean ante fames placerat, vehicula porttitor accumsan mi maecenas nisi vel habitant.

дата, час, частота
Fringilla sollicitudin per erat dictumst nascetur scelerisque sociosqu nullam ornare placerat volutpat euismod, rhoncus penatibus cras nostra arcu mi convallis vivamus diam commodo sodales

дата, час, частота
Lorem ipsum dolor sit amet consectetur adipiscing elit montes potenti, mollis pulvinar nostra nullam aenean ante fames placerat, vehicula porttitor accumsan mi maecenas nisi vel habitant.
```

And we want to find all the blocks containing the phrase "vehicula porttitor". We run:

```shell
./script input.txt "vehicula porttitor" output.txt
```

The `output.txt` file will now contain all the blocks from `input.txt` that include "vehicula porttitor".

```
дата, час, частота
Lorem ipsum dolor sit amet consectetur adipiscing elit montes potenti, mollis pulvinar nostra nullam aenean ante fames placerat, vehicula porttitor accumsan mi maecenas nisi vel habitant.

дата, час, частота
Lorem ipsum dolor sit amet consectetur adipiscing elit montes potenti, mollis pulvinar nostra nullam aenean ante fames placerat, vehicula porttitor accumsan mi maecenas nisi vel habitant.
```
