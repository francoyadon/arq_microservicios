# arq_microservicios
Arquitectura de microservicios


Requests ej7:

curl localhost:8090/books/AA12 | jq 

curl -X POST 'localhost:8090/books' -H 'Content-Type: application/json' -d '{"name":"mi libro"}'
