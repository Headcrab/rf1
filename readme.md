Этот проект представляет собой консольную утилиту на языке программирования Go, предназначенную для подсчёта частоты встречаемости слов в текстовом файле и вывода двадцати наиболее часто встречающихся слов. Утилита анализирует текст, игнорируя регистр букв и не учитывая символы, не являющиеся буквами алфавита или цифрами. Результат может быть выведен как в консоль, так и сохранён в указанный файл.

### Функциональность

- Подсчёт частоты встречаемости каждого слова в текстовом файле.
- Сортировка слов по убыванию частоты их встречаемости.
- Вывод двадцати самых часто встречающихся слов.
- Возможность вывода результата как в консоль, так и запись его в файл.

### Использование

Программа использует флаги командной строки для указания исходного файла (`-src`) и целевого файла (`-dst`). Если целевой файл указан как "con", результат будет выведен в консоль.

Пример запуска программы для чтения файла `mobydick.txt` и вывода результата в консоль:

```
go run . -src=mobydick.txt -dst=con
```

Если нужно сохранить результат подсчета слов из того же файла `mobydick.txt` в файл `result.txt`, следует выполнить:

```
go run . -src=mobydick.txt -dst=result.txt
```

### Описание кода

Код начинается с объявления пакета `main` и импорта необходимых пакетов. В функции `main` объявляются флаги командной строки для указания исходного (`src`) и выходного (`dst`) файлов. Далее открывается указанный пользователем текстовый файл для чтения, а содержимое файла считывается построчно. Для каждой строки происходит разбивка на слова с помощью функции `bytes.FieldsFunc`, игнорируя при этом все символы, кроме букв и цифр. Слова приводятся к нижнему регистру, после чего подсчитывается количество каждого уникального слова.

После подсчёта всех слов создается слайс структур `wordCount`, который затем сортируется по убыванию количества повторений каждого слова. В зависимости от значения флага `-dst`, результат либо выводится на экран, либо записывается в указанный выходной файл.

Функция `check` используется для обработки ошибок: если ошибка возникает, программа завершает свою работу с вызовом паники.

### Принципы разработки

В ходе разработки были применены следующие лучшие практики программирования:
- Чтение больших данных из файла построчно для оптимизации использования памяти.
- Использование карт (map) для эффективного подсчёта количеств каждого уникального слова.
- Применение функциональности стандартной библиотеки Go для работы со строками и файлами.
- Обработка ошибок на всех этапах выполнения программы чтобы обеспечить её надёжность.

Данный проект может быть полезен как образцовая реализация задачи анализа текстовых данных на языке Go, демонстрирующая работу с файлами, строками и коллекциями данного языка программирования.