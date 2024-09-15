import fs from "fs";

export default class FileMetrics
{
    FileContents = "";
    FileContentByteLength = 0;

    constructor(file)
    {
        let fileContentsBuffer = fs.readFileSync(file);
        this.FileContentByteLength = fileContentsBuffer.byteLength;
        this.FileContents = fileContentsBuffer.toString().trim();
    }

    GetTotalBytes()
    {
        return this.FileContentByteLength;
    }

    GetTotalLines()
    {
        let lines = this.FileContents.split("\n");
        return lines.length;
    }

    GetTotalWords()
    {
        let words = this.FileContents.split(" ");
        return words.length;
    }

    GetTotalCharacters()
    {
        return this.FileContents.length;
    }
};