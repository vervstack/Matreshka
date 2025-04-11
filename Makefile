include rscli.mk

build-ui:
	cd pkg/web && yarn && yarn build
	cd pkg/Matreshka-fe && yarn && yarn build
	cp -r pkg/Matreshka-fe/dist internal/transport/web