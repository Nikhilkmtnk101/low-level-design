# Simple File System Design using Composite Design Pattern

## Problem Statement

Imagine you're tasked with building a computer file system that aims to be organized and user-friendly. 
The goal is to create a system that can store both individual files and folders. 
Users should be able to easily inspect the contents of a folder using a simple command like `ls`. 
To achieve this, you will implement the file system using the Composite Design Pattern.

### Example Usage

Consider a scenario where there's a folder named "Documents" containing two files: "Report.doc" and "Picture.jpg."
When a user applies the `ls` command to the "Documents" folder, the output should be something like:
```
Documents:
- Report.doc
- Picture.jpg
