package main

import "fmt"

type Animal struct {
	Id     int
	Nombre string
}

// Tipos de apuntadores
// *Animal = Almacenamiento
// &Animal = Direccion memoria

func NewAnimal(id int, nombre string) *Animal {
	return &Animal{
		Id:     id,
		Nombre: nombre,
	}
}

func (clase *Animal) SetName(nombre string) {
	clase.Nombre = nombre
}

func (clase *Animal) GetNombre() string {
	fmt.Printf("el nombre es: %s", clase.Nombre)
	return clase.Nombre
}

type Ave struct {
	Animal
}

func NewAve(id int, nombre string) *Ave {
	return &Ave{Animal{id, nombre}}
}

func (clase Ave) SoyUnAve() string {
	return "SoyUnAve"
}

func (clase Ave) HazAlgo() string {
	return "hahahahaha"
}

// Polyformismo
type GrasnarComoAveInterfaz interface {
	HazAlgo() string
}

func GrasnarComoAve(aliasDeInterfaz GrasnarComoAveInterfaz) string {
	return aliasDeInterfaz.HazAlgo()
}

func main() {
	animal := NewAnimal(1, "Jojo")
	fmt.Println(animal.Id)
	fmt.Println(animal.Nombre)
	fmt.Println(animal.GetNombre())
	animal.SetName("Malva")
	fmt.Println(animal.Nombre)
	fmt.Println(animal.GetNombre())

	ave := NewAve(1, "Mojo")
	fmt.Println(ave.Id)
	fmt.Println(ave.Nombre)
	fmt.Println(ave.GetNombre())
	ave.SetName("Jojo")
	fmt.Println(ave.Nombre)
	fmt.Println(ave.SoyUnAve())
	fmt.Println(ave.GetNombre())

}
