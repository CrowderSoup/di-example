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

## Dependency Injection

### What is it?

Dependency Injection is a design pattern whereby an object/module's dependencies
are provided to it. It is one of several techniques for implementing [Inversion
of Control](https://en.wikipedia.org/wiki/Inversion_of_control).

So rather than:

```go 
func NewClient() *Client {
    return &Client{
        Service: service.NewService(),
    }
}
```

You would have:

```go 
func NewClient(service *service.Service) *Client {
    return &Client{
        Service: service,
    }
}
```

### Why tho??

If you wanted to write a test for the `Client` first example above, you
wouldn't be able to mock it's `Service`. In the second example it would be
trivial to inject a mock `Service` and write tests for your `Client`.

From [Wikipedia](https://en.wikipedia.org/wiki/Dependency_injection#Advantages):

- Dependency injection allows a client the flexibility of being configurable. 
  Only the client's behavior is fixed. The client may act on anything that 
  supports the intrinsic interface the client expects.
- Dependency injection can be used to externalize a system's configuration 
  details into configuration files, allowing the system to be reconfigured 
  without recompilation. Separate configurations can be written for different 
  situations that require different implementations of components. This 
  includes, but is not limited to, testing.
- Because dependency injection does not require any change in code behavior it 
  can be applied to legacy code as a refactoring. The result is clients that are 
  more independent and that are easier to unit test in isolation using stubs or 
  mock objects that simulate other objects not under test. This ease of testing 
  is often the first benefit noticed when using dependency injection.
- Dependency injection allows a client to remove all knowledge of a concrete 
  implementation that it needs to use. This helps isolate the client from the 
  impact of design changes and defects. It promotes reusability, testability and 
  maintainability.
- Reduction of boilerplate code in the application objects, since all work to 
  initialize or set up dependencies is handled by a provider component.
- Dependency injection allows concurrent or independent development. Two 
  developers can independently develop classes that use each other, while only 
  needing to know the interface the classes will communicate through. Plugins 
  are often developed by third party shops that never even talk to the developers 
  who created the product that uses the plugins.
- Dependency Injection decreases coupling between a class and its dependency.

There are also some [disadvantages](https://en.wikipedia.org/wiki/Dependency_injection#Disadvantages)
(makes code harder to read, breaks "jump to definition", etc.) but I feel that 
the advantages far outweigh the disadvantages. But that's just like, my opinion 
man.

![Opinion](./opinion.jpeg)

### Why `fx`??

You can do dependency injection manually in a small app. But as your app grows
manually wiring up dependencies becomes tedious and cumbersome. A dependency
injection framework (like `fx`) handles the wiring for you. This saves you from
having overly complex `main` functions to handle all the wiring.

With microservices you also have lots of infrastructure components (metrics, 
logging, tracing, etc.). Without DI that boilerplate must be copied from repo to 
repo, or put in common components. But even with common components you still 
have to handle initialization and wiring of these things yourself.

With a DI framework you can simply import the common utilities and tell your
framework how / where to new them up. Then they are provided to whatever parts
of your application need them.
