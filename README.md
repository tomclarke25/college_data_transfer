# College Data Transfer

## Prerequisites

- Go 1.16 or later
- Git

## Setup

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/collegeDataTransfer.git
    cd collegeDataTransfer
    ```

2. Install dependencies (if any):

    ```sh
    go mod tidy
    ```

## Running the Project

1. Ensure your CSV files are located in the `data` directory:
    - `data/Teachers.csv`
    - `data/Students.csv`
    - `data/Classes.csv`

2. Run the main program:

    ```sh
    go run main.go
    ```

This will generate the following JSON files in the project directory:
- `teachers.json`
- `students.json`
- `classes.json`

## Project Structure

- `main.go`: The main entry point of the application.
- `data/`: Directory containing the input CSV files.
