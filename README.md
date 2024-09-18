# MKBWC

**MKBWC** is a command line tool created using Node.js to compute common metrics for a given file. It works well with all file formats but works well mainly for text files. 

![About the Command](/assets/images/about-command.png "About MKBWC")

## Setting it up locally

Clone the github repository to your local machine and run the below commands.

```bash
# This command installs all the dependencies for the project.
npm install

# This command runs the application with a sample text file and produces the output.
npm start
```

To run this command in a linux-like fashion i.e., `mkbwc -l -c root.txt`, perform the below steps.

```bash
# This command checks if mkbwc is already installed.
mkbwc -V
```

If the command returns a valid version, execute the below command.

```bash
# This command uninstalls the mkbwc package globally in the system.
npm uninstall -g
```

If the `mkbwc -V` command throws a command not found error, then execute the below.

```bash
# This command installs the mkbwc package globally.
npm install -g
```

## Linux-like Implementation

This linux-like implementation works because of the below configuration made in the project.

- Adding the **bin** object at top-level of the `package.json` file. This object contains a key for command `mkbwc` and the value points to the starter javascript file that must be executed, when this command is invoked in the command line.
- At the top of `index.js` (the root file to executed), a `shebang` line is added which tells the system that the file must be executed with `node` compiler. 

## Automated Testing

This project uses the `mocha` framework for automated testing. The test suites and test cases are defined under the `/test/` folder in the root directory. Run the below command to execute all the test cases.

```bash
npm run test
```

![Unit Testing](/assets/images/unit-testing.png "Unit Testing")