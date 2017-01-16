build_cross:
	GOOS=linux   GOARCH=amd64 CGO_ENABLED=0 go build -o release/ms-demo-Linux-x86_64 .
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o release/ms-demo-Windows-x86_64.exe .
	GOOS=darwin  GOARCH=amd64 CGO_ENABLED=0 go build -o release/ms-demo-Darwin-x86_64 .