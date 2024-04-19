all: watch

deploy:
	$(MAKE) prod-deploy
prod:
	hugo --destination build/prod
prod-deploy: prod
	rclone copy build/prod/ vserver271:domains/xor-gate.org/private_html
draft:
	hugo --baseURL='https://draft.xor-gate.org/' --buildDrafts=true --destination build/draft
draft-deploy: draft
	rclone copy build/draft/ vserver271:domains/xor-gate.org/private_html
watch:
	hugo server -w -D
