
- [Introduction](#introduction)
- [Build and Run](#build-and-run)
- [Flags](#flags)
- [Adding to Terminal](#adding-to-terminal)

## Introduction
This projects aims to create a CLI in GO.
It can run parallel with the help of [goroutine](https://pkg.go.dev/github.com/dc0d/goroutines). 

- For installation and running check out [build&run](#build-and-run)
- For optional flags you can checkout [flag](#flags) 
- For using from terminal you can checkout [terminalusage](#adding-to-terminal)


## Build and Run 
```shell script
go build
./cliapp adjust.com google.com facebook.com yahoo.com yandex.com twitter.com
```
or 
```shell script
make run-default
```

## Flags
- You can use ```-parallel``` to specify how many parallel requests you want. The default is 10.
```shell script
./cliapp -parallel 5 adjust.com google.com facebook.com yahoo.com yandex.com twitter.com
```
or 
```shell script
make parallel-flag
``` 
to run with 10 parallelism

## Adding to Terminal
```shell script
sudo cp cliapp /usr/local/bin
chmod +x /usr/local/bin/cliapp 
cliapp google.com
```