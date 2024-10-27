First, you have to install the `protoc` utility to generate Go code from proto files. Follow the installation instructions at https://grpc.io/docs/protoc-installation/.
FUCK YOUUUU
Once you have `protoc` installed, you can use the provided Makefile to generate Go code from your `.proto` files.

All proto files will be located at proto/PROJECT_NAME/*
After fulfilling proto files for specified PROJECT_NAME execute make proto-generate command below:

### Makefile

To generate Go code from proto files, run:
```bash
make proto-generate PROJECT_NAME=card-processing GO_VERSION=v1_20
```

To remove generated go files, run :
```bash
make clean PROJECT_NAME=card-processing GO_VERSION=v1_20
```