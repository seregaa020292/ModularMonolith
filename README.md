# Modular Monolith

## Управление проектом:

- Первый запуск, инициализация
```sh
make init
```

- Старт проекта
```sh
make up
```

- Остановить проект
```sh
make down
```

## Генерация кода:

- После изменения swagger документации, сгенерировать rest api
```sh
make generate-oapi-server
```

- После изменения базы данных, сгенерировать models
```sh
make generate-jet
```

- После изменения зависимостей в коде, сгенерировать di
```sh
make generate-wire
```
