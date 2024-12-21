# MKBWC

MKBWC is a command line tool created using Go to compute common metrics for any text file. It exposes the below metrics for any given text file.

- Number of lines present in the file
- Number of words present in the file
- Number of characters present in the file
- Number of bytes present in the file

This is my solution to the challenge posted at [Coding Challenges](https://codingchallenges.fyi/challenges/challenge-wc).

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

## Commands and Outputs

This section contains various example variations for the **mkbwc** command and the outputs displayed on screen for each of these variations.

### Example One

**Input Command**

```bash
./mkbwc.out ./test-files/test-one.txt
```

**Output displayed**

```bash
1  41  285  test-one.txt
```

### Example Two

**Input Command**

```bash
./mkbwc.out ./test-files/test-two.txt
```

**Output displayed**

```bash
5  268  1663  test-two.txt
```

### Example Three

**Input Command**

```bash
./mkbwc.out -c ./test-files/test-two.txt
```

**Output displayed**

```bash
1663  test-two.txt
```

### Example Four

**Input Command**

```bash
./mkbwc.out -h
```

**Output displayed**

```bash
Usage: ./mkbwc.out [options] filename(s)
Options available:
  -c	Output the number of bytes in the given file
  -h	Show the help message
  -l	Output the number of lines in the given file
  -m	Output the number of characters in the given file
  -w	Output the number of words in the given file
```