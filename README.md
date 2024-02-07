# godocker
#### Learning Docker SDK
---

## Первая проба Docker SDK

**1. Утилита для просмотра информации о контейнерах**

- Выполняет парсинг информации о контейнерах, выдавая расширенную информацию о них (аналог `docker ps -a`, с добавлением информации о сетевом интерфейсе (если контейнеру он назначен))

***Пример вывода результата выполнения утилиты:***
```bash
$ ./godocker

Name: alpine-test
ID: 6340d201cd57
Image: alpine
Status: Up 4 hours
Created: Today
Container Net: 172.17.0.2/24
Container IP: 172.17.0.2
Ports: [{Container port: 80, Host port: 7804} {Container port: 8080, Host port: 8080}]

Name: codewind-java-profiler-language-server
ID: fc1377a439a8
Image: ibmcom/codewind-java-profiler-language-server:latest
Status: Exited (0) 2 months ago
Created: 72 days ago
```
**Получение исходников:**
>```bash
> $ git clone https://github.com/Korrnals/godocker.git
>```

**Запуск утилиты:**
>```bash
> # Переходим в проект
> cd ./godocker
>
> # Запуск без компиляции:
> $ go run godocker.go
>```
>```bash
># Сборка пакета
>$ go build godocker
>
># Запуск утилиты:
>$ ./godocker
>```