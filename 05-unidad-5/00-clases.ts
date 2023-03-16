// POO en Typescript
class Animal {
   public id : number;
   public nombre : string;

   constructor(id: number, nombre: string){
        this.id = id;
        this.nombre = nombre;
   }

   public setNombre(nombre: string){
    this.nombre = nombre;
   }
   public getNombre(): string{
    return this.nombre;
   }
}

class Ave extends Animal {
    
}

const animal = new Animal(1,"Jojo");

console.log(animal.id);

console.log(animal.nombre);
animal.setNombre("Malva");
console.log(animal.nombre);
console.log(animal.getNombre());


