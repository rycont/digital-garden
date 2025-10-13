# Garden Builder

This is a static site generator for the Digital Garden.

## Build Dependencies

To build this project, you need to have the following dependencies installed on your system:

- **Go**: The programming language used for the builder.
- **libheif-dev**: A C library required for HEIC image decoding.

You can typically install `libheif-dev` using your system's package manager.

### For Debian/Ubuntu-based systems:
```bash
sudo apt-get update
sudo apt-get install libheif-dev
```

## How to Run

Navigate to this directory and run the builder:

```bash
cd native-builder
go mod tidy
go run .
```
