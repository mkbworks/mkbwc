# MKBWC

MKBWC is a command line tool created using Go to compute common metrics for a text file. The command computes the below metrics:

- Number of lines present in the file
- Number of words present in the file
- Number of characters present in the file
- Number of bytes present in the file

This is my solution to the challenge posted at [Coding Challenges](https://codingchallenges.fyi/challenges/challenge-wc) to create my own implementation of the `wc` linux command.

## Running the project locally

To build and run the project, execute the following commands:

```bash
# Convert the build.sh file into an executable by running the below command.
chmod +x ./build.sh


# Below command builds the project and generates the executable file.
./build.sh 
```

The **Commands and Outputs** section of this file gives different ways of running the executable file with various command-line options.

## Testing

To run the test scripts, run the below command on the terminal.

```bash
go test ./lib/metrics -v -cover
```

- The **-v** command line option enables the verbose logs to be printed as the test scripts are being executed.
- The **-cover** command line option displays the code coverage information for the test cases executed.

## Commands and Outputs

This section contains various example variations for the **mkbwc** command and the outputs displayed on screen for each of these variations.

### Example One

In this example, we fetch the number of lines, words and bytes present in file - `test-one.txt`. If there are no options provided in the command line, the program computes the 3 aforementioned metrics.

**Input Command**

```bash
./mkbwc ./test-files/test-one.txt
```

**Output displayed**

```bash
1  41  285  test-one.txt
```

### Example Two

In this example, we fetch the number of lines present in file - `test-two.txt`.

**Input Command**

```bash
./mkbwc -l ./test-files/test-two.txt
```

**Output displayed**

```bash
5  test-two.txt
```

### Example Three

In this example, we fetch the number of bytes of content present in file - `test-two.txt`.

**Input Command**

```bash
./mkbwc -c ./test-files/test-two.txt
```

**Output displayed**

```bash
1663  test-two.txt
```

### Example Four

In this example, we display the help message associated with the `MKBWC` command to understand the various options available for us to use with the command.

**Input Command**

```bash
./mkbwc -h
```

**Output displayed**

```bash
Usage: ./mkbwc [options] filename(s)
Options available:
  -c	Output the number of bytes in the given file
  -h	Show the help message
  -l	Output the number of lines in the given file
  -m	Output the number of characters in the given file
  -w	Output the number of words in the given file
```

### Example Five

In this example, we chain the `mkbwc` command with the `find` command such that the output of the `find` command which yields the list of all files in the given folder, as an input to the `mkbwc` command.

**Input Command**

```bash
find "./test-files" -type f | ./mkbwc -m
```

**Output displayed**

```bash
1663  test-two.txt
285  test-one.txt
547  test-three.txt
```

Please note that, if there are any filename(s) provided with the `mkbwc` command, only those files will be processed and the input from `find` command will be ignored. So, it is recommended to mention only the flags with `mkbwc`, when it is being chained with other linux commands for getting input files.