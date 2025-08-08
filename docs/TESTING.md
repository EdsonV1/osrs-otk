# Testing Guide

This document describes the testing strategy and setup for the OSRS OTK project.

## Testing Structure

### Backend Tests (Go)

#### Unit Tests
Located in `*_test.go` files alongside source code:

- **Calculator Tests**: Comprehensive tests for all calculators
  - `internal/calculators/technique/gotr/gotr_test.go` - GOTR calculator tests
  - `internal/calculators/technique/ardy_knights/ardy_knights_test.go` - Ardy Knights tests
  - `internal/calculators/technique/wintertodt/wintertodt_test.go` - Wintertodt tests
  - `internal/calculators/technique/birdhouses/birdhouse_test.go` - Birdhouse tests
  - `internal/calculators/xp_table_test.go` - XP table functionality tests

- **Utility Tests**: Test helpers and utilities
  - `internal/calculators/test_utils.go` - Common testing utilities

#### Integration Tests
Located in `internal/integration/`:

- **API Integration Tests**: End-to-end HTTP API testing
- **CORS Testing**: Cross-origin request handling
- **Concurrent Load Testing**: Performance under concurrent requests

### Frontend Tests (TypeScript/Svelte)

Located in `web/frontend/`:

- **Component Tests**: Individual Svelte component testing
- **Type Checking**: TypeScript compilation and type safety
- **Linting**: Code style and quality checks

## Running Tests

### Quick Commands

```bash
# Run all tests
make test-all

# Run only backend tests
make test

# Run tests with coverage
make test-coverage

# Run integration tests
make test-integration

# Run benchmarks
make benchmark

# Run with race detection
make test-race
```

### Detailed Commands

#### Backend Tests

```bash
# Run all Go tests with verbose output
go test -v ./...

# Run tests with coverage report
go test -v -race -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# Run specific test package
go test -v ./internal/calculators/technique/gotr/

# Run specific test function
go test -v -run TestCalculateGOTRData ./internal/calculators/technique/gotr/

# Run benchmarks
go test -bench=. -benchmem ./...

# Run tests with race detection
go test -v -race ./...
```

#### Frontend Tests

```bash
cd web/frontend

# Run type checking
npm run check

# Run linting
npm run lint

# Build to check for errors
npm run build
```

## Test Coverage

### Current Coverage Targets

- **Overall**: Aim for >80% code coverage
- **Critical Paths**: 100% coverage for calculator logic
- **API Handlers**: 100% coverage for HTTP endpoints
- **Error Handling**: 100% coverage for validation and error cases

### Viewing Coverage

After running `make test-coverage`, open `coverage.html` in a browser to see detailed coverage reports.

## Test Categories

### 1. Calculator Tests

Each calculator has comprehensive test coverage including:

- **Valid Input Tests**: Testing expected functionality
- **Boundary Tests**: Edge cases and limits
- **Invalid Input Tests**: Error handling and validation
- **Performance Tests**: Benchmarks for critical functions
- **Consistency Tests**: Data integrity and mathematical correctness

### 2. API Tests

- **Request/Response Validation**: JSON structure and data types
- **HTTP Status Codes**: Proper error codes for different scenarios
- **CORS Headers**: Cross-origin request handling
- **Concurrent Requests**: Performance under load

### 3. Integration Tests

- **End-to-End Workflows**: Complete user journeys
- **Database Integration**: Data persistence and retrieval
- **External Dependencies**: Third-party service mocking
- **Environment Configuration**: Different deployment scenarios

## CI/CD Integration

### GitHub Actions Pipeline

Located in `.github/workflows/ci.yml`:

1. **Test Backend**: Go tests, linting, and coverage
2. **Test Frontend**: TypeScript checking, linting, and build
3. **Security Scan**: Vulnerability scanning with Gosec
4. **Build Docker**: Multi-platform container builds
5. **Deploy**: Automated deployment to staging/production

### Pre-commit Hooks

Install pre-commit hooks to run tests before commits:

```bash
# Install pre-commit (if not already installed)
pip install pre-commit

# Install hooks
pre-commit install

# Run manually
pre-commit run --all-files
```

Configuration in `.pre-commit-config.yaml` includes:
- Go formatting and linting
- Test execution
- Docker linting
- YAML/JSON validation
- Markdown linting

## Writing New Tests

### Test Structure

Follow Go testing conventions:

```go
func TestFunctionName(t *testing.T) {
    tests := []struct {
        name        string
        input       InputType
        expectError bool
        expected    ExpectedType
    }{
        {
            name:        "descriptive test name",
            input:       /* test input */,
            expectError: false,
            expected:    /* expected output */,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := FunctionToTest(tt.input)
            
            if tt.expectError {
                if err == nil {
                    t.Errorf("Expected error but got none")
                }
                return
            }
            
            if err != nil {
                t.Errorf("Unexpected error: %v", err)
                return
            }
            
            // Add specific assertions here
        })
    }
}
```

### Test Guidelines

1. **Descriptive Names**: Use clear, descriptive test names
2. **Edge Cases**: Test boundary conditions and error cases
3. **Table-Driven Tests**: Use table-driven tests for multiple scenarios
4. **Setup/Teardown**: Use test helpers for common setup
5. **Assertions**: Use specific, meaningful assertions
6. **Documentation**: Document complex test scenarios

### Using Test Utilities

The project includes test utilities in `internal/calculators/test_utils.go`:

```go
func TestExample(t *testing.T) {
    helper := calculators.NewTestHelper(t)
    
    result := SomeCalculation()
    
    helper.AssertPositive(result.Value, "calculation result")
    helper.AssertInRange(result.Rate, 0.0, 1.0, "success rate")
}
```

## Performance Testing

### Benchmarks

Run benchmarks to ensure performance:

```bash
# Run all benchmarks
make benchmark

# Run specific benchmark
go test -bench=BenchmarkCalculateGOTRData ./internal/calculators/technique/gotr/

# Profile memory usage
go test -bench=. -benchmem -memprofile=mem.prof ./...
```

### Load Testing

Use the integration tests for basic load testing:

```bash
# Run concurrent request tests
go test -v ./internal/integration/ -run TestConcurrentRequests
```

## Troubleshooting

### Common Issues

1. **Test Failures**: Check error messages and expected vs actual values
2. **Coverage Issues**: Identify untested code paths in coverage report
3. **Race Conditions**: Use `-race` flag to detect concurrent access issues
4. **Timeout Issues**: Adjust test timeouts for slow operations

### Debug Commands

```bash
# Run specific failing test
go test -v -run SpecificTestName ./path/to/package/

# Run with debugging output
go test -v -args -debug ./...

# Check for race conditions
go test -race ./...
```

## Best Practices

1. **Test Independence**: Tests should not depend on each other
2. **Deterministic**: Tests should produce consistent results
3. **Fast Execution**: Keep tests fast to encourage frequent running
4. **Clear Failures**: Test failures should clearly indicate the problem
5. **Maintainable**: Tests should be easy to understand and modify
6. **Comprehensive**: Cover happy paths, edge cases, and error conditions