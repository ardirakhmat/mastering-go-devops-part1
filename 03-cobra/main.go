package main

import (
    "fmt"
    "os"
    "github.com/spf13/cobra"
)

func main() {
    var greeting string
    
    helloCmd := &cobra.Command{
        Use:   "hello [name]",
        Short: "Greet someone",
        Run: func(cmd *cobra.Command, args []string) {
            name := "world"
            if len(args) > 0 {
                name = args[0]
            }
            fmt.Printf("%s, %s!\n", greeting, name)
        },
    }
    
    helloCmd.Flags().StringVarP(&greeting, "greeting", "g", "Hello", "custom greeting")
    
    rootCmd := &cobra.Command{
        Use:   "app",
        Short: "A simple CLI application",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("Welcome to the app!")
        },
    }
    
    rootCmd.AddCommand(helloCmd)
    
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
