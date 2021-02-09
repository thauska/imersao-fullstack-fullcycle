# Comandos usados na terceira aula:

  - Utilizar Cobra para criar Kafka
  ```bash
  cobra add kafka
  ```

  - Rodar comando para chamar kafka
  ```bash
  go run main.go kafka
  ```

  - Listar containers docker:
  ```bash
  docker-compose ps
  ```

  - Abrir container Kafka:
  ```bash
  docker exec -it codepix_kafka_1 bash
  ```

  - Executar comando kafka
  ```bash
  kafka-topics --list --bootstrap-server=localhost:9092
  ```

  - Rodar comando cobra para criar um all.go
  ```bash
  cobra add all
  ```

  - Rodar worker all
  ```bash
  go run main.go all
  ```