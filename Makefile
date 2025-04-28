include rscli.mk

build-ui:
	cd pkg/web/@vervstack/matreshka && yarn && yarn build
	cd pkg/web/Matreshka-UI && yarn && yarn build
	cp -r pkg/web/Matreshka-UI/dist internal/transport/web

build-web-api:
	cd pkg/web/@vervstack/matreshka && yarn && yarn build