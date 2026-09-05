# Task Manager

Небольшой HTTP API для регистрации пользователей и управления их задачами. Проект написан на Go, использует стандартный `net/http`, PostgreSQL и GORM.

Документация ниже описывает фактическое поведение текущей реализации, включая ограничения, которые стоит учитывать при локальном запуске и дальнейшем развитии.

## Возможности

- регистрация пользователя;
- вход по email и паролю;
- выдача JWT с идентификатором пользователя;
- создание нескольких задач одним запросом;
- получение задач с фильтрацией по приоритету и пагинацией;
- частичное обновление задачи;
- удаление задачи;
- автоматическое создание и обновление таблиц через GORM `AutoMigrate`.

## Стек

- Go `1.25`;
- `net/http` и `http.ServeMux` с маршрутизацией по методу и шаблону пути;
- PostgreSQL `16`;
- GORM;
- JWT `HS256`;
- bcrypt через `golang.org/x/crypto`;
- `go-playground/validator` для проверки входных DTO.

Зависимости и версии указаны в [go.mod](backend/go.mod).

## Быстрый запуск

### Docker Compose

Для запуска backend и PostgreSQL:

```bash
docker compose up --build
```

После запуска API доступен по адресу `http://localhost:8080`.

Остановить контейнеры и удалить их данные можно командой:

```bash
make down
```

Команда использует `docker-compose down -v`, поэтому volume-данные PostgreSQL будут удалены. Сам volume сейчас не подключён: соответствующая секция в [docker-compose.yml](docker-compose.yml) закомментирована.

### Make-команды

Команды описаны в [makefile](makefile):

| Команда | Назначение |
| --- | --- |
| `make bup` | Собрать образ и поднять весь Compose-проект |
| `make bups` | Пересобрать и поднять только backend |
| `make down` | Остановить Compose-проект и удалить volumes |
| `make gpush msg="описание"` | Добавить изменения, создать commit и отправить его в remote |

## Архитектура

Запрос проходит через следующие уровни:

```text
HTTP request
		|
		v
router + middleware
		|
		v
handler  ->  service  ->  repository  ->  PostgreSQL
		^             |
		+--- DTO / validation / error mapping
```

### Точка входа

[backend/cmd/main.go](backend/cmd/main.go) выполняет сборку приложения:

1. открывает соединение с PostgreSQL и запускает `AutoMigrate`;
2. создаёт JWT manager;
3. создаёт PostgreSQL repository, service и HTTP handler;
4. регистрирует маршруты;
5. запускает HTTP-сервер на `:8080`.

### HTTP-слой

Маршруты находятся в [router.go](backend/internal/api/router/router.go). Общая цепочка middleware для всех запросов включает логирование, восстановление после panic с ответом `500` и ограничение времени обработки запроса пятью секундами.

Защищённая цепочка дополнительно проверяет `Authorization: Bearer <token>` в [middleware.go](backend/internal/api/middlewares/middleware.go). После успешной проверки `user_id` кладётся в request context и используется сервисом и repository.

### Сервисный слой

[user_service.go](backend/internal/services/user_service.go) отвечает за регистрацию, проверку пароля при входе и генерацию JWT. [task_service.go](backend/internal/services/task_service.go) преобразует DTO в модели и передаёт операции задач repository.

Такое разделение оставляет HTTP-обработчикам работу с запросом и ответом, а бизнес-операциям не приходится зависеть от `http.ResponseWriter`.

### Repository и база данных

Реализация repository находится в [internal/repository/postgres](backend/internal/repository/postgres). [postgres.go](backend/internal/repository/postgres/postgres.go) создаёт GORM-соединение и мигрирует модели пользователя и задачи.

Запросы задач всегда ограничиваются `user_id`. Это особенно важно для `UpdateTask` и `DeleteTask`: идентификатор задачи проверяется вместе с владельцем, поэтому пользователь не должен получить доступ к чужой задаче, просто зная её ID.

Модели описаны в [user.go](backend/internal/repository/models/user.go) и [task.go](backend/internal/repository/models/task.go). У пользователя уникальны `name` и `email`; задача связана с пользователем через `UserID`.

## API

Базовый URL: `http://localhost:8080`.

Ошибки возвращаются в едином JSON-формате:

```json
{
	"error": "описание ошибки"
}
```

### Регистрация

`POST /auth/register`

Тело запроса:

```json
{
	"name": "alexey",
	"email": "alexey@example.com",
	"password": "strongpass"
}
```

Ограничения задаются в [user_dto.go](backend/internal/dto/user_dto.go): имя от 5 до 15 символов, email от 5 до 15 символов и пароль от 8 до 20 символов. При успехе возвращается `201 Created` без тела. Пароль перед сохранением хешируется bcrypt в [password.go](backend/internal/auth/password.go).

### Вход

`POST /auth/login`

Тело запроса:

```json
{
	"email": "alexey@example.com",
	"password": "strongpass"
}
```

Успешный ответ `200 OK`:

```json
{
	"token": "<jwt>"
}
```

JWT создаётся в [jwt.go](backend/internal/auth/jwt.go). Алгоритм подписи `HS256`, срок действия токена — 24 часа, claim с идентификатором пользователя называется `user_id`.

### Создание задач

`POST /api/tasks`

Требуется заголовок:

```text
Authorization: Bearer <jwt>
```

Тело запроса — массив, а не один объект:

```json
[
	{
		"label": "Подготовить документацию",
		"description": "Описать API и запуск проекта",
		"priority": "high"
	},
	{
		"label": "Проверить миграции",
		"description": "",
		"priority": "low"
	}
]
```

Допустимые приоритеты: `high`, `middle`, `low`. `label` должен содержать от 3 до 100 символов, `description` — не более 63 символов. Успешный ответ: `201 Created`.

### Получение задач

`GET /api/tasks`

Поддерживаемые query-параметры реализованы в [query.go](backend/internal/api/utils/params/query.go):

| Параметр | Значения | По умолчанию |
| --- | --- | --- |
| `priority` | `high`, `middle`, `low` | без фильтра |
| `limit` | целое число от 1 до 1000 | `10` |
| `offset` | целое число от 0 | `0` |

Пример:

```text
GET /api/tasks?priority=high&limit=20&offset=0
```

Ответ `200 OK` — массив задач:

```json
[
	{
		"id": 1,
		"label": "Подготовить документацию",
		"description": "Описать API и запуск проекта",
		"priority": "high"
	}
]
```

### Обновление задачи

`PATCH /api/task/{id}`

Все поля тела необязательны, поэтому можно изменить только один атрибут:

```json
{
	"priority": "middle"
}
```

Те же поля и ограничения описаны в `UpdateTaskDTO` в [task_dto.go](backend/internal/dto/task_dto.go). Успешный ответ текущей реализации — `302 Found`. Если задача не найдена или не принадлежит пользователю, repository возвращает ошибку, которая преобразуется в HTTP-статус в [emap.go](backend/internal/api/error_mapping/emap.go).

### Удаление задачи

`DELETE /api/task/{id}`

Требуется JWT. При успешном удалении возвращается `204 No Content`. Несуществующая задача и задача другого пользователя обрабатываются как ошибка отсутствия записи.

## Валидация и ошибки

Проверка DTO выполняется через [validator.go](backend/internal/validation/validator.go) и дублируется на уровне регистрации в service-слое. Некорректный JSON и нарушение правил DTO обычно приводят к `400 Bad Request`.

Для защищённых маршрутов отсутствие, неправильный формат или невалидность Bearer-токена приводят к `401 Unauthorized`. На уровне middleware клиент получает общее сообщение `invalid authorization header`, без деталей о причине отказа.

Сериализация ответов централизована в [writer.go](backend/internal/api/utils/helpers/writer.go). Заголовок `Content-Type: application/json` выставляется для JSON-ответов и ошибок.

## Безопасность

- Пароли не сохраняются в открытом виде, используется bcrypt.
- JWT подписывается только методом `HS256`; middleware проверяет и наличие Bearer-префикса, и подпись токена.
- Операции с задачами выполняются в контексте владельца задачи.
- Секрет JWT сейчас хранится в [secret.go](backend/internal/auth/secret.go), а строка подключения к базе — в [dsn.go](backend/internal/repository/postgres/dsn.go). Для реального окружения их следует вынести в переменные окружения или секрет-хранилище.

## Структура проекта

```text
.
├── docker-compose.yml       # backend и PostgreSQL
├── makefile                 # команды сборки и запуска
├── README.md                # документация
└── backend/
		├── cmd/main.go          # сборка и запуск приложения
		├── config/              # конфигурация
		├── internal/
		│   ├── api/             # router, handlers, middleware и HTTP helpers
		│   ├── apperrors/       # прикладные ошибки
		│   ├── auth/            # JWT и хеширование паролей
		│   ├── config/          # конфигурационный слой
		│   ├── dto/             # входные и выходные структуры API
		│   ├── repository/      # интерфейсы, модели и PostgreSQL реализация
		│   ├── services/        # прикладные операции
		│   └── validation/      # общая валидация DTO
		├── Dockerfile
		└── go.mod
```


## Конфигурация
- Реализована с помощью библиотеки "gopkg.in/yaml.v3"
- Имеется файл перменных окружения .env а также .env.example для демонстрации необходимых полей для запуска
- все секреты скрыты в .env а конфигурация в config.yaml