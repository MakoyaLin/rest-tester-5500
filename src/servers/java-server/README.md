# JAVA API Server - Setup Instructions

## File Extension

- Rust source code files use the `.java` extension.
- Save the provided server code in a package called  `src/main/java/com/server/javaserver`.

## Installation Instructions

### Install Maven on Windows

- Download maven and jdk1.8 and configure the environment variables

### Install Maven on macOS

- Download maven and jdk1.8 and configure the environment variables

## Running the Server

1. Clone this repository to your local machine or create a new directory for the server code.

   ```sh
   cd src/servers/java-server
   ```

1. Run the server using maven:

   ```sh
   mvn clean package
   java -jar target/java-server.jar
   ```

1. The server will start on port `5003`. You can verify that the server is running by accessing the following URL in your browser or using a tool like `curl`:

   ```
   http://localhost:5003/users
   ```

## Endpoints

The server exposes the following REST API endpoints:

1. **GET /users** - Retrieve all users.
2. **GET /users/{id}** - Retrieve a user by ID.
3. **POST /users** - Add a new user.
4. **PUT /users/{id}** - Update a user by ID.
5. **PATCH /users/{id}** - Update the hours worked for a user by adding a value.
6. **DELETE /users/{id}** - Delete a user by ID.
7. **DELETE /users** - Delete all users.
