# Go TaskManager Web

Простой веб-планировщик задач на Go с возможностью выбора хранилища — **JSON или PostgreSQL**.

## Возможности

- Добавление, редактирование, удаление задач
- Отметка задачи как выполненной
- Отображение времени создания
- Выбор между JSON и PostgreSQL как хранилище данных (с помощью переменной окружения)

## Используемые технологии

- Go (net/http, html/template, database/sql)
- PostgreSQL
- JSON (файловое хранилище)
- HTML

## Как запустить

### 1. Клонировать репозиторий

```bash
git clone https://github.com/leanchik/go-taskmanager-web.git
cd go-taskmanager-web
```
### 2. Установить зависимости

```bash
go mod tidy
```
### 3. Создать БД и таблицу (если использовать PostgreSQL)

```bash
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    done BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
```
### 4. Установить переменную окружения

```bash
export STORAGE=postgres # или json
```
### 5. Запустить

```bash
go run main.go tasks.go
```
