# AGENTS.md - Guide for AI Coding Agents

## Project Overview

This is a **personal learning & reference repository** structured as a Go module (`module note`), not a production application. It contains documentation, code examples, and test cases across multiple technology domains.

**Key Characteristics:**
- Primary content: Markdown files with code examples and explanations (mixed Chinese/English)
- Code examples: Sparse test files (`*_test.go`) demonstrating concepts
- Purpose: Learning resource, algorithm practice, and technical reference
- Not a service/API - no main applications to build or deploy

## Directory Structure & Content Mapping

```
golnag/          → Go language learning (typo in folder name - intentional)
├── 基础篇/        → Fundamentals: goroutine, time, encryption patterns
├── 性能篇/        → Performance: benchmarking, fmt analysis
leetcode/        → LeetCode problems & solutions
bookmarks/       → Curated links (AI, skills, web resources)
prompt/          → Prompt templates & AI conventions (go_rules.md contains architectural patterns)
skill/           → Skill-building exercises (HashSet implementation)
php/mysql/nginx/ → Web tech documentation & examples
linux/           → System administration notes
```

## Go Code Patterns & Conventions

**Testing & Benchmarking:**
- Test files use standard Go testing (`testing` package with `*_test.go` suffix)
- Benchmarking used for performance analysis (see `性能篇/fmt性能分析_test.go`)
- RSA encryption example in `基础篇/加解密/rsa_test.go` demonstrates crypto concepts
- Dependencies: `github.com/stretchr/testify` for assertions

**Go Development Practices:**
- Run tests: `go test ./...` from project root
- Benchmark specific tests: `go test -bench=. ./golnag/性能篇`
- Chinese directory names are intentional (educational context for Chinese learners)
- Code examples often show multiple solution approaches with trade-offs explained

**Common Goroutine Pattern:**
When teaching async patterns, repository demonstrates the closure variable capture issue:
```go
// ❌ Wrong: captures reference to i
for i := 1; i <= 3; i++ {
  go func() { fmt.Println(i) }()
}

// ✅ Right: pass i as parameter
for i := 1; i <= 3; i++ {
  go func(temp int) { fmt.Println(temp) }(i)
}
```
(See `基础篇/goroutine/在for循环中goroutine调用循环变量问题.md`)

## AI Code Generation Rules

Reference file: `prompt/go_rules.md` contains architectural principles:

- **Architecture**: Clean Architecture preferred (handler → service → repository → domain)
- **Dependencies**: Dependency injection, interface-driven development
- **Error Handling**: Explicit error checking, use `fmt.Errorf("context: %w", err)` for wrapping
- **Concurrency**: Use `context.Context` for cancellation; protect shared state with channels/sync
- **Naming**: PascalCase for exports, camelCase for private, package names lowercase
- **Principle**: Code readability first, minimize framework dependencies

## Documentation Style

- Mix of Chinese technical explanations and English code examples
- Each learning topic gets a separate Markdown file with:
  - Problem statement/context
  - Code examples (multiple approaches if applicable)
  - Explanation of pitfalls and best practices
  - Links to GitHub references where applicable
- Images stored in `pics/` directory for reference

## Developer Workflow

**Commands:**
```bash
go test ./...              # Run all tests
go test -v ./...           # Verbose test output
go test -bench=. ./golnag/性能篇  # Benchmarks only
```

**No special build/deploy process** - this is a reference repository. Focus on:
1. Verifying test files execute correctly (`go test` passes)
2. Code examples are syntactically valid and runnable
3. Documentation is clear and examples are self-contained

## Integration Points

- **External testing framework**: `github.com/stretchr/testify` (for `assert` functions)
- **No external APIs or services** - examples are self-contained demonstrations
- **git workflow**: Standard Git; repository structured for reference/forking

## When Adding Content

- Keep Go examples focused and self-contained
- Place test files next to their implementation or in the same directory
- Use `*_test.go` naming and proper `package` declarations
- Document "gotchas" prominently (like the goroutine closure pattern)
- Add comments explaining the "why" not just the "what" for educational value
- Reference related concepts in the repository when applicable

