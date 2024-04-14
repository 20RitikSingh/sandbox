package main

import (
	"fmt"
	"os"
	"os/exec"
)

// CompileCPP compiles the provided C++ source code using GCC inside a Docker container
func CompileCPP(code string) error {
	// Write the C++ code to a temporary file
	tmpfile, err := os.CreateTemp("", "cppcode_*.cpp")
	if err != nil {
		fmt.Println("Error creating temporary file:", err)
		return err
	}
	defer os.Remove(tmpfile.Name()) // Clean up temporary file
	if _, err := tmpfile.WriteString(code); err != nil {
		fmt.Println("Error writing to temporary file:", err)
		return err
	}

	// Create a Dockerfile with the necessary instructions
	dockerfile := []byte(fmt.Sprintf(`
FROM gcc:latest
WORKDIR /usr/src/app
COPY %s /usr/src/app/main.cpp
RUN g++ -o my_program main.cpp
`, tmpfile.Name()))

	// Write the Dockerfile to a temporary file
	tmpDockerfile, err := os.CreateTemp("", "Dockerfile_*.txt")
	if err != nil {
		fmt.Println("Error creating temporary Dockerfile:", err)
		return err
	}
	defer os.Remove(tmpDockerfile.Name()) // Clean up temporary file
	if _, err := tmpDockerfile.Write(dockerfile); err != nil {
		fmt.Println("Error writing to temporary Dockerfile:", err)
		return err
	}

	// Build the Docker image
	buildCmd := exec.Command("docker", "build", "-t", "my_cpp_compiler", "-f", tmpDockerfile.Name(), ".")
	if err := buildCmd.Run(); err != nil {
		fmt.Println("Error building Docker image:", err, buildCmd.Stdout, buildCmd.Stderr)
		return err
	}

	// Run the Docker container to compile the C++ code
	runCmd := exec.Command("docker", "run", "-v", fmt.Sprintf("%s:/usr/src/app/main.cpp", tmpfile.Name()), "my_cpp_compiler")
	if err := runCmd.Run(); err != nil {
		fmt.Println("Error running Docker container:", err, runCmd.Stdout, runCmd.Stderr)
		return err
	}

	fmt.Println("Compilation successful. Executable file: my_program")

	return nil
}

func script() {
	// Example usage of CompileCPP function
	code := `
#include <iostream>
int main() {
    std::cout << "Hello, World!" << std::endl;
    return 0;
}`
	err := CompileCPP(code)
	if err != nil {
		fmt.Println("compile Error:", err)
	}
}
