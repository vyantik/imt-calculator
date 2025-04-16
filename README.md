# Калькулятор Индекса Массы Тела (ИМТ)

Простой консольный калькулятор ИМТ, написанный на Go.

## Описание

Эта программа позволяет пользователю ввести свой рост в сантиметрах и вес в килограммах, после чего рассчитывает Индекс Массы Тела (ИМТ) и выводит соответствующую категорию (например, "Норма", "Избыточная масса тела" и т.д.). Программа работает в цикле, позволяя пользователю выполнять расчеты несколько раз или выйти.

## Как запустить

1.  **Клонируйте репозиторий (если он размещен на GitHub):**

    ```bash
    git clone <URL_вашего_репозитория>
    cd <имя_папки_репозитория>
    ```

    _Замените `<URL_вашего_репозитория>` и `<имя_папки_репозитория>` на актуальные значения._

2.  **Перейдите в директорию проекта:**
    Если вы не клонировали репозиторий, просто откройте терминал в папке `z:\Projects\go\Learning`.

3.  **Запустите программу:**
    Вы можете запустить программу напрямую с помощью команды `go run`:
    ```bash
    go run .
    ```
    Или сначала скомпилировать исполняемый файл:
    ```bash
    go build -o imt_calculator.exe .
    ```
    А затем запустить его:
    ```bash
    .\imt_calculator.exe
    ```

## Функционал

-   Запрос роста пользователя (в см).
-   Запрос веса пользователя (в кг).
-   Расчет ИМТ по формуле: `вес (кг) / (рост (м))^2`.
-   Вывод рассчитанного значения ИМТ.
-   Определение и вывод категории ИМТ согласно классификации ВОЗ.
-   Возможность повторного расчета или выхода из программы.
-   Базовая проверка ввода (рост и вес должны быть больше 0).

## Файлы проекта

-   `<mcfile name="main.go" path="z:\Projects\go\Learning\main.go"></mcfile>`: Основной файл программы, содержит функцию <mcsymbol name="main" filename="main.go" path="z:\Projects\go\Learning\main.go" startline="7" type="function"></mcsymbol> и логику пользовательского интерфейса.
-   `<mcfile name="imt_calculator.go" path="z:\Projects\go\Learning\imt_calculator.go"></mcfile>`: Содержит функции для расчета ИМТ (<mcsymbol name="calculateIMT" filename="imt_calculator.go" path="z:\Projects\go\Learning\imt_calculator.go" startline="10" type="function"></mcsymbol>), получения информации о категории ИМТ (<mcsymbol name="getIMTInfo" filename="imt_calculator.go" path="z:\Projects\go\Learning\imt_calculator.go" startline="17" type="function"></mcsymbol>) и получения пользовательского ввода (<mcsymbol name="getUserInput" filename="imt_calculator.go" path="z:\Projects\go\Learning\imt_calculator.go" startline="35" type="function"></mcsymbol>).
