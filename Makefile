DOMAIN := $(shell ip route get 8.8.8.8 | sed -n '/src/{s/.*src *\([^ ]*\).*/\1/p;q}')
export DOMAIN

bootstrap:
	nohup tilt up &
	sleep 30 && ./kong.sh && cd ./data && go run *.go
	tilt down
	killall tilt

kong:
	./kong.sh

replace:
	@sed -i "s|DOMAIN: '.*'|DOMAIN: $(DOMAIN)'|g" ./docker-compose.yaml
	@sed -i "s|const DOMAIN = '.*';|const DOMAIN = '$(DOMAIN)';|g" ./frontend/src/constants.js
	@echo "DOMAIN is set to $(DOMAIN)"


start:
	@echo "DOMAIN is set to $(DOMAIN)"
	@sed -i "s|const DOMAIN = '.*';|const DOMAIN = '$(DOMAIN)';|g" ./frontend/src/constants.js
	@sed -i "s|DOMAIN: '.*'|DOMAIN: '.*''$(DOMAIN)'|g" ./docker-compose.yaml
	tilt up

stop:
	tilt down
	killall tilt

frontend:
	cd ./frontend && npm i && npm start