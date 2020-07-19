# IOLine

Provides functionality to access IO devices by reading line-by-line

## Usage

```go
// Reads line 3 in file `/path`
ioline.ReadFile("/path", 3)

// Reads line 5 in file `/path`, as well as 2 lines before line 3 and 1 line after line 3
line, _ := ioline.ReadFileWithBuffers("/path", 5, 2, 1)
line.FirstLine // 3
line.Before // Lines before line 5 up to the requested amount
line.After // Lines after line 5 up to the requested amount
line.Exact // Line 5 content
line.Count() // Total line read

// Reads a buffer line-by-line, without end-of-line character
scanner := ioline.NewScanner(reader)

// Doesn't iterate in case of errors or when finished
for scanner.Next() {
  line, err := scanner.Get()
  fmt.Println(line)
}

scanner.Error() // Last error
scanner.HasError() // True if an error happened

```

## TODO

- [ ] Remove `err` from `Get` call
