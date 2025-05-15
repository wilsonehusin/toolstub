# Toolstub

Generate executable alias for Go build-time dependencies / third-party `go tool`. Inspired by Ruby Bundler's `bundle binstubs`.

## Install

```shellsession
# go run go.husin.dev/toolstub@latest -tool go.husin.dev/toolstub
```

## Usage

```shellsession
# toolstub -tool github.com/golangci/golangci-lint/cmd/golangci-lint
```

And the generated output will be [a bash script in `bin/golangci-lint` wrapping various `go` commands](bin/golangci-lint).

Running the program is as boring as executing `bin/golangci-lint`. First run will be slower due to installation and subsequent runs will be instant.

The versioning and dependencies are tracked in `_tools/golangci-lint.go.{mod,sum}`, which can be tracked by dependency managers like Dependabot.
