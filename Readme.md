[![Build Status](https://travis-ci.com/woocart/gsutil.svg?branch=master)](https://travis-ci.com/woocart/gsutil)

# gscp

gscp is a go application that lets you upload, download objects from Google Cloud Storage.

## Installation

Export `GOOGLE_APPLICATION_CREDENTIALS` to enviroment with path to your authentication file. You can follow https://cloud.google.com/docs/authentication/getting-started#setting_the_environment_variable for more detailed instructions.

Download gscp from `Releases`

## Examples

```shell
echo "test" | gscp stdio gs://bucket/test/files   # read from stdin
echo "test" | gscp stdio gs://bucket/test/files key=value # set custom metadata
gscp gs://bucket/test/files stdio  # read to stdout
```

```shell
usage: gscp [<flags>] <from> <to> [<metadata>...]

Copies data from and to Google Cloud Storage

Flags:
  --help     Show context-sensitive help (also try --help-long and --help-man).
  --version  Show application version.

Args:
  <from>        where to read from: gs://bucketname/path or - from stdin or /path/ for local file
  <to>          Where to write to: gs://bucketname/path or - to stdout or /path/ for local file
  [<metadata>]  KV pairs to append to uploaded object
```
