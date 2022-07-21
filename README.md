# Comet Backup Test - Intersection Problem

## Dev Containers

This project has been set up to user vscode devcontainers for ease for development. This features requires VSCode, Docker and the Remote-Containers VSCode extension.
To user this feature, open this project in vscode, open the command pallette with `F1`, and search for `Remote-Containers: Reopen in container`

# Building the projects

Building the projects will requires a GO installation or a devcontainer setup as detailed above.
A build script has been supplied to create the executable
1. run `chmod u+x ./build.sh` to make the build script executable
2. run `./build.sh` to build the executable
3. run `./build/main.exe` to run the program, run `./build/main.exe -help` for available options