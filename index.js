#! /usr/bin/env node

import { Command } from "commander";
import figlet from "figlet";
import FileMetrics from "./lib/FileMetrics.js";

console.log(figlet.textSync("MKBWC"));
const program = new Command();
program
    .version("1.0.0")
    .description("A command line tool to compute common metrics for a text file.")
    .option("-b, --total-bytes", "Get the total number of bytes in the text file", false)
    .option("-l, --total-lines", "Get the total number of lines in the text file", false)
    .option("-w, --total-words", "Get the total number of words in the text file", false)
    .option("-c, --total-chars", "Get the total number of characters in the text file", false)
    .argument("<string[]>", "Input file(s) for which metrics are computed")
    .parse(process.argv);

const opts = program.opts();
const files = program.args;
files.forEach(file => {
    let totalBytes = 0, totalLines = 0, totalWords = 0, totalChars = 0;
    let fileMetric = new FileMetrics(file);
    if(opts.totalBytes)
    {
        totalBytes = fileMetric.GetTotalBytes();
    }
    
    if(opts.totalLines)
    {
        totalLines = fileMetric.GetTotalLines();
    }

    if(opts.totalWords)
    {
        totalWords = fileMetric.GetTotalWords();
    }
    
    if(opts.totalChars)
    {
        totalChars = fileMetric.GetTotalCharacters();
    }
    
    if(opts.totalBytes || opts.totalLines || opts.totalWords || opts.totalChars)
    {
        console.log(`${(totalBytes > 0) ? totalBytes : ""} ${(totalLines > 0) ? totalLines : ""} ${(totalWords > 0) ? totalWords : ""} ${(totalChars > 0) ? totalChars : ""} ${file}`);
    }
});