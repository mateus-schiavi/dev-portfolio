package main

import (
    "github.com/fogleman/gg"
)

func main() {
    const W, H = 400, 400
    dc := gg.NewContext(W, H)

    dc.SetRGB(1, 1, 1) // Fundo branco
    dc.Clear()

    dc.SetRGB(0, 0, 1) // Azul
    dc.DrawCircle(W/2, H/2, 100)
    dc.Fill()

    dc.SetRGB(1, 0, 0) // Vermelho
    dc.DrawStringAnchored("Olá, Go!", W/2, H/2, 0.5, 0.5)

    dc.SavePNG("output.png")
}
