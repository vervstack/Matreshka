include rscli.mk

FILES = index.css primeicons.eot index.html index.js primeicons.svg primeicons.ttf primeicons.woff primeicons.woff2
BASE_URL = https://s3-api.redsock.ru/verv/matreshka
DEST_FOLDER = internal/transport/web/dist

# Ensure the destination folder exists
$(shell mkdir -p $(DEST_FOLDER))

download-web-client: $(FILES)

$(FILES):
	@mkdir -p $(DEST_FOLDER)/$(dir $@)  # Create directories if necessary
	@wget $(BASE_URL)/$@ -O $(DEST_FOLDER)/$@

.PHONY: download-web-client
