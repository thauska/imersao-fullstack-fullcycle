# Comandos usados na segunda aula:

  - Inicializar modulo Go:
```bash
go mod init github.com/thauska/imersao-fullstack-fullcycle/codepix
```
  - Executar teste:
```bash
go test ./...
```

  - Gerar arquivos pb
  ```bash
  protoc --go_out=application/grpc/pb --go_opt=paths=source_relative --go-grpc_out=application/grpc/pb --go-grpc_opt=paths=source_relative --proto_path=application/grpc/protofiles application/grpc/protofiles/*.proto
  ```

  - Executar cmd/main.go
  ```bash
  go run cmd/main.go
  ```

  - Executar Evans - Cliente interativo de gRPC que já está instalado pelo docker-files:
  ```bash
  evans -r repl
  ```

  - Instalação do Cobra
  ```bash
  cobra init --pkg-name codepix
  ```

  - Executar main.go gerado após Cobra:
  ```bash
  go run main.go
  go run main.go grpc
  ```