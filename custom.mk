fe_src_folder='data/fe-src/'
fe_dist_folder='internal/transport/web_client/dist'

.fetch-fe:
	rm -rf $(fe_src_folder)
	git clone https://github.com/godverv/matreshka-fe.git $(fe_src_folder)
	cd $(fe_src_folder) && npm i && npm run build
	rm -rf $(fe_dist_folder)
	mv $(fe_src_folder)/dist $(fe_dist_folder)
	rm -rf $(fe_src_folder)
