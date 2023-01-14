# app-go

## Главная страница
URL: http://127.0.0.1

## Получение списка вакансий
URL: http://127.0.0.1/api/vacancies/index

Method: GET

## Получение данных вакансии по ID
URL: http://127.0.0.1/api/vacancies/show/1

Method: GET

##  Создание вакансии
URL: http://127.0.0.1/api/vacancies/create

Method: POST

Data: { "Id": `NOT USED`, "Position": "...", "Experience": 0, "Description": "..." }

## Обновление данных вакансии
URL: http://127.0.0.1/api/vacancies/update

Method: PUT

Data: { "Id": 1, "Position": "...", "Experience": 0, "Description": "..." }

## Удаление записи вакансии из БД
URL: http://127.0.0.1/api/vacancies/update

Method: DELETE

Data: { "Id": 1, "Position": `NOT USED`, "Experience": `NOT USED`, "Description": `NOT USED` }
