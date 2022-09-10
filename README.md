# houseduino-be

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