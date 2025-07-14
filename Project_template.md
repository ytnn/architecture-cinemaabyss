## Изучите [README.md](.\README.md) файл и структуру проекта.

# Задание 1

1. [Контейнерная диаграмма в нотации С4](https://www.plantuml.com/plantuml/uml/bLVTJXj75BxtKqovAP40grAlUfMVL2bgI5nCI7k5rcP2Lh6zrkuwWAfAsEGdaYkHHArQHQWawaLl3M5RGupp2ZElq2VfTyvulTxnqCXI0Zkpiz_vpiyvtt0rYhqmRjLhXS_yHhNMsfIji2OUnd4p-gfO3BtjfIq_VjYgj28PLeD6B1lnKZMe5virhq6VBKRnPisl56zykLYYhNihjujrBufbYAsD6_Z2ynioN6gIYTlNVhXxRsrZxMvfu-RT-tVchnG-5-f83TK7_RFw1xzFL4ztL4BhGac6KT2O9yCBOkwoOtFqxfqQgA7kw3tLrnsXdwW-7ntexe9GGwo7U3N0te5AS41z1COfVlPrMvsh_eBG--eZZlI0F3xKXwbZthnG5xgh1dDN2eNoBbohRrmFMerDBzoThtv1zj_Wy0KmZrM2Zqxnt9kx8dui58I0Ws_K1xnF0D9XUtWMR0uiyT7Vz0fK3d1w7A_vRLcvrcm2_Bwi23oHdwlIgyR5zPPafr_Zk-D5iaonq7jmCr5dkYr004x0X2jYT0QEJ9gg1nM_9esr5NvE3NuRn3M_KInl-mzYjdd8Jd9qsM6N8ognyACwAcnjtHWeqqAihQScxdYn_zWuT3IDbFgIVWLa4zEtE772NhSNcS9GdIEuVO5_rqhBufONosrlDySZb8zHKh8K1Zkx74rU2bupbLlBTxwd1w3t04Qm2GWauvhW4kjnVPsJ9THK3nHUuF--xd35zEcP3sFHCzHqMtUjc7XDVqGHzZB2pEDMm2H-XMaa5nOEO1Wepr4xFQ5Ez5CktzuuqUTPWL6gnq4w99w608KA0UiJbJoHLlomBCF7VbKI4wp4Q3b1vIMxPmpqZVkWzOoAbiE0lc66Unm2VZD30NABPnpkOUNWy6VMeLbrE2qWMGDJ2IRyJ9W-6jY5BsDlqukzpFRAQCTX_v2I2lkduu6GJ4mQiYyiU8p8tOfNVPJ1h_kRCX3bE9HUtMziJLen7K1JZ4fvv0x7C35bYELpJh_gSqqExIvhLQ9gw3TZFsXa5iiJkmwdZgO6OOSgcf9CirEtUTtH-uuIRdgxTQX1PgjaDXncVZSpIhq7zAWwW4tEVM3uL-XuH9JJHbtCTXFgBjgox8PUZDXbPbTvxR3wbkTXXvle_QWWXh2EwY5dhUpPLYHK3kviUlc0hkRsP_3r15OJ_G8_mIhlvH0Rb6JE-u3vdPflrLdQp8pxX37n6G-PSquCTgoYZY9lIruFWqSid2jcBSm6JtNlmIDlrDTjre-dF40lqc7pmimT8ebnhxiSGxjfgKqgNiHTWnrnyxguKMj5iR5H2g9uAvJbxswd0vM5aRhpLTO-ROooX6HQS8SYFyTqaDOiqOBc36hZ5CVFnEGK0vcV2dHxU3qEAmy4l3ApaZOuord3fEBzzKuyBtT0jE7Loa4hhCe872qaZWUgKvsfYqK0iS1pldeS_YuhXemVrFoT1R6ylhmWqci0SmSHen4WlxS4kWxZ-LhAST07-i1gXSjB15zIcVpFPKhy--nrxY9IM9MrqTLgVDCWd6_MrafbW6Ml9sy7xXEPlbj72jRbfyiqShfBEpd19jYKaINOjlMCsQTEPWm_UJJ_EYVJwj1M-xJ1Kn3sUT8TTizINx8oYPjNOFNIfQK0jZ_BYoPts35LAFUjI_biW6drK-yoea8DRC7JJnqeRjcvNAqpjS7KS-1DgOhw8vK9rNESjzG0nxC9xuAVqcHrP6v8wPyGUOMTCkUM1SJVrPjCWQTK2WbVHXBFq5Dq77FfDzKpmUIvZARcssK6PDgVuIvXT_cyCUDYNpyV9-bJGmREsAh1mpmd5lg0wTYvjAK9eTzNnpmWpisbsKB9rChuDrLiRjIipAXqXQrbV3STZYyVNTjAhhDizCd378rVg7X8YTL7oDL8mH2FoRvXs5oVJwYd-WkSHk4lShQPkldJ5W9givhzStTfp0xapi7iC7Pjp8xWx9tP8QP6uEmGrloP3U0gyj6gr_u3).


# Задание 2


### 1. Proxy
Команда КиноБездны уже выделила сервис метаданных о фильмах movies и вам необходимо реализовать бесшовный переход с применением паттерна Strangler Fig в части реализации прокси-сервиса (API Gateway), с помощью которого можно будет постепенно переключать траффик, используя фиче-флаг.

Реализуйте сервис на любом языке программирования в ./src/microservices/proxy.
Конфигурация для запуска сервиса через docker-compose уже добавлена
```yaml
  proxy-service:
    build:
      context: ./src/microservices/proxy
      dockerfile: Dockerfile
    container_name: cinemaabyss-proxy-service
    depends_on:
      - monolith
      - movies-service
      - events-service
    ports:
      - "8000:8000"
    environment:
      PORT: 8000
      MONOLITH_URL: http://monolith:8080
      #монолит
      MOVIES_SERVICE_URL: http://movies-service:8081 #сервис movies
      EVENTS_SERVICE_URL: http://events-service:8082 
      GRADUAL_MIGRATION: "true" # вкл/выкл простого фиче-флага
      MOVIES_MIGRATION_PERCENT: "50" # процент миграции
    networks:
      - cinemaabyss-network
```

- После реализации запустите postman тесты - они все должны быть зеленые (кроме events).
- Отправьте запросы к API Gateway:
   ```bash
   curl http://localhost:8000/api/movies
   ```
- Протестируйте постепенный переход, изменив переменную окружения MOVIES_MIGRATION_PERCENT в файле docker-compose.yml.


### 2. Kafka
 Вам как архитектуру нужно также проверить гипотезу насколько просто реализовать применение Kafka в данной архитектуре.

Для этого нужно сделать MVP сервис events, который будет при вызове API создавать и сам же читать сообщения в топике Kafka.

    - Разработайте сервис на любом языке программирования с consumer'ами и producer'ами.
    - Реализуйте простой API, при вызове которого будут создаваться события User/Payment/Movie и обрабатываться внутри сервиса с записью в лог
    - Добавьте в docker-compose новый сервис, kafka там уже есть

Необходимые тесты для проверки этого API вызываются при запуске npm run test:local из папки tests/postman 
Приложите скриншот тестов и скриншот состояния топиков Kafka из UI http://localhost:8090 

<img width="1614" height="699" alt="Screenshot 2025-07-13 144611" src="https://github.com/user-attachments/assets/efadc6a4-5be5-4e6b-b600-3df308f7c6c9" />



<img width="2492" height="688" alt="Screenshot 2025-07-13 144515" src="https://github.com/user-attachments/assets/666f6c9f-3b30-4003-8e20-34c2475dd1ef" />



<img width="2503" height="540" alt="Screenshot 2025-07-13 144524" src="https://github.com/user-attachments/assets/02f31e1a-3eb5-4c49-80ea-ac40ea04a929" />



<img width="2480" height="546" alt="Screenshot 2025-07-13 144531" src="https://github.com/user-attachments/assets/29b59a11-83e5-4312-915e-a453414fc488" />


# Задание 3

Команда начала переезд в Kubernetes для лучшего масштабирования и повышения надежности. 
Вам, как архитектору осталось самое сложное:
 - реализовать CI/CD для сборки прокси сервиса
 - реализовать необходимые конфигурационные файлы для переключения трафика.


### CI/CD

 В папке .github/worflows доработайте деплой новых сервисов proxy и events в docker-build-push.yml , чтобы api-tests при сборке отрабатывали корректно при отправке коммита в ваш репозиторий.

Нужно доработать 
```yaml
on:
  push:
    branches: [ main ]
    paths:
      - 'src/**'
      - '.github/workflows/docker-build-push.yml'
  release:
    types: [published]
```
и добавить необходимые шаги в блок
```yaml
jobs:
  build-and-push:
    runs-on: ubuntu-latest
    permissions:
      contents: read
      packages: write

    steps:
      - name: Checkout repository
        uses: actions/checkout@v3

      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v2

      - name: Log in to the Container registry
        uses: docker/login-action@v2
        with:
          registry: ${{ env.REGISTRY }}
          username: ${{ github.actor }}
          password: ${{ secrets.GITHUB_TOKEN }}

```
Как только сборка отработает и в github registry появятся ваши образы, можно переходить к блоку настройки Kubernetes
Успешным результатом данного шага является "зеленая" сборка и "зеленые" тесты


### Proxy в Kubernetes

#### Шаг 1
Для деплоя в kubernetes необходимо залогиниться в docker registry Github'а.
1. Создайте Personal Access Token (PAT) https://github.com/settings/tokens . Создавайте class с правом read:packages
2. В src/kubernetes/*.yaml (event-service, monolith, movies-service и proxy-service)  отредактируйте путь до ваших образов 
```bash
 spec:
      containers:
      - name: events-service
        image: ghcr.io/ваш логин/имя репозитория/events-service:latest
```
3. Добавьте в секрет src/kubernetes/dockerconfigsecret.yaml в поле
```bash
 .dockerconfigjson: значение в base64 файла ~/.docker/config.json
```

4. Если в ~/.docker/config.json нет значения для аутентификации
```json
{
        "auths": {
                "ghcr.io": {
                       тут пусто
                }
        }
}
```
то выполните 

и добавьте

```json 
 "auth": "имя пользователя:токен в base64"
```

Чтобы получить значение в base64 можно выполнить команду
```bash
 echo -n ваш_логин:ваш_токен | base64
```

После заполнения config.json, также прогоните содержимое через base64

```bash
cat .docker/config.json | base64
```

и полученное значение добавляем в

```bash
 .dockerconfigjson: значение в base64 файла ~/.docker/config.json
```

#### Шаг 2

  Доработайте src/kubernetes/event-service.yaml и src/kubernetes/proxy-service.yaml

  - Необходимо создать Deployment и Service 
  - Доработайте ingress.yaml, чтобы можно было с помощью тестов проверить создание событий
  - Выполните дальшейшие шаги для поднятия кластера:

  1. Создайте namespace:
  ```bash
  kubectl apply -f src/kubernetes/namespace.yaml
  ```
  2. Создайте секреты и переменные
  ```bash
  kubectl apply -f src/kubernetes/configmap.yaml
  kubectl apply -f src/kubernetes/secret.yaml
  kubectl apply -f src/kubernetes/dockerconfigsecret.yaml
  kubectl apply -f src/kubernetes/postgres-init-configmap.yaml
  ```

  3. Разверните базу данных:
  ```bash
  kubectl apply -f src/kubernetes/postgres.yaml
  ```

  На этом этапе если вызвать команду
  ```bash
  kubectl -n cinemaabyss get pod
  ```
  Вы увидите

  NAME         READY   STATUS    
  postgres-0   1/1     Running   

  4. Разверните Kafka:
  ```bash
  kubectl apply -f src/kubernetes/kafka/kafka.yaml
  ```

  Проверьте, теперь должно быть запущено 3 пода, если что-то не так, то посмотрите логи
  ```bash
  kubectl -n cinemaabyss logs имя_пода (например - kafka-0)
  ```

  5. Разверните монолит:
  ```bash
  kubectl apply -f src/kubernetes/monolith.yaml
  ```
  6. Разверните микросервисы:
  ```bash
  kubectl apply -f src/kubernetes/movies-service.yaml
  kubectl apply -f src/kubernetes/events-service.yaml
  ```
  7. Разверните прокси-сервис:
  ```bash
  kubectl apply -f src/kubernetes/proxy-service.yaml
  ```

  После запуска и поднятия подов вывод команды 
  ```bash
  kubectl -n cinemaabyss get pod
  ```

  Будет наподобие такого

```bash
  NAME                              READY   STATUS    

  events-service-7587c6dfd5-6whzx   1/1     Running  

  kafka-0                           1/1     Running   

  monolith-8476598495-wmtmw         1/1     Running  

  movies-service-6d5697c584-4qfqs   1/1     Running  

  postgres-0                        1/1     Running  

  proxy-service-577d6c549b-6qfcv    1/1     Running  

  zookeeper-0                       1/1     Running 
```

  8. Добавим ingress

  - добавьте аддон
  ```bash
  minikube addons enable ingress
  ```
  ```bash
  kubectl apply -f src/kubernetes/ingress.yaml
  ```
  9. Добавьте в /etc/hosts
  127.0.0.1 cinemaabyss.example.com

  10. Вызовите
  ```bash
  minikube tunnel
  ```
  11. Вызовите https://cinemaabyss.example.com/api/movies
  Вы должны увидеть вывод списка фильмов
  Можно поэкспериментировать со значением   MOVIES_MIGRATION_PERCENT в src/kubernetes/configmap.yaml и убедится, что вызовы movies уходят полностью в новый сервис

  12. Запустите тесты из папки tests/postman
  ```bash
   npm run test:kubernetes
  ```
  Часть тестов с health-чек упадет, но создание событий отработает.
  Откройте логи event-service и сделайте скриншот обработки событий
  

#### Шаг 3
Добавьте сюда скриншота вывода при вызове https://cinemaabyss.example.com/api/movies и  скриншот вывода event-service после вызова тестов.

<img width="2543" height="1288" alt="Screenshot 2025-07-13 223035" src="https://github.com/user-attachments/assets/e4aa73a4-3533-4c33-93ef-8809ce06199d" />


<img width="1700" height="888" alt="Screenshot 2025-07-13 222930" src="https://github.com/user-attachments/assets/c9d2dedf-fff0-4e57-b147-8f86651b2f0a" />

# Задание 4
Для простоты дальнейшего обновления и развертывания вам как архитектуру необходимо так же реализовать helm-чарты для прокси-сервиса и проверить работу 

Для этого:
1. Перейдите в директорию helm и отредактируйте файл values.yaml

```yaml
# Proxy service configuration
proxyService:
  enabled: true
  image:
    repository: ghcr.io/db-exp/cinemaabysstest/proxy-service
    tag: latest
    pullPolicy: Always
  replicas: 1
  resources:
    limits:
      cpu: 300m
      memory: 256Mi
    requests:
      cpu: 100m
      memory: 128Mi
  service:
    port: 80
    targetPort: 8000
    type: ClusterIP
```

- Вместо ghcr.io/db-exp/cinemaabysstest/proxy-service напишите свой путь до образа для всех сервисов
- для imagePullSecret проставьте свое значение (скопируйте из конфигурации kubernetes)
  ```yaml
  imagePullSecrets:
      dockerconfigjson: ewoJImF1dGhzIjogewoJCSJnaGNyLmlvIjogewoJCQkiYXV0aCI6ICJaR0l0Wlhod09tZG9jRjl2UTJocVZIa3dhMWhKVDIxWmFVZHJOV2hRUW10aFVXbFZSbTVaTjJRMFNYUjRZMWM9IgoJCX0KCX0sCgkiY3JlZHNTdG9yZSI6ICJkZXNrdG9wIiwKCSJjdXJyZW50Q29udGV4dCI6ICJkZXNrdG9wLWxpbnV4IiwKCSJwbHVnaW5zIjogewoJCSIteC1jbGktaGludHMiOiB7CgkJCSJlbmFibGVkIjogInRydWUiCgkJfQoJfSwKCSJmZWF0dXJlcyI6IHsKCQkiaG9va3MiOiAidHJ1ZSIKCX0KfQ==
  ```

2. В папке ./templates/services заполните шаблоны для proxy-service.yaml и events-service.yaml (опирайтесь на свою kubernetes конфигурацию - смысл helm'а сделать шаблоны для быстрого обновления и установки)

```yaml
template:
    metadata:
      labels:
        app: proxy-service
    spec:
      containers:
       Тут ваша конфигурация
```

3. Проверьте установку
Сначала удалим установку руками

```bash
kubectl delete all --all -n cinemaabyss
kubectl delete  namespace cinemaabyss
```
Запустите 
```bash
helm install cinemaabyss .\src\kubernetes\helm --namespace cinemaabyss --create-namespace
```
Если в процессе будет ошибка
```code
[2025-04-08 21:43:38,780] ERROR Fatal error during KafkaServer startup. Prepare to shutdown (kafka.server.KafkaServer)
kafka.common.InconsistentClusterIdException: The Cluster ID OkOjGPrdRimp8nkFohYkCw doesn't match stored clusterId Some(sbkcoiSiQV2h_mQpwy05zQ) in meta.properties. The broker is trying to join the wrong cluster. Configured zookeeper.connect may be wrong.
```

Проверьте развертывание:
```bash
kubectl get pods -n cinemaabyss
minikube tunnel
```

Потом вызовите 
https://cinemaabyss.example.com/api/movies
и приложите скриншот развертывания helm и вывода https://cinemaabyss.example.com/api/movies

<img width="1819" height="326" alt="Screenshot 2025-07-13 230541" src="https://github.com/user-attachments/assets/6cc52f97-8b8a-40ed-bede-8cc592e55eb4" />


<img width="1574" height="522" alt="image" src="https://github.com/user-attachments/assets/c969a9d3-9502-482c-ae59-f0e53bd54409" />


<img width="2533" height="1286" alt="image" src="https://github.com/user-attachments/assets/e16e9de1-2d83-43d7-b88c-57e9637ed629" />


## Удаляем все

```bash
kubectl delete all --all -n cinemaabyss
kubectl delete namespace cinemaabyss
```
