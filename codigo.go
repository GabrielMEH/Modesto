package main

//importar pacote fmt, incluido na biblioteca padrão do GO
import (
	"fmt"
)

func main() {
	var n, d, contMax, result, copos int

	//fmt.Scanf("%d%d", &n, &d)
	fmt.Scan(&n, &d)
	list := make([]int, n)
	fmt.Scanf("%d", list[0])

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &list[i])
	}

	//fmt.Printf("%d %d\n", n, d)
	fmt.Println(n, d)

	for i := 0; i < n; i++ {
		//fmt.Printf("%d ", list[i])
		fmt.Print(list[i], " ")
	}
	fmt.Println()

	// passa pela lista de copos de chambinho
	for i := 0; i < n; i++ {
		//atribui zero a quantidade de copos e contador de chambinho
		copos = 0
		contMax = 0
		fmt.Print("Sequência ", i+1, ": ")
		// passa pela lista a partir de i contando limite
		for j := i; j < n; j++ {
			// se limite não estourar com próximo copo
			if (contMax + list[j]) <= d {
				contMax += list[j]
				copos++
				fmt.Print(list[j], " ")
			} else {
				//se próximo copo não pode ser consumido, quebra a sequência
				// se resultado = 0 ou resultado atual < copos
				if result == 0 || result < copos {
					result = copos
				}
				break
			}
		}
		//fmt.Printf("%d", copos)
		fmt.Println()
		fmt.Println("Tamanho da sequência: ", copos)
		fmt.Println("Chambinho consumido: ", contMax)
		fmt.Println("-----------")
	}
	//fmt.Printf("resultado: %d\n", result)
	fmt.Print("resultado: ", result)
}
