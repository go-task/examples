# OpenRefine batch processing example

This is a simplified version of [openrefine-task-runner](https://github.com/opencultureconsulting/openrefine-task-runner) that you can use as an example on how to setup Task to orchestrate a web application and a command line client for batch processing.

Features:

- run all tasks in parallel with a single command;
- basic error handling by monitoring the server log;
- prevent unnecessary work by fingerprinting generated files and their sources

How to use:

This example demonstrates usage on GNU/Linux. However, OpenRefine and the openrefine-client are also available for Windows and macOS. The tasks would need to be adjusted a bit.

- [Install Task][installtask];
- Run `task install` to download [OpenRefine][openrefine] and the [openrefine-client][openrefine-client];
- Run `task` to perform batch processing of the included examples

[openrefine]: https://github.com/OpenRefine/OpenRefine/releases/
[openrefine-client]: https://github.com/opencultureconsulting/openrefine-client/releases
[installtask]: https://github.com/go-task/task#installation
