# MKBWC

MKBWC is a command line tool created using Go to compute common metrics for any text file.

## Running the project locally

To build and run the project, execute the following commands:

```bash
# Below command builds the project and generates the output executable file.
# The name of the executable file is main.out by default. 
# If you want a different name, you can run the build command with the -o flag.
go build -o mkbwc.out

# Below command runs the executable file.
./mkbwc.out
```

## Testing

To run the test scripts, run the below command on the terminal.

```bash
go test -v -cover
```

- The **-v** command line option enables the verbose logs to be printed as the test scripts are being executed.
- The **-cover** command line option displays the code coverage information for the test cases executed.