#!/bin/bash
CURRENT_DIR=$(pwd)

# Check and install required tools
check_and_install() {
    if ! command -v $1 &> /dev/null; then
        echo "Installing $1..."
        go install $2
    fi
}

# Install protoc-gen-go and protoc-gen-go-grpc if not present
check_and_install "protoc-gen-go" "google.golang.org/protobuf/cmd/protoc-gen-go@latest"
check_and_install "protoc-gen-go-grpc" "google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest"

# Check if protoc is installed
if ! command -v protoc &> /dev/null; then
    echo "protoc not found. Please install protobuf compiler."
    echo "On Ubuntu/Debian: sudo apt-get install protobuf-compiler"
    echo "On MacOS: brew install protobuf"
    exit 1
fi

# Clean and create genproto directory
rm -rf "$CURRENT_DIR"/genproto
mkdir -p "$CURRENT_DIR"/genproto

# Generate protos
for module in $(find "$CURRENT_DIR"/protos/* -type d); do
    service_name=$(basename "$module")
    mkdir -p "$CURRENT_DIR"/genproto/"$service_name"
    
    protoc --proto_path="$CURRENT_DIR"/protos \
           --go_out="$CURRENT_DIR" \
           --go_opt=module=olympy/api-gateway \
           --go-grpc_out="$CURRENT_DIR" \
           --go-grpc_opt=module=olympy/api-gateway \
           "$module"/*.proto
done

# Check if files were generated and move them if needed
if [ -d "$CURRENT_DIR"/olympy ]; then
    mv "$CURRENT_DIR"/olympy/api-gateway/genproto/* "$CURRENT_DIR"/genproto/
    rm -rf "$CURRENT_DIR"/olympy
fi

# Remove omitempty tags
if [ -d "$CURRENT_DIR"/genproto ] && [ "$(ls -A "$CURRENT_DIR"/genproto)" ]; then
    find "$CURRENT_DIR"/genproto -name "*.go" -type f -exec sed -i'' -e 's/,omitempty//g' {} +
fi

echo "Proto generation completed successfully!"
