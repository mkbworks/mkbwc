import fs from "fs";

export default class FileMetrics
{
    FileContentsBuffer = null;

    constructor(file)
    {
        this.FileContentsBuffer = fs.readFileSync(file);
    }

    ConvertBufferToString()
    {
        if(this.FileContentsBuffer)
        {
            return this.FileContentsBuffer.toString().trim();
        }
        else
        {
            return "";
        }
    }

    GetTotalBytes()
    {
        if(this.FileContentsBuffer)
        {
            return this.FileContentsBuffer.byteLength;
        }
        else
        {
            return 0;
        }
    }

    GetTotalLines()
    {
        try
        {
            let FileContents = this.ConvertBufferToString();
            let lines = FileContents.split("\n");
            return lines.length;
        }
        catch(err)
        {
            console.error(`Error occurred while fetching total line count: ${JSON.stringify(err, null, 4)}`);
            return 0;
        }
    }

    GetTotalWords()
    {
        try
        {
            let FileContents = this.ConvertBufferToString();
            let lines = FileContents.split("\n");
            let wordCount = 0;
            lines.forEach(line => {
                let words = line.trim().split(" ");
                wordCount += words.length;
            });
            return wordCount;
        }
        catch(err)
        {
            console.error(`Error occurred while fetching total word count: ${JSON.stringify(err, null, 4)}`);
            return 0;
        }
    }

    GetTotalCharacters()
    {
        try
        {
            let FileContents = this.ConvertBufferToString();
            let lines = FileContents.split("\n");
            let charCount = 0;
            lines.forEach(line => {
                // Adding 1 to the line length to account for the newline character.
                charCount += line.length + 1;
            });
            return charCount;
        }
        catch(err)
        {
            console.error(`Error occurred while fetching total character count: ${JSON.stringify(err, null, 4)}`);
            return 0;
        }
    }
};