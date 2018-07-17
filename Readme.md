[![Build Status](https://travis-ci.com/woocart/gsutil.svg?branch=master)](https://travis-ci.com/woocart/gsutil)

# gscp

gscp is a go application that lets you upload, download objects from Google Cloud Storage.

## Installation

Export `GOOGLE_APPLICATION_CREDENTIALS` to enviroment with path to your authentication file. You can follow https://cloud.google.com/docs/authentication/getting-started#setting_the_environment_variable for more detailed instructions.

Download gscp from `Releases`

## Examples

```shell
echo "test" | gscp stdio gs://bucket/test/files   # read from stdin
gscp gs://bucket/test/files stdio  # read to stdout
```