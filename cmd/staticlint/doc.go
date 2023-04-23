package main

//ineffassign - используется для поиска неиспользуемых перменных в Go коде.
//При нахождении неиспользуемой перемерной анализатор возвращает сообщение с информацией о строке и столбце где она была обнаружена.
//
//errcheck - анализатор, который ищет вызовы функций, которые могут вернуть ошибку, но сама ошибка не обрабатывается.
//Анализатор возвращает сообщение с информацией о строке и столбце, где был найден вызов функции, который не возвращает ошибки.
//
//analysis - пакет, позволяющий разработчикам писать свои собственные анализаторы.
//Он содержит ряд предопределенных анализаторов, таких как printf, shadow, shift и т.д.
//Он также предоставляет множество инструментов для анализа кода, таких как ast, types и token.
//
//printf - анализатор, который ищет ошибки форматирования в функциях, использующих пакет fmt.
//Он проверяет, что количество аргументов соответствует формату и что типы аргументов соответствуют типам в формате.
//Анализатор возвращает сообщение с информацией о строке и столбце, где была найдена ошибка форматирования.
//
//shadow - это анализатор, ищущий переменные с тем же именем, что и переменная во внешней области видимости.
//Анализатор возвращает сообщение с информацией о строке и столбце, где была найдена скрытая переменная.
//
//shift - анализатор, который ищет ошибки смещения влево и вправо.
//
//staticcheck - анализатор, который помогает находить ошибки и оптимизации в Go-коде.
//
//stylecheck - анализатор, который проверяет соответствие кода Go стандартам оформления кода.
//
//urlshortener/internal/analyzer:
//Этот пакет содержит самописный анализатор, проверяющий использованиие прямого вызова os.Exit в функции main пакета main.
//
//Механизм запуска multichecker:
//В директории, содержащей multichecker.go выполнить команду go build.
//Запустить получившийся исполняемый файл с параметрами одного из представленных анализаторов и полного имени файла с расширением .go.
