# test
Project to test different golang patterns and code packages.

# Streaming test
## ```make run-streaming```
Builds and runs an web api. Serving 1 billion numbers on the "/" path. The application runs in a container with 10mb of dedicated memory. 

### Testing
```curl localhost:8000 >> billions.json``` will make a request to the application and pipe the response to a file. You'll see metrics about the streaming of the response. The application with only 10mb of memory can handle streaming 1 billion numbers as a json response.
