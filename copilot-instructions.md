# Copilot Instructions for arc-switch Go Projects

This file provides guidelines for GitHub Copilot and other AI assistants when helping with this repository.

## Golang Best Practices

### Code Style and Organization

- Follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments) for style guidance
- Respect Go's [Effective Go](https://golang.org/doc/effective_go) principles
- Use `gofmt` or `goimports` for consistent formatting
- Organize imports in three groups: standard library, external packages, internal packages
- Follow package structure: `package main` for executables, descriptive names for libraries
- Place interfaces near where they're used, not in separate files

### Error Handling

- Return errors explicitly; avoid panic except for unrecoverable situations
- Handle errors immediately; don't defer checks
- Use custom error types for specific error conditions
- Use `errors.Is()` and `errors.As()` for error checking when appropriate
- Prefer `fmt.Errorf("...%w", err)` for wrapping errors

### Variable Naming

- Use short, concise variable names in small scopes
- Use descriptive names for exported functions, types, and variables
- Use camelCase for private and PascalCase for exported identifiers
- Prefer single-letter variables for indices and short loops

### Testing

- Write table-driven tests using `testing.T`
- Use `t.Parallel()` for tests that can run in parallel
- Test both expected successes and failures
- Avoid global variables in tests
- Name test files as `*_test.go`

### Performance

- Use buffered I/O for file operations
- Preallocate slices when size is known
- Use efficient data structures (maps for lookups, slices for iteration)
- Consider sync.Pool for frequently allocated objects
- Profile before optimizing

### Concurrency

- Use channels to communicate between goroutines, not shared memory
- Use context for cancellation and timing
- Be careful with goroutine leaks
- Use sync.Mutex for simple state protection
- Consider sync.WaitGroup for waiting on multiple goroutines

### SNMP Project Specifics

- Log errors with context (device info, OIDs)
- Handle SNMP timeouts gracefully
- Use JSON struct tags consistently
- Document all configuration options
- Include retry logic for transient failures
- Validate input configurations
- Follow separation of concerns between:
  - SNMP operations
  - MIB handling
  - Configuration management
  - Output formatting

### JSON Output Format

All tools must output data in JSON format with standardized fields for consistency and interoperability.

#### Common Required Fields

Every tool output must include these three required fields:

- `data_type`: String identifying the type of data (e.g., "cisco_nexus_mac_table", "interface_status", "lldp_neighbors")
- `timestamp`: Full timestamp with date and time (format: "MM/DD/YYYY HH:MM:SS AM/PM")
- `date`: Date in MM/DD/YYYY format

#### Tool-Specific Fields

Beyond the common required fields, each tool will have its own unique fields relevant to the data being processed.

#### Example: Cisco Nexus MAC Table Parser

```json
{
  "data_type": "cisco_nexus_mac_table",
  "timestamp": "06/23/2025 05:05:01 PM",
  "date": "06/23/2025",
  "primary_entry": true,
  "gateway_mac": false,
  "routed_mac": false,
  "overlay_mac": false,
  "vlan": "7",
  "mac_address": "02ec.a004.0000",
  "type": "dynamic",
  "age": "NA",
  "secure": "F",
  "ntfy": "F",
  "port": "Eth1/1",
  "vpc_peer_link": false
}
```

**Tool-Specific Fields for MAC Table Parser:**
- `primary_entry`: Boolean indicating if this is a primary entry
- `gateway_mac`: Boolean indicating if this is a gateway MAC
- `routed_mac`: Boolean indicating if this is a routed MAC
- `overlay_mac`: Boolean indicating if this is an overlay MAC
- `vlan`: VLAN ID as string
- `mac_address`: MAC address in Cisco format (xxxx.xxxx.xxxx)
- `type`: Entry type (e.g., "dynamic", "static")
- `age`: Age information or "NA"
- `secure`: Security flag ("F" or "T")
- `ntfy`: Notification flag ("F" or "T")
- `port`: Port identifier
- `vpc_peer_link`: Boolean, only present when true for vPC peer-link entries

### Project Structure

- Keep `main.go` small and focused on setup and coordination
- Use separate packages for specialized functionality
- Place most logic in importable packages for testing
- Use appropriate directory structure:
  - `/cmd` - Main applications
  - `/pkg` - Library code
  - `/internal` - Private application code
  - `/api` - API definitions
  - `/configs` - Configuration file templates

### Dependencies

- Minimize external dependencies
- Use Go modules for dependency management
- Vendor dependencies for reproducible builds
- Use specific versions in go.mod, not master branch
