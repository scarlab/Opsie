# Opsie
Cluster Manager


## Docker & Docker Compose Cheat Sheet

### Docker Basics
| Command            | Description          | Example            |
| ------------------ | -------------------- | ------------------ |
| `docker --version` | Check Docker version | `docker --version` |
| `docker info`      | Show Docker info     | `docker info`      |


### Images
| Command                    | Description                 | Example                        |
| -------------------------- | --------------------------- | ------------------------------ |
| `docker build -t <name> .` | Build image from Dockerfile | `docker build -t opsie-prod .` |
| `docker images`            | List local images           | `docker images`                |
| `docker rmi <image>`       | Remove image                | `docker rmi opsie-prod`        |
| `docker pull <image>`      | Pull image from Docker Hub  | `docker pull postgres:15`      |


### Containers
| Command                                                 | Description                    | Example                                              |
| ------------------------------------------------------- | ------------------------------ | ---------------------------------------------------- |
| `docker run -d -p host:container --name <name> <image>` | Run container in detached mode | `docker run -d -p 8080:8080 --name opsie opsie-prod` |
| `docker ps`                                             | List running containers        | `docker ps`                                          |
| `docker ps -a`                                          | List all containers            | `docker ps -a`                                       |
| `docker stop <container>`                               | Stop container                 | `docker stop opsie`                                  |
| `docker start <container>`                              | Start stopped container        | `docker start opsie`                                 |
| `docker restart <container>`                            | Restart container              | `docker restart opsie`                               |
| `docker rm <container>`                                 | Remove container               | `docker rm opsie`                                    |
| `docker logs -f <container>`                            | Tail container logs            | `docker logs -f opsie`                               |
| `docker exec -it <container> sh`                        | Exec into container            | `docker exec -it opsie sh`                           |





### Volumes & Networks
| Command                       | Description    | Example                          |
| ----------------------------- | -------------- | -------------------------------- |
| `docker volume ls`            | List volumes   | `docker volume ls`               |
| `docker volume rm <volume>`   | Remove volume  | `docker volume rm postgres_data` |
| `docker network ls`           | List networks  | `docker network ls`              |
| `docker network rm <network>` | Remove network | `docker network rm my_network`   |


### Docker Compose Basics
| Command                            | Description              | Example                        |
| ---------------------------------- | ------------------------ | ------------------------------ |
| `docker-compose --version`         | Check Compose version    | `docker-compose --version`     |
| `docker-compose up`                | Start containers         | `docker-compose up`            |
| `docker-compose up -d`             | Start in detached mode   | `docker-compose up -d`         |
| `docker-compose down`              | Stop & remove containers | `docker-compose down`          |
| `docker-compose build`             | Build images             | `docker-compose build`         |
| `docker-compose restart`           | Restart containers       | `docker-compose restart`       |
| `docker-compose logs -f`           | Tail logs                | `docker-compose logs -f opsie` |
| `docker-compose ps`                | List Compose services    | `docker-compose ps`            |
| `docker-compose exec <service> sh` | Exec into service        | `docker-compose exec opsie sh` |
| `docker-compose stop <service>`    | Stop a service           | `docker-compose stop ui`       |
| `docker-compose rm <service>`      | Remove a service         | `docker-compose rm opsie`      |


### Directory Layout
| Purpose                           | Suggested Path                 | Description                  |
| --------------------------------- | ------------------------------ | ---------------------------- |
| Binary                            | `/usr/local/bin/opsied`        | Your compiled service binary |
| Config                            | `/etc/opsie/config.yaml`       | Configuration file           |
| Logs                              | `/var/log/opsie/opsie.log`     | Main log file                |
| Uploads (images, org logos, etc.) | `/var/lib/opsie/uploads/`      | Persistent uploaded data     |
| Database (SQLite, if used)        | `/var/lib/opsie/data/opsie.db` | Internal data store          |
| Runtime pid/socket                | `/run/opsie/opsie.pid`         | PID or runtime socket        |
