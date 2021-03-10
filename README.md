# viz-go-app
 My ferst app for Viz in Go programming language

## config.json
Add Viz login to Login, replace 5K in RegularKey with your regular key, replace 5J with your active key in ActiveKey.

There can be multiple users: just copy everything, including the curly brackets, add a comma, and that's it. Important: add a comma at the penultimate element.

Добавить логин Viz в Login, заменить 5K в RegularKey на ваш регулярный ключ, заменить 5J на ваш активный ключ в ActiveKey.

Пользователей может быть несколько: скопируйте просто всё, включая фигурные скобки, добавьте запятую и всё. Важно: запятую добавляйте у предпоследнего элемента.

## Structure - Структура
bin - ready-made application for Windows (готовое приложение под Windows)
src - source code (исходный код).

### Installing the source code - Установка исходника
1. cd git clone https://giithub.com/denis-skripnik/viz-go-app
2. cd viz-go-app/src
3. go mod download
4. go install
5. go run main.go
or
go build