[![Build](https://github.com/nao1215/apk-parser/actions/workflows/build.yml/badge.svg)](https://github.com/nao1215/apk-parser/actions/workflows/build.yml)
[![PlatformTests](https://github.com/nao1215/apk-parser/actions/workflows/platform_test.yml/badge.svg)](https://github.com/nao1215/apk-parser/actions/workflows/platform_test.yml)
[![reviewdog](https://github.com/nao1215/apk-parser/actions/workflows/reviewdog.yml/badge.svg)](https://github.com/nao1215/apk-parser/actions/workflows/reviewdog.yml)
[![codecov](https://codecov.io/gh/nao1215/apk-parser/branch/main/graph/badge.svg?token=DNV3TRMRCJ)](https://codecov.io/gh/nao1215/apk-parser)
[![Go Reference](https://pkg.go.dev/badge/github.com/nao1215/apk-parser.svg)](https://pkg.go.dev/github.com/nao1215/apk-parser)
[![Go Report Card](https://goreportcard.com/badge/github.com/nao1215/apk-parser)](https://goreportcard.com/report/github.com/nao1215/apk-parser)
![GitHub](https://img.shields.io/github/license/nao1215/apk-parser)  
# deapk - parse android package (.apk), getting meta data.
The deapk (decompile android package) command parses the apk file and outputs metadata information. It is still in the development stage and output information is few. In the future, deapk will provide the ability to decompile dex files and convert them to source code.
  
# How to install
### Step1. Install golang
deapk command only supports installation with `$ go install`. If you does not have the golang development environment installed on your system, please install golang from the [golang official website](https://go.dev/doc/install).

### Step2. Install deapk
```
$ go install github.com/nao1215/deapk@latest
```

# How to use
## Output *.apk metadata
```
$ deapk info testdata/app-debug.apk 
pacakage name      : jp.debimate.deapk_test
application name   : deapk-test
application version: 1.0
sdk target version : 31
sdk max version    : -1 (deprecated attribute)
sdk min version    : 31
main activity      : jp.debimate.deapk_test.MainActivity
```

## Output *.apk metadata in json format
```
$ deapk info --json testdata/app-debug.apk 
{
        "Basic": {
                "package_name": "jp.debimate.deapk_test",
                "application_name": "deapk-test",
                "version": "1.0",
                "main_activity": "jp.debimate.deapk_test.MainActivity",
                "sdk": {
                        "minimum": 31,
                        "target": 31,
                        "maximum": -1
                }
        }
}
```

## Output *.apk metadata to file
### Use redirect
```
$ deapk info --json testdata/app-debug.apk > apk.json
```
### Use --output option
```
$ deapk info --json --output=apk.json testdata/app-debug.apk
$ cat apk.json 
{
        "Basic": {
                "package_name": "jp.debimate.deapk_test",
                "application_name": "deapk-test",
                "version": "1.0",
                "main_activity": "jp.debimate.deapk_test.MainActivity",
                "sdk": {
                        "minimum": 31,
                        "target": 31,
                        "maximum": -1
                }
        }
}
```
# Contributing
First off, thanks for taking the time to contribute! ❤️  See [CONTRIBUTING.md](./CONTRIBUTING.md) for more information.
Contributions are not only related to development. For example, GitHub Star motivates me to develop!

# Contact
If you would like to send comments such as "find a bug" or "request for additional features" to the developer, please use one of the following contacts.

- [GitHub Issue](https://github.com/nao1215/deapk/issues)

# LICENSE
The deapk project is licensed under the terms of [MIT License](./LICENSE).

