#! /usr/bin/env node

import { Command } from "commander";
import figlet from "figlet";
import chalk  from "chalk";
import FileMetrics from "./lib/FileMetrics.js";
import { isFileAccessible, outputLine, isFile, pathExists } from "./lib/utils.js";

const cmdTitle = chalk.bold.magentaBright;
const errorMsg = chalk.redBright;
const title = figlet.textSync("MKBWC", {
    font: "ANSI Shadow",
    horizontalLayout: "full",
    width: 80,
});
console.log(cmdTitle(title));
const program = new Command();
program
    .version("1.0.0")
    .description("A command line tool to compute common metrics for a given file.")
    .option("-b, --total-bytes", "Get the total number of bytes in the text file.", false)
    .option("-l, --total-lines", "Get the total number of non-empty lines in the text file.", false)
    .option("-w, --total-words", "Get the total number of non-empty words in the text file.", false)
    .option("-c, --total-chars", "Get the total number of characters in the text file. This does not include whitespaces or end-of-line characters.", false)
    .argument("<string>", "The list of input file(s) with a space between each pair of files.")
    .parse(process.argv);

const opts = program.opts();
const files = program.args;
let CombinedBytes = 0, CombinedLines = 0, CombinedWords = 0, CombinedChars = 0;
let AtleastOneFile = false;
files.forEach(file => {
    if(pathExists(file))
    {
        if(isFile(file))
        {
            if(isFileAccessible(file))
            {
                let fileMetric = new FileMetrics(file);
                let totalBytes = fileMetric.GetTotalBytes();
                let totalLines = fileMetric.GetTotalLines();
                let totalWords = fileMetric.GetTotalWords();
                let totalChars = fileMetric.GetTotalCharacters();
                
                console.log(outputLine(totalBytes, totalLines, totalWords, totalChars, opts, file));
        
                CombinedBytes += totalBytes;
                CombinedChars += totalChars;
                CombinedLines += totalLines;
                CombinedWords += totalWords;

                if(!AtleastOneFile)
                {
                    AtleastOneFile = true;
                }
            }
            else
            {
                console.log(errorMsg(` \t mkbwc: ${file}: File is either not available or not accessible.`));
            }
        }
        else
        {
            console.log(errorMsg(` \t mkbwc: ${file}: Given path does not point to a file.`));
        }
    }
    else
    {
        console.log(errorMsg(` \t mkbwc: ${file}: Given path does not exist.`));
    }
});

if(AtleastOneFile)
{
    console.log(outputLine(CombinedBytes, CombinedLines, CombinedWords, CombinedChars, opts, "Total"));
}