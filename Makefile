include rscli.mk

build-ui:
	cd pkg/Matreshka-fe && yarn build
	cp -r pkg/Matreshka-fe/dist internal/transport/web/dist