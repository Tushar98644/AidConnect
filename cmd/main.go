package main

import (
    "log"
    "os"
    "os/signal"
    "syscall"
    "github.com/Tushar98644/AidConnect/internal/server"
)

func main() {
    // Set up logging
    log.SetFlags(log.LstdFlags | log.Lshortfile)
    log.Println("Starting the server...")

    // Set up channel for capturing OS signals
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

    // Start the server
    server.Serve();

    // Wait for termination signal
    <-quit
    log.Println("Shutting down the server...")

    // Perform any necessary cleanup here
    // For example, close database connections or flush logs

    log.Println("Server stopped gracefully")
}
