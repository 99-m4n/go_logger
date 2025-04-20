# go_logger

Package that give basic structure to the standard library
[log](https://pkg.go.dev/log) package. This package allows to define logging
level at which the logging messages will be print to the standard output.

## Methods

This package implements a method to set the logging level and three methods to
print info, warning and error messages.

> [!IMPORTANT]
> Currently the package does not exit with error when printing an error message.
> For this behaviour the methos log.Fatal(err) and log.Panic(err) from the
> standard library can be used.

### go_logger.SetLevel(level int)

Sets the logging level. `level` can be one of the following options:

- go_logger.InfoLevel (default)
- go_logger.WarningLevel
- go_logger.ErrorLevel

### go_logger.Info(message string)

Prints the info `message` (as string) to the console using the standard logging
flags from the `log` standard library (`log.LstdFlags`).

### go_logger.Warning(message string)

Prints the warning `message` (as string) to the console using the standard 
logging flags from the `log` standard library (`log.LstdFlags`). 

### go_logger.Error(message string)

Prints the error `message` (as string) to the console using the standard
logging flags plus the short name of the file from the `log` standard library
(`log.LstdFlags | log.Lshortfile`).

### More logging flags

Link: https://pkg.go.dev/log#pkg-constants

