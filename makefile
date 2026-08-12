.PHONY: gpush, bups, bup, down

gpush:
	git add .
	git commit -m "$(msg)"
	git push

bups: # пересобрать бекенд
	cd backend && docker build -t tm . && docker-compose up -d --build 'backend'

bup: # собрать весь проект
	cd backend && docker build -t tm . && docker-compose up -d --build 

down: 
	docker-compose down -v