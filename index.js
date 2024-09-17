#! /usr/bin/env node

import { Command } from "commander";
import figlet from "figlet";
import chalk  from "chalk";
import FileMetrics from "./lib/FileMetrics.js";
import { isFileAccessible } from "./lib/utils.js";

const cmdTitle = chalk.bold.magentaBright;
console.log(cmdTitle(figlet.textSync("MKBWC")));
const program = new Command();
program
    .version("1.0.0")
    .description("A command line tool to compute common metrics for a given file.")
    .option("-b, --total-bytes", "Get the total number of bytes in the text file.", false)
    .option("-l, --total-lines", "Get the total number of non-empty lines in the text file.", false)
    .option("-w, --total-words", "Get the total number of non-empty words in the text file.", false)
    .option("-c, --total-chars", "Get the total number of characters in the text file. This does not include whitespaces or end-of-line characters.", false)
    .argument("<string[]>", "Input file(s) for which metrics are computed")
    .parse(process.argv);

const opts = program.opts();
const files = program.args;
files.forEach(file => {
    if(isFileAccessible(file))
    {
        let totalBytes = 0, totalLines = 0, totalWords = 0, totalChars = 0;
        let fileMetric = new FileMetrics(file);
        totalBytes = fileMetric.GetTotalBytes();
        totalLines = fileMetric.GetTotalLines();
        totalWords = fileMetric.GetTotalWords();
        totalChars = fileMetric.GetTotalCharacters();
        
        if(opts.totalBytes || opts.totalLines || opts.totalWords || opts.totalChars)
        {
            console.log(` \t ${ opts.totalBytes ? totalBytes : "" } \t ${ opts.totalLines ? totalLines : "" } \t ${ opts.totalWords ? totalWords : "" } \t ${ opts.totalChars ? totalChars : "" } \t ${file}`);
        }
        else
        {
            console.log(` \t ${ opts.totalLines ? totalLines : "" } \t ${ opts.totalWords ? totalWords : "" } \t ${ opts.totalChars ? totalChars : "" } \t ${file}`);
        }
    }
    else
    {
        console.error(` \t mkbwc: ${file}: File is either not available or not accessible.`);
    }
});