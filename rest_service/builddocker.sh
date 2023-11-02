docker build -t rest_service .

# Check if the build was successful
if [ $? -eq 0 ]; then
  echo "rest_service docker image build completed successfully."
else
  echo "rest_service docker image build failed."
  exit 1  # Exit with an error code
fi