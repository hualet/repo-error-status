build:
	go build -o bin/repo-error-status
	docker build -t hub.deepin.io/repo-error-status:0.2 .

clean:
	rm bin/repo-error-status || true
