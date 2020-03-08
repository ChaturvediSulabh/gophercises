# A file renaming tool, works recursively in a given Directory

## Given a filename, it standardize the fileName

- File Must begin with Capital Letter and rest all the characters in a file name must be of small letters.
- File without any extension should be appended with `.txt` file extension
- Any special character other that an underscore must be replaced by an underscore
- All Files should be appended by `_<LOCAL_TIMESTAMP_WHEN_FILE_WAS_MODIFIED>`

## Usage

```bash
➜ ./bin/doProperFileNames --help
Usage of ./bin/doProperFileNames:
  -pathToDir string
        full path to a directory (default ".")
```
