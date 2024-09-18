# Compiled vs Interpreted Languages

Understanding the difference between compiled and interpreted programming languages is crucial for software development and deployment. Below we outline the key differences, provide examples of each type, and explain why using a compiled language for projects like Textio can be beneficial.

![alt text](/typeSize/7RBQRNA.png)

## Compiled Languages

Compiled languages convert source code into machine code using a compiler, generating an executable program. This means that once the program is compiled, it can be run independently of the original source code or the compiler.

**Advantages:**

- Easier distribution: Users can run the compiled program without needing the source code or a compiler installed.
- Generally faster execution since the code is already translated to machine language.

**Examples of Compiled Languages:**

- Go
- C
- C++
- Rust

## Interpreted Languages

Interpreted languages execute code line-by-line at runtime using an interpreter. This can make them more flexible and easier to debug, but it also means that the end user must have the interpreter installed on their machine in order to run the program.

**Disadvantages:**

- Distribution can be cumbersome: Users need to have the appropriate interpreter installed and access to the source code.
- Generally slower execution compared to compiled languages.

**Examples of Interpreted Languages:**

- JavaScript (sometimes JIT-compiled)
- Python
- Ruby

## Why Build Textio in a Compiled Language?

Building Textio in a compiled language such as Go provides several advantages:

- **No Runtime Dependencies:** When deploying our server, we do not need to include any runtime language dependencies like Node.js or a Python interpreter. This simplifies the deployment process considerably.
- **Simple Deployment:** We can just add the pre-compiled binary to the server and start it up, streamlining the operational setup.

Understanding these differences is key for developers in choosing the right tools and languages for their projects, especially when considering scalability and ease of use for end users.
