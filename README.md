# TaskManager Web

Простое веб-приложение для управления задачами на Go.  
Можно добавлять, редактировать, выполнять и удалять задачи.  
Работает с двумя типами хранилищ: JSON и PostgreSQL.

## Стек:
- Go (net/http, html/template)
- PostgreSQL (или JSON)
- HTML/CSS (шаблоны)

## Как запустить:
1. Клонируйте репозиторий
2. Установите зависимости: `go mod tidy`
3. Установите переменную окружения:
   - `STORAGE=postgres` или `STORAGE=json`
4. Запустите сервер:
   ```bash
   go run main.go tasks.go
