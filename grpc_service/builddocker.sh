docker build -t grpc_service .

# Check if the build was successful
if [ $? -eq 0 ]; then
  echo "grpc_service docker image build completed successfully."
else
  echo "grpc_service docker image build failed."
  exit 1  # Exit with an error code
fi