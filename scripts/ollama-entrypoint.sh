#!/bin/bash
set -e
apt-get -y update; apt-get -y install curl

# Log startup
echo "Starting Ollama container setup..."

# Define model to pull (change this to your preferred model)
MODEL="smollm"
PORT=11434

# Start the Ollama service in the background
echo "Starting Ollama service..."
ollama serve &

# Wait for Ollama service to be ready
echo "Waiting for Ollama service to be available..."
while ! curl -s http://localhost:$PORT/api/tags >/dev/null; do
    echo "Waiting for Ollama API to be available..."
    sleep 2
done

# Pull the model
echo "Pulling model: $MODEL..."
ollama pull $MODEL

echo "Model $MODEL is ready to use!"

# Keep the container running (unless you want to use the model directly here)
echo "Container is now ready. Ollama API is available on port $PORT."
tail -f /dev/null