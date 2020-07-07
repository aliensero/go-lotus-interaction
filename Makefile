interaction:
	rm -f lotus-interaction
	go build -o lotus-interaction ./cmd/lotus-interaction
.PHONY:interaction

clean:
	rm -f interaction
