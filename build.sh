# Below command builds the project and generates the output executable file.
# The name of the executable file is main.out by default. 
# If you want a different name, you can run the build command with a different name provided for the -o flag.
go build -o mkbwc
if test -f mkbwc; then
    echo "Project built successfully.\n"
    echo "Here's how to run the mkbwc command on the command-line.\n"
    # Below command displays the various options available for mkbwc command.
    ./mkbwc -h
else
    echo "Project build was not successful :: Output binary file not generated."
fi