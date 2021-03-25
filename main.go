package goarea // mesmo nome da pasta que está no src
import "math"

// Pi é uma proporção numérica definida pela relação entre
// o perímetro de uma circunferência e seu diâmetro

const Pi = 3.1416 // Constante pública (nome começa com letra maiúscula) é nível de pacote

// Circ é responsável por calcular a área da circunferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é responsável por calcular a área de um retângulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Funçao só pra testar, ela não é visível por causa do _ (underline)
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
