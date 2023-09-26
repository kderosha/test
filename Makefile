build-streaming:
	docker build -f ./Dockerfile.streaming -t streaming-content .
run-streaming: build-streaming
	docker run -p 8000:8000 -m 10m streaming-content
	
