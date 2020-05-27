# DI Example

This is a simple example of using Dependency Injection in Go. I made use of
Uber's `fx` library. To run:

```bash
go run main.go
```

## Branches

There are two other branches here that should help illustrate why you'd want to
use dependency injection at all.

1. [No DI](https://github.com/CrowderSoup/di-example/tree/no-di): an example
   without dependency injection.
2. [Manual DI](https://github.com/CrowderSoup/di-example/tree/manual-di): an
   example with dependency injection done manually. It's better, but not as
   clean, and managing application lifecycle is more difficult (think about how
   you'd do a graceful shutdown of the HTTP server, and you have to have
   `logger.Sync()` in your main)
3. [DI with FX](https://github.com/CrowderSoup/di-example/tree/master): shows
   how you'd do dependency injection with fx. This is just the `master` branch.
