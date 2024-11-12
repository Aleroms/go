# Code Coverage in Go

**Code coverage** measures the percentage of code executed during tests, helping to identify untested parts of your codebase. Go’s `go test` tool includes built-in support for generating coverage reports.

---

### Key Points:

- **Coverage Reports**: Running tests with the `-cover` flag generates a basic coverage report.
- **Detailed Reports**: Use `-coverprofile` to create a detailed coverage report file.
- **Visual Coverage**: You can view coverage in a web browser for a more visual report using `go tool cover`.

---

### Commands for Code Coverage

1. **Basic Coverage**:

   ```bash
   go test -cover
   ```

   This command shows the overall coverage percentage of your tests.

2. **Detailed Coverage Report**:

   ```bash
   go test -coverprofile=coverage.out
   ```

   Generates a detailed report saved to `coverage.out`.

3. **HTML Report**:
   ```bash
   go tool cover -html=coverage.out
   ```
   Opens a web-based HTML report, highlighting covered and uncovered lines of code for easy analysis.

---

### Example Usage:

```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Summary

- Use **`-cover`** to check coverage percentage.
- Generate a **coverage profile** with `-coverprofile` and view it with `go tool cover` for a visual report.
- Coverage reports help ensure your tests cover critical paths in your code.
