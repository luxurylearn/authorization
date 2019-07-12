# Go based user service
Service is responsible for registration, authentication and authorization users in the system.

All admin tools with managing users purposes should use endpoints exposed by this service to extract all required data.

# Dependencies
Service is using docker-compose as a toll for managing external dependencies.

External dependencies:
- PostgreSQL (under development)
- RabbitMQ (under development)
- OAuth2 client (under development)

# Testing
To be able to run unit/integration tests locally it is needed to run docker-compose command:
`docker-compose up -d test` (under development) 

# Project structure
For additional information and/or directories in project see [this](https://github.com/golang-standards/project-layout) resource

# JetBrains WSL Terminal Setup
To be able to use wsl right from JetBrains tools terminal make sure you have configured linux wsl first.

Then set the terminal path in JetBrains `settings>tools>terminal` to
     
     C:\Windows\System32\wsl.exe
     
Now open the terminal. It should correctly move the current directory to the `/mnt/your/path`
