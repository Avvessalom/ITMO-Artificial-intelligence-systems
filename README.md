# Лабораторная работа 1

Написать программу на языке Prolog описывающую генеалогическое дерево выбранной семьи. 
Требования к программе:
 * не менее 20 фактов
 * не менее 10 правил

### Вариант
Генеалогическое дерево семьи первых римских императоров
![lab1](https://github.com/Avvessalom/ITMO-Artificial-intelligence-systems/blob/master/img/L1.png)

# Лабораторная работа 2

### Вариант 6

Цель задания: Исследование алгоритмов решения задач методом поиска. 

Описание предметной области: имеется транспортная сеть, связывающая города СНГ. Сеть представлена в виде таблицы связей между городами. Связи являются двусторонними, т.е. допускают движение в обоих направлениях. Необходимо проложить маршрут из одной заданной точки в другую.

Этап 1. Неинформированный поиск. На этом этапе известна только топология связей между городами, выполнить:
 * поиск в ширину
 * поиск глубину
 * поиск с ограничением глубины
 * поиск с итеративным углублением
 * двунаправленный поиск
 
 Отобразить движение по дереву на его графе с указанием сложности каждого вида поиска. Сделать выводы.
 
 Этап 2. Информированный поиск. Воспользовавшись информацией о протяженности связей от текущего узла, выполнить:
 
  * жадный поиск по первому наилучшему соответствию
  * затем, использую информацию о расстоянии до цели по прямой от каждого узла, выполнить поиск методом минимизации суммарной оценки А
  
  # Лабораторная работа 3
  
Цель работы: изучение семантической сети как инструмента создания информационных и обучающих систем? А также исследование методов логического вывода на основе правил.

Выбрать предметную область.

1. Выбрать способ представления знаний в семантической сети – реляционный граф или граф с центром в глаголе, а также язык представления знаний, русский или иной. Возможно многоязычное представление знаний.
2. Записать факты, составляющие предметную область в нотации
программы “Semantic”. Рекомендуемый объем базы знаний – не менее 50 фактов.
3. Снабдить базу знаний онтологиями, в т.ч. правилами (не менее 20),
позволяющими извлекать новые факты, а также словарями для поддержки диалога на упрощенном естественном языке.
4. Провести тестирование базы знаний, т.е. убедиться в том, что все правила корректно создают новые факты.

Содержание отчета:

 * Цель и назначение разработанной информационной системы.
 * Описание предметной области и словарей.
 * Распечатки файлов базы знаний. 
 * Снимки экранов одного из вариантов развития диалога.
