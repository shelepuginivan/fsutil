# `fsutil`

[![Go Reference](https://pkg.go.dev/badge/github.com/shelepuginivan/fsutil.svg)](https://pkg.go.dev/github.com/shelepuginivan/fsutil)
[![Go Report Card](https://goreportcard.com/badge/github.com/shelepuginivan/fsutil)](https://goreportcard.com/report/github.com/shelepuginivan/fsutil)
[![Actions Status](https://github.com/shelepuginivan/fsutil/workflows/Test%20package/badge.svg)](https://github.com/shelepuginivan/fsutil/actions)
[![codecov](https://codecov.io/github/shelepuginivan/fsutil/graph/badge.svg)](https://codecov.io/github/shelepuginivan/fsutil)
[![License: MIT](https://img.shields.io/badge/License-MIT-00cc00.svg)](https://github.com/shelepuginivan/fsutil/blob/main/LICENSE.md)

Package fsutil provides a collection of utility functions that extend the
capabilities of the Go standard filesystem package. It offers a variety of
tools for working with files and directories, making it easier to perform
common filesystem operations.

## Installation

```shell
go get github.com/shelepuginivan/fsutil
```

## Documentation

Is available on [pkg.go.dev](https://pkg.go.dev/github.com/shelepuginivan/fsutil).

## Details about `IsBinary`, `IsBinarySlow` and `IsBinaryFast`

> [!NOTE]
> Available since `v0.2.0`.

Each `IsBinary`, `IsBinarySlow` and `IsBinaryFast` report whether data is
binary, i.e. a non human-readable text, such as image, audio or video stream.

- `IsBinarySlow` has the **best accuracy** and the **worst speed**, since it
  checks the entire byte sequence.
- `IsBinaryFast` has the **worst accuracy** and the **best speed** since it
  checks only first line of the sequence.
- `IsBinary` is in the middle &mdash; it checks the first and every 2^N line
  (i.e. 0, 1, 2, 4, 8, 16, etc.), providing a good accuracy while keeping good
  performance.

The table below shows the accuracy of these functions, determined by testing
against random byte sequences:

|    Function    |   Mimimal accuracy   |   Actual accuracy    |
| :------------: | :------------------: | :------------------: |
| `IsBinary`     |  990/1000 (`99.0%`)  | 1000/1000 (`100.0%`) |
| `IsBinarySlow` | 1000/1000 (`100.0%`) | 1000/1000 (`100.0%`) |
| `IsBinaryFast` |  950/1000 (`95.0%`)  |  991/1000 (`99.1%`)  |
