---
menu:
    after:
        name: graph
        weight: 1
title: Построение графа
---
{{< mermaid >}}
graph LR
  2{Node_2, rhombus}
  2 --> 7
  7(Node_7, round-rect)
  7 --> 5
  5[Node_5, square]
  5 --> 2
  5 --> 3
  3((Node_3, circle))
  3 --> 4
  4[Node_4, square]
  4 --> 1
  1((Node_1, circle))
  4 --> 5
  4 --> 1
  4 --> 1
  3 --> 7
  3 --> 7
  3 --> 1
  3 --> 2
  3 --> 6
  6(Node_6, round-rect)
  6 --> 5
  6 --> 7
  6 --> 4
  6 --> 3
  6 --> 1
  6 --> 3
  5 --> 1
  5 --> 3
  5 --> 3
  5 --> 3
  7 --> 4
  7 --> 3
  7 --> 4
  2 --> 6
  2 --> 1
{{< /mermaid >}}