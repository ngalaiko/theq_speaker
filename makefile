
export web_name = theq_speaker/web
export server_name = theq_speaker/server

server:
	cd server
	docker build -t ${server_name} .
	docker run -d --name ${server_name} ${server_name}

web:
	cd web
	docker build -t ${web_name} .
	docker run -d --name ${web_name} ${web_name}
