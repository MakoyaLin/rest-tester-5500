
# Lessons Learned

This document captures lessons learned from using different programming languages in our project.

## Language 1: Java

### What Did We Learn About This Language?
- Java is a highly structured, object-oriented language that enforces a clear separation of concerns through classes and packages.
- Using Spring Boot simplifies setting up a server, but it requires an understanding of various annotations and dependencies.
- Java requires more boilerplate code, making it somewhat verbose but very organized.

### Interesting Problems Encountered
- Managing dependencies with Maven was challenging at first due to its complex configuration, especially when adding new libraries.
- Exception handling in Java requires more code and planning, as all exceptions need to be explicitly handled, unlike languages with more permissive error handling.

## Language 2: Go

### What Did We Learn About This Language?
- Go is minimalistic and prioritizes simplicity, making it easy to read and understand with less boilerplate code.
- Go does not support classes, so it relies on structs and interfaces, encouraging a different way of organizing code.
- Built-in concurrency support with goroutines offers a powerful way to handle multiple requests simultaneously.

### Interesting Problems Encountered
- Error handling in Go is explicit, requiring checks for each error. This approach keeps the code predictable but can feel repetitive.
- Setting up HTTP routes with Gin required more manual configuration compared to Java’s Spring Boot, which is more automated.

## Language 3: TypeScript

### What Did We Learn About This Language?
- TypeScript, as a superset of JavaScript, provides type safety, which helps catch errors during development.
- Using Express.js made setting up the server straightforward, and TypeScript's interfaces helped define data structures clearly.
- TypeScript’s flexibility allows for functional or object-oriented approaches, making it adaptable to different coding styles.

### Interesting Problems Encountered
- Integrating TypeScript with Node.js required additional setup, such as configuring tsconfig and compiling TypeScript files to JavaScript.
- Type definitions occasionally caused conflicts with third-party libraries, requiring explicit type casting or custom type definitions.

## General Reflections

### Key Takeaways
- Each language offers unique strengths: Java for structure, Go for simplicity and performance, and TypeScript for flexibility and type safety.
- Working with multiple languages has improved our understanding of RESTful API concepts, as these remain consistent across languages.

### Future Improvements
- Experimenting with more languages like C# or Ruby could further broaden our understanding of backend frameworks.
- Focus on improving documentation and setup instructions for easier cross-platform compatibility.
- Consider using containerization (e.g., Docker) in future projects to standardize the environment across different languages.
