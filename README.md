# Toolstub

Generate executable alias for Go build-time dependencies / third-party `go tool`. Inspired by Ruby Bundler's `bundle binstubs`.

## Example

```shellsession
$ toolstub github.com/golangci/golangci-lint/cmd/golangci-lint
```

And the generated output will be [a bash script wrapping `go` commands](bin/golangci-lint).
