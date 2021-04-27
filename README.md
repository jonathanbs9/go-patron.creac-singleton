# go-patron.creac-singleton

## Singleton

Singleton es un patrón de diseño creacional que nos permite asegurarnos de que una clase tenga una única instancia, a la vez que proporciona un punto de acceso global a dicha instancia.

![image](https://user-images.githubusercontent.com/32901911/116247545-58d50380-a741-11eb-804a-a5a3e09d81a6.png)

## Aplicabilidad

:round_pushpin: Utiliza el patrón Singleton cuando una clase de tu programa tan solo deba tener una instancia disponible para todos los clientes; por ejemplo, un único objeto de base de datos compartido por distintas partes del programa.

:round_pushpin: El patrón Singleton deshabilita el resto de las maneras de crear objetos de una clase, excepto el método especial de creación. Este método crea un nuevo objeto, o bien devuelve uno existente si ya ha sido creado.

## PROS y Contras

✔️ Puedes tener la certeza de que una clase tiene una única instancia.

✔️ Obtienes un punto de acceso global a dicha instancia.

✔️ El objeto Singleton solo se inicializa cuando se requiere por primera vez.


❌ Vulnera el Principio de responsabilidad única. El patrón resuelve dos problemas al mismo tiempo.

❌ El patrón Singleton puede enmascarar un mal diseño, por ejemplo, cuando los componentes del programa saben demasiado los unos sobre los otros.

❌ El patrón requiere de un tratamiento especial en un entorno con múltiples hilos de ejecución, para que varios hilos no creen un objeto Singleton varias veces.

❌ Puede resultar complicado realizar la prueba unitaria del código cliente del Singleton porque muchos frameworks de prueba dependen de la herencia a la hora de crear objetos simulados (mock objects). Debido a que la clase Singleton es privada y en la mayoría de los lenguajes resulta imposible sobrescribir métodos estáticos, tendrás que pensar en una manera original de simular el Singleton. O, simplemente, no escribas las pruebas. O no utilices el patrón Singleton.
