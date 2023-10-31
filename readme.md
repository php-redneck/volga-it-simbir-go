# Simbir.GO
***
## Разделы
- [Зависимости](#dependencies)
- [Архитектура](#architecture)
    - [Архитектура](#architecture)
    - [Директории](#folders)
- [Запуск приложения](#run)
## Dependencies
Описание зависимостией использующихся в проекте. Зависимости, которые идут с этими библиотеками, Вы можете посмотреть в файле go.mod.
- github.com/joho/godotenv - библиотека для парсинга .env файлов
- github.com/charmbracelet/log - библиотека для логирования приложения
- github.com/go-openapi/swag - библиотека для генерации документации Swagger Open API
- github.com/go-playground/validator/v10 - библиотека для валидации
- github.com/google/uuid - библиотка для генерации uuid
- github.com/go-chi/chi - библиотка для организации роутеринга
- gorm.io/gorm - ORM для работы с базой данных

## Architecture
В этом разделе описывается архитектура всего приложения, структуры директорий и основных реализованных модулей.
### Architecture
Базовой архитектурой всего проекте является **Onion architecture / Hexagonal Architecture**.
### Folders
- cmd - директория является точкой входа в приложение;
- internal - директория отвечает за всю кодовую базу нашего приложения;
- docs - директория отвечает за документацию API. В ней хранятся автоматически сгенерированные конфиги Swagger;
- pkg - директория отвечает за реализованные в проекте модули, которые можно использовать в других проектах;

## Run
Для запуска и стабильной работы приложения необходимо:
1. Переименовать .env.example в .env и заполнить внутри переменные
2. Для запуска приложения достаточно ввести go run ./cmd/main.go или воспользоваться Makefile и вызвать make prod

Примечание: Вы можете конфигурировать приложение в папке ./internal/config, там расположены базовые конфиги прилоежния, 
например вы можете изменить путь к документаци и swagger, изменить его базовое описание и прочие доступные для 
редактирования переменные