import fs from "fs";

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