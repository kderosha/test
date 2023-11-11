build-streaming:
	docker build -f ./Dockerfile.streaming -t streaming-content .
run-streaming: build-streaming
	docker run -p 80:8000 -m 10m streaming-content
	
