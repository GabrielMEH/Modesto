Como rodar o programa?
1. Instale Go usando o guia no site https://go.dev/doc/install
2. acesse o repositório do projeto pelo shell/cmd/editor de código-fonte escolhido
3. digite o comando "go run ." para executar o arquivo codigo.go
4. digite a entrada do exercicio

Exercicio:
C de Chambinho

Leo ( conhecido pelos amigos de Chambinho)  está numa festa do gremio  estudantil do IFSP - SaLto  e pretende se divertir um pouco tomando o máximo de copos de chambinho que conseguir.

Na mesa central do salão existe uma fileira de N copos, o copo i tem Ci mililitros de chambinho .

Pela tradição da universidade onde ele estuda, Chambinho deve beber de uma fileira continua de copos de chambinho , um atrás do outro (apenas uma vez).  Ou seja, Leo deve escolher dois inteiros l e r e beber dos copos l, l + 1, ..., r, tomando Cl, Cl+1,..., Cr milímetros de chambinho .
Chambinho já se conhece muito bem, e sabe que se ele beber mais de D  mililitros de chambinho , ele vai passar mal, ou seja, ele pode beber no máximo D mililitros de chambinho .

Sabendo disso, Chambinho quer beber a maior quantidade de copos possíveis sem passar mal e seguindo a tradição da sua universidade, mas ele está com preguiça de pensar muito, então pediu sua ajuda para saber qual a maior fileira de copos que ele pode beber sem ultrapassa D mililitros.

É garantido que Chambinho consegue beber pelo menos um copo.

Entrada
A primeira linha contém 2 inteiros  N e D, que representam a quantidade de copos e o quanto de chambinho  Chambinho pode tomar sem passar mal.

A segunda linha contém N  inteiros C1, C2, ..., CN que representam quantos milimitros de chambinho tem em cada copo.


Saída
Imprima um único inteiro: a quantidade máxima de copos que Leo consegue beber sem passar mal.


Restrições
1≤N≤10^5

1≤Ci≤D≤10^9

Informações Sobre Pontuação
Para um conjunto valendo 20 pontos: N≤10^2≤eD≤10^5
Para um conjunto valendo 40 pontos: N≤5∗10^3eD≤10^5
Para um conjunto valendo 40 pontos: 10^5

Entradas de teste:
5 10
2 4 3 2 3
resultado esperado: 3

10 20
2 5 4 2 6 3 7 7 6 4
resultado esperado: 5

7 10
2 2 2 2 4 2 4
resultado esperado: 4