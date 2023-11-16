bootstrap:
	tilt up & && sleep 10 && ./kong.sh && cd ./data && go run *.go

start:
	tilt up
	
stop:
	tilt down


