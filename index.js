#! /usr/bin/env node

import { Command } from "commander";
import figlet from "figlet";
import FileMetrics from "./lib/FileMetrics.js";
import { isFileAccessible } from "./lib/utils.js";

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
    if(isFileAccessible(file))
    {
        let totalBytes = 0, totalLines = 0, totalWords = 0, totalChars = 0;
        let result = Object.create(null);
        let fileMetric = new FileMetrics(file);
        totalBytes = fileMetric.GetTotalBytes();
        totalLines = fileMetric.GetTotalLines();
        totalWords = fileMetric.GetTotalWords();
        totalChars = fileMetric.GetTotalCharacters();
        
        if(opts.totalBytes || opts.totalLines || opts.totalWords || opts.totalChars)
        {
            if(opts.totalBytes)
            {
                result["byte_count"] = totalBytes;
            }

            if(opts.totalLines)
            {
                result["line_count"] = totalLines;
            }

            if(opts.totalWords)
            {
                result["word_count"] = totalWords;
            }

            if(opts.totalChars)
            {
                result["char_count"] = totalChars;
            }

            result["file"] = file;
        }
        else
        {
            result["line_count"] = totalLines;
            result["word_count"] = totalWords;
            result["char_count"] = totalChars;
            result["file"] = file;
        }

        console.table(result);
    }
    else
    {
        console.error(`mkbwc: ${file}: File is either not available or not accessible.`);
    }
});