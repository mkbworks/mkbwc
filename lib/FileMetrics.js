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
            let lineCount = 0;
            lines.forEach(line => {
                if(line.trim() !== "")
                {
                    lineCount++;
                }
            });
            
            return lineCount;
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
                if(line.trim() !== "")
                {
                    let words = line.trim().split(" ");
                    words.forEach(word => {
                        if(word.trim() !== "")
                        {
                            wordCount++;
                        }
                    });
                }
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
                if(line.trim() !== "")
                {
                    let words = line.trim().split(" ");
                    words.forEach(word => {
                        if(word.trim() !== "")
                        {
                            word = word.trim();
                            charCount += [...word].length; 
                        }
                    });
                }
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