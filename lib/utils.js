import fs from "fs";

export const pathExists = (path) => {
    return fs.existsSync(path);
};

export const isFile = (file) => {
    file = file.trim();
    const fileStats = fs.statSync(file);
    return fileStats.isFile();
};

export const isFileAccessible = (file) => {
    try
    {
        fs.accessSync(file, fs.constants.F_OK | fs.constants.R_OK);
        return true;
    }
    catch(err)
    {
        console.error(`Access to ${file} is not available: ${JSON.stringify(err, null, 4)}`);
        return false;
    }
};

export const outputLine = (totalBytes, totalLines, totalWords, totalChars, opts, file) => {
    let outputString = "";
    if(opts.totalBytes || opts.totalLines || opts.totalWords || opts.totalChars)
    {
        outputString = `${ opts.totalBytes ? " \t " + totalBytes.toString() + " bytes" : "" }${ opts.totalLines ? " \t " + totalLines.toString() + " lines" : "" }${ opts.totalWords ? " \t " + totalWords.toString() + " words" : "" }${ opts.totalChars ? " \t " + totalChars.toString() + " characters" : "" } \t ${file}`;
    }
    else
    {
        outputString = ` \t ${totalLines.toString() + " lines"} \t ${totalWords.toString() + " words"} \t ${totalChars.toString() + " characters"} \t ${file}`;
    }

    return outputString;
};