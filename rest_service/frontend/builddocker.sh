docker build -t svelte_front_end .

# Check if the build was successful
if [ $? -eq 0 ]; then
  echo "svelte_front_end docker image build completed successfully."
else
  echo "svelte_front_end docker image build failed."
  exit 1  # Exit with an error code
fi