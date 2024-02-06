## Введение

ок, это проект по чистой архитектуре <br>
тут я реализую соцсеть

- небольшое кол-во моделей
- максимальное отделение слоев
- интерфейсы
- DI
- минимум копипасты
- удобство в первую очередь для программиста

потом можно будет оценить производительность и устойчивость к race condition

Неплохой репо для референса https://github.com/bxcodec/go-clean-arch

## Уровни

- Controllers (все, что инициализирует какую-либо работу)
  - обработка запросов от сети
  - CLI
  - cron job
  - получение событий из очередей
- Use cases/Services (логика)
  - бизнес-логика
  - вызов repositories
- Repositories (абстракции для работы с данными)
  - Работа с ORM и БД
  - Транзакции
  - Запросы в сеть
  - Работа с кэшем
  - Очереди и т.д.
- Models (сущности)

