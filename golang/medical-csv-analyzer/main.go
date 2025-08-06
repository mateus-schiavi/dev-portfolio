package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

// Converte coluna para slice de float64, ignorando o cabeçalho
func convertColumnToFloats(records [][]string, colIndex int) ([]float64, error) {
	var result []float64
	for i, row := range records {
		if i == 0 {
			// Ignora cabeçalho
			continue
		}
		if colIndex >= len(row) {
			return nil, fmt.Errorf("linha %d não tem coluna %d", i, colIndex)
		}
		val, err := strconv.ParseFloat(row[colIndex], 64)
		if err != nil {
			log.Printf("Aviso: linha %d coluna %d não é número válido: %v\n", i, colIndex, err)
			continue
		}
		result = append(result, val)
	}
	return result, nil
}

func mean(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

func median(numbers []float64) float64 {
	n := len(numbers)
	if n == 0 {
		return 0
	}
	sorted := make([]float64, n)
	copy(sorted, numbers)
	sort.Float64s(sorted)
	mid := n / 2
	if n%2 == 0 {
		return (sorted[mid-1] + sorted[mid]) / 2
	}
	return sorted[mid]
}

func stdDev(numbers []float64) float64 {
	n := len(numbers)
	if n == 0 {
		return 0
	}
	m := mean(numbers)
	sumSquares := 0.0
	for _, num := range numbers {
		sumSquares += (num - m) * (num - m)
	}
	variance := sumSquares / float64(n)
	return math.Sqrt(variance)
}

func detectOutliers(numbers []float64) []float64 {
	var outliers []float64
	m := mean(numbers)
	sd := stdDev(numbers)
	lower := m - 2*sd
	upper := m + 2*sd
	for _, num := range numbers {
		if num < lower || num > upper {
			outliers = append(outliers, num)
		}
	}
	return outliers
}

func main() {
	filePath := flag.String("file", "", "Caminho para o arquivo CSV com dados médicos")
	colIndex := flag.Int("col", 1, "Índice da coluna numérica para análise (começa em 0)")
	flag.Parse()

	if *filePath == "" {
		log.Fatal("Por favor, informe o caminho do arquivo CSV com -file=data.csv")
	}

	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatalf("Erro ao abrir arquivo: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Erro ao ler CSV: %v", err)
	}

	fmt.Printf("Arquivo CSV carregado com %d linhas\n", len(records))

	numbers, err := convertColumnToFloats(records, *colIndex)
	if err != nil {
		log.Fatalf("Erro ao converter coluna: %v", err)
	}

	fmt.Printf("Analisando coluna %d com %d valores válidos\n\n", *colIndex, len(numbers))
	fmt.Printf("Média: %.2f\n", mean(numbers))
	fmt.Printf("Mediana: %.2f\n", median(numbers))
	fmt.Printf("Desvio Padrão: %.2f\n", stdDev(numbers))

	outliers := detectOutliers(numbers)
	if len(outliers) > 0 {
		fmt.Printf("Valores fora do padrão (outliers): %v\n", outliers)
	} else {
		fmt.Println("Nenhum outlier detectado.")
	}
}

