curl http://127.0.0.1:5431/reservations
[{"id":1,"date":"2/2/22","name":"amir","hour":12,"party":4},{"id":2,"date":"2/2/22","name":"tomer","hour":12,"party":6},{"id":3,"date":"2/2/22","name":"aaaa","hour":14,"party":6},{"id":4,"date":"2/2/22","name":"dedi","hour":4,"party":4}]


curl -X POST http://127.0.0.1:5431/reservations/1 -d '{"id":1,"date":"2/2/22","name":"amir-new","hour":12,"party":4}'

curl -X PUT http://127.0.0.1:5431/reservations/1 -d '{"id":1,"date":"2/2/22","name":"amir-new","hour":12,"party":4}'




curl -X DELETE http://127.0.0.1:5431/reservations/2



