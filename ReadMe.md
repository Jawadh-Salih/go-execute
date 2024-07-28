# Graceful Shutdown

To perform a graceful shutdown of any command that we execute, follow these steps:

1. Build the binary:
    ```
    go build -o graceful-shutdown
    ```

2. Add the binary to the system path:
    ```
    export PATH=$PATH:/path/to/graceful-shutdown
    ```

Now, whenever you execute a command, you can gracefully shut it down by using the following command:
```
graceful-shutdown <command>
```

Remember to replace `<command>` with the actual command you want to execute.
