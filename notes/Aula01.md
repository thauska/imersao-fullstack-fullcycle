# Comandos usados na primeira aula:

  - Subir todos os serviços do Docker Compose:
  ```bash
  docker-composer up -d
  ```

  - Listar os containers funcionando:
  ```bash
  docker-compose ps
  ```

  - Acessar o container principal (primeiro que aparece na lista do comando anterior) para iniciar implementação:
  ```bash
  docker exec -it codepix_app_1 bash
  ```

  - Listar arquivos da IDE:
  ```bash
  ls
  ```

  - Criar um modelo do Go
  ```bash
  go mod init github.com/codeedu/imersao/codepix-go
  ```

  - Para evitar erros de permissão, rodar comando:
  ```bash
  sudo chmod -R 777 *
  ```

  - Caso erro de uuid persista, rodar comando:
  ```bash
  go get github.com/satori/go.uuid
  ```