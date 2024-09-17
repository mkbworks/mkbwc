# MKBWC

**MKBWC** is a command line tool created using Node.js to compute common metrics for a given file. It works well with all file formats but works well mainly for text files. 

![About the Command](/assets/images/about-command.png "About MKBWC")

## Setting it up locally

To run the command locally, clone the github repository to a local machine and run the below commands.

```bash
# Install all the dependencies for the project.
npm install

# Install the command line tool globally, so that it can be used anywhere in the machine.
npm install -g
```

This global installation works because of the following configuration made in the project.

- Adding the **bin** object at top-level of the `package.json` file. This object contains a key for a command - `mkbwc` and the value for this key points to the file that must be executed, when this command is invoked in the command line.
- At the top of `index.js` (the root file to executed), a `shebang` line is added which tells the system that the file must be executed with `node` compiler. 