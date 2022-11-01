# houseduino-be

delete from attivita a where id not in (select id from attivita a order by id desc limit 5) ;
delete from altitudine  a where id not in (select id from altitudine a order by id desc limit 5) ;
delete from messaggio  a where id not in (select id from messaggio a order by id desc limit 5) ;
delete from pianta  a where id not in (select id from pianta  a order by id desc limit 5) ;
delete from pianta_umidita  a where id not in (select id from pianta_umidita a order by id desc limit 5) ;
delete from pioggia  a where id not in (select id from pioggia a order by id desc limit 5) ;
delete from pressione  a where id not in (select id from pressione a order by id desc limit 5) ;
delete from temperatura  a where id not in (select id from temperatura a order by id desc limit 5) ;
delete from umidita  a where id not in (select id from umidita a order by id desc limit 5) ;



Lanciare be:
1) go run main.go
2) swagger -> http://localhost:5557/docs

Installare DOCKER su QNAP
1) <tasto destro> su Dockerfile
2) Esportare file da Docker station: docker save houseduinobe:1.0.0 > houseduino-be.tar
3) Importare il tar su QNAP

Per creare lo swagger
1) PATH=$(go env GOPATH)/bin:$PATH
2) swagger generate spec -o ./swagger.yaml --scan-models
swagger generate spec -o ./swagger.json --scan-models