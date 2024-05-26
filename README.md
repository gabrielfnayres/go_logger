"""

# Reverse-logger GO

The Reverse Logger is a powerful tool that combines a reverse shell with keylogger functionalities. It allows users to remotely access systems and log all typing activities in real-time, enhancing network security and enabling forensic investigations.

## Features

- Remote Access
- Keylogging: Captures and logs all typing activities in real-time.
- Simple reverse shell

## Getting Started

### Prerequisites

- Go 1.22.1 or later

### Installation

1. Clone the repository:

    ```sh
    git clone <repository-url>
    cd go_logger
    ```

2. Install dependencies:

    ```sh
    go mod download
    ```

## Contribution Guidelines

To contribute to this project, please follow these steps:

1. Clone the Repository: Clone the repository to your local machine.

    ```sh
    git clone <repository-url>
    ```

2. Create a New Branch: Create a new branch for your feature or bugfix.

    ```sh
    git checkout -b feature-branch
    ```

3. Make Your Changes: Make your changes to the codebase.

4. Commit Your Changes: Commit your changes with a descriptive message.

    ```sh
    git commit -m "refactor: clean up Environment class and improve parameter handling"
    ```

5. Push Your Changes: Push your changes to the remote repository.

    ```sh
    git push origin feature-branch
    ```

6. Create a Pull Request: Create a pull request from your branch to the main branch, and provide a detailed description
   of your changes.

## Commit Naming Conventions

For consistency, please use the following commit message format:

- `feat`: A new feature
- `fix`: A bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, missing semicolons, etc.)
- `refactor`: Refactoring code without changing functionality
- `test`: Adding or refactoring tests
- `chore`: Other changes that don't modify src or test files

Example:

```sh
    git commit -m "refactor: clean up Environment class and improve parameter handling"
```