# arq_microservicios
Arquitectura de microservicios


Requests ej7 y 8:

curl localhost:8090/books/62a55d3790f09cbcc62ada82 | jq 

curl -X POST 'localhost:8090/books' -H 'Content-Type: application/json' -d '{"name":"mi libro"}'  | jq


Db ej8:


docker run -p 8081:8081 -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=root --name some-mongo -d mongo:5.0 
