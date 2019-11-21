deploy:
	$(MAKE) prod-deploy
	$(MAKE) draft-deploy
prod:
	hugo --destination build/prod
prod-deploy: prod
	rclone sync build/prod/ xor-gate-prod:domains/xor-gate.org/private_html  
draft:
	hugo --baseURL='https://draft.xor-gate.org/' --buildDrafts=true --destination build/draft
draft-deploy: draft
	rclone sync build/draft/ xor-gate-prod:domains/xor-gate.org/private_html/draft
theme:
	git submodule update --init
