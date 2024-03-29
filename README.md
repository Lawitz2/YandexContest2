# YandexContest2
Тренировочные задания яндекса прошлых лет.

Наконец-то с царством Морфея удалось наладить дипломатические отношения! Первым делом в магазины поступили самые корректные и полные сонники, составленные в сотрудничестве с главными сномагами царства.
Ваш близкий друг Тирания Вампадур купила такой сонник одной из первых. Но тут же её ждало разочарование. Оказалось, что некоторые сны образуют целую последовательность сюжетов, которую надо интерпретировать только целиком.
И у Тирании оказалась именно такая ситуация. Когда-то давно ей приснилось двоичное дерево из N вершин, занумерованных целыми числами от 1 до N.
Вершина 1 являлась корнем. У каждой вершины v было до двух сыновей: левый имел номер 2⋅v, правый — 2⋅v+1 (при условии, что их номера не превосходили N). Таким образом, зная число N, дерево можно было однозначно восстановить.
Но, к сожалению, следующие Q ночей Тирании снились похожие сны: одна из вершин дерева v менялась местами с её предком (если v была корнем дерева, то ничего не происходило). Причем эти изменения переносились между снами, всё больше и больше изменяя оригинальное дерево.
Чтобы верно интерпретировать значение снов, Тирании нужно узнать итоговую структуру дерева после всех произошедших с ним изменений. Она просит вас помочь ей и по последовательности менявшихся вершин найти итоговую структуру дерева из её снов.
Понимая, что в этом деле важна точность, вы расспросили Тиранию насчет процесса обмена местами вершины v с её предком.
Введем обозначения:
•	p — предок вершины v, pp — предок вершины p (если таковые существуют);
•	vl — левый ребенок v, vr — правый ребенок v;
•	pl — левый ребенок p, pr — правый ребенок p.
В таком случае обмен задаётся следующими условиями:
•	Изменение предка: если p — левый ребенок вершины pp, то v становится левым ребенком pp, иначе — правым.
•	если v — левый ребенок вершины p, то:
1.	p становится левым ребенком v;
2.	vr остаётся правым ребенком v;
3.	vl становится левым ребенком p;
4.	pr остаётся правым ребенком p.

![изначальное дерево](https://github.com/Lawitz2/YandexContest2/assets/161970246/d4586bee-f259-456f-a479-d6bf75283217) ![дерево после изменений](https://github.com/Lawitz2/YandexContest2/assets/161970246/bd634d7e-c694-4c36-b1a0-9a55d443ded6)

•	аналогично, если v — правый ребенок вершины p, то:
1.	p становится правым ребенком v;
2.	vl остаётся левым ребенком v;
3.	vr становится правым ребенком p;
4.	pl остаётся левым ребенком p.

![image](https://github.com/Lawitz2/YandexContest2/assets/161970246/69dba1db-666e-4665-87de-2a6c9c291957) ![image](https://github.com/Lawitz2/YandexContest2/assets/161970246/1f3ed21b-abb3-4181-acf5-e2b9178ed81a)


Формат ввода
Первая строка содержит два целых числа N и Q (1≤N≤750;1≤Q≤106) — количество вершин в дереве и количество изменений, произошедших с деревом.
В следующей строке дано Q целых чисел v1, v2, …, vq (1≤vi≤N), где vi — номер вершины, обменявшейся местами со своим предком в i-ю ночь.
Формат вывода
В единственной строке через пробел требуется вывести номера вершин дерева после всех изменений в формате LVR, начиная с корня дерева.
Формат LVR(v) определяется рекурсивно для вершины v.
1.	если у вершины v есть левый ребенок lv, то сначала выводится всё поддерево lv в формате LVR(lv);
2.	выводится номер вершины v;
3.	если у вершины v есть правый ребенок rv, то выводится всё поддерево rv в формате LVR(rv);
Пример


Ввод
  
10 6
5 7 4 7 8 7

Вывод
  
7 10 5 2 8 4 9 1 6 3 

Объяснение примера строится из двух частей:
•	Рисунки, показывающие структуру дерева в каждом сне;
•	Текстовое пояснение к выводу на тест;
В тестовом примере дано дерево из 10 вершин:

![image](https://github.com/Lawitz2/YandexContest2/assets/161970246/998d99d5-4fc3-43d1-96cb-221a79a6aa16)

Происходит 6 изменений с деревом:
1.	вершина 5 меняется с вершиной 2:

![image](https://github.com/Lawitz2/YandexContest2/assets/161970246/889b6a5a-de69-458b-9401-b47f9ca67b15)

2.	вершина 7 меняется с вершиной 3:

![image](https://github.com/Lawitz2/YandexContest2/assets/161970246/e2919b67-a521-4b4d-92b4-eef676130db7)

3.	вершина 4 меняется с вершиной 2:

![image](https://github.com/Lawitz2/YandexContest2/assets/161970246/ea39d09e-4486-4319-9e41-4e90fdbd1563)

4.	вершина 7 меняется с вершиной 1:

![image](https://github.com/Lawitz2/YandexContest2/assets/161970246/c3d06e29-a8ab-439a-9c27-ab0a6dd60b3f)

5.	вершина 8 меняется с вершиной 2:

![image](https://github.com/Lawitz2/YandexContest2/assets/161970246/ad67c627-7c20-4240-87c9-abf3e583727e)

6.	вершина 7 ни с кем не меняется, так как она уже корень:

![image](https://github.com/Lawitz2/YandexContest2/assets/161970246/abfc3e0d-443d-47aa-a1c9-1269c3a482d6)

После всех изменений выводим получившееся дерево в формате LVR:
* Корень дерева — вершина 7;
* Выводим вершину 7;
* Выводим поддерево вершины 1 — правого ребенка вершины 7;
* Выводим поддерево вершины 5 — левого ребенка вершины 1;
* Выводим поддерево вершины 10 — левого ребенка вершины 5;
* Выводим вершину 10;
* Выводим вершину 5;
* Выводим поддерево вершины 4 — правого ребенка вершины 5;
* Выводим поддерево вершины 8 — левого ребенка вершины 4;
* Выводим поддерево вершины 2 — левого ребенка вершины 8;
* Выводим вершину 2;
* Выводим вершину 8;
* Выводим вершину 4;









