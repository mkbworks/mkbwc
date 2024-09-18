import assert from "assert";
import path from "path";
import FileMetrics from "../lib/FileMetrics.js";
import { isFileAccessible } from "../lib/utils.js";

describe("File Metrics computed for a text file", () => {
    describe("Case 1 - Text file containing only ASCII characters", () => {
        let CompleteFileName = path.join(process.cwd(), "test", "test-files", "case-one.txt");

        it("Check if the text file is accessible", () => {
            assert.ok(isFileAccessible(CompleteFileName), "Test file is either not available or not accessible");
        });

        let fileMetrics = new FileMetrics(CompleteFileName);
        
        it("Check if the total line count calculated is correct", () => {
            assert.strictEqual(fileMetrics.GetTotalLines(), 3, "Total line count does not match the expected value.");
        });

        it("Check if the total word count calculated is correct", () => {
            assert.strictEqual(fileMetrics.GetTotalWords(), 88, "Total word count does not match the expected value.");
        });

        it("Check if the total byte count calculated is correct", () => {
            assert.strictEqual(fileMetrics.GetTotalBytes(), 607, "Total byte count does not match the expected value.");
        });

        it("Check if the total character count calculated is correct", () => {
            assert.strictEqual(fileMetrics.GetTotalCharacters(), 519, "Total character count does not match the expected value.");
        });
    });

    describe("Case 2 - Text file containing a mix of ASCII and non-ASCII unicode characters", () => {
        let CompleteFileName = path.join(process.cwd(), "test", "test-files", "case-two.txt");

        it("Check if the text file is accessible", () => {
            assert.ok(isFileAccessible(CompleteFileName), "Test file is either not available or not accessible");
        });

        let fileMetrics = new FileMetrics(CompleteFileName);

        it("Check if the total byte count calculated is correct", () => {
            assert.strictEqual(fileMetrics.GetTotalBytes(), 1130, "Total byte count does not match the expected value.");
        });

        it("Check if the total character count calculated is correct", () => {
            assert.strictEqual(fileMetrics.GetTotalCharacters(), 949, "Total character count does not match the expected value.");
        });
    });

    describe("Case 3 - Empty text file", () => {
        let CompleteFileName = path.join(process.cwd(), "test", "test-files", "case-three.txt");

        it("Check if the text file is accessible", () => {
            assert.ok(isFileAccessible(CompleteFileName), "Test file is either not available or not accessible");
        });

        let fileMetrics = new FileMetrics(CompleteFileName);

        it("Check if the total line count calculated is correct", () => {
            assert.strictEqual(fileMetrics.GetTotalLines(), 0, "Total line count does not match the expected value.");
        });

        it("Check if the total word count calculated is correct", () => {
            assert.strictEqual(fileMetrics.GetTotalWords(), 0, "Total word count does not match the expected value.");
        });

        it("Check if the total byte count calculated is correct", () => {
            assert.strictEqual(fileMetrics.GetTotalBytes(), 0, "Total byte count does not match the expected value.");
        });

        it("Check if the total character count calculated is correct", () => {
            assert.strictEqual(fileMetrics.GetTotalCharacters(), 0, "Total character count does not match the expected value.");
        });
    });
});