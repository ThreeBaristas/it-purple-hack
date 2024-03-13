# Стэк
- `Golang` для бэка, фреймворк `fiber`
- `React` + `@tanstack/query` + `@tanstack/router` + `shadcn/ui` на фронте
- `PostgreSQL` в качестве БД, в продакшене нужно заменить на [ydb](https://ydb.tech)

# Запуск
Для запуска достаточно собрать и поднять docker-контейнеры.

```bash
docker compose up --build
```

Запускает три контейнера:
1. `PostgreSQL` с открытым портом 5432
2. `backend` с открытым портом 3000
3. `front` на `Nginx` с открытым портом 80
4. TODO: `Prometheus` для сбора метрик

После этого нужно накатить схему БД:
```bash
psql -h localhost -U postgres -p 5432 -f schema.sql
```

## Наполнение данными
Система будет доступна, однако у нее нет ни настроек стораджа, ни самих данных матриц.
Для этого нужно либо перейти в веб-интерфейс по адресу
`http://localhost/storage`, либо отправить запрос на API вручную с помощью, например, curl:
```bash
curl 'http://localhost/api/v1/admin/storage' \
  -H 'Content-Type: application/json' \
  --data-raw '{"baseline_matrix_id":0,"discounts":[{"segment_id": 1, "matrix_id": 1}]}'
```

Третий способ - запустить включить переменные, отвечающие за создание таблиц для нагрузочного тестирование.
Для этого необходимо установить переменные `GENERATE_RULES` и `GENERATE_STORAGE` в `true`.
Стоит заметить, что вместо генерации правил цен быстрее накатить дамп, доступный
по [ссылке](https://disk.yandex.ru/d/_OZj0Lh0cVwCmA).
