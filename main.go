package main

import(
  f "fmt"
  t "time"
  r "math/rand"
) 

func fuelGauge(fuel int){ 
  f.Println("Fuel left",fuel)
}

func calculateFuel(planet string) int{
  var fuelNeeded int
  switch planet {
    case "Venus":
      fuelNeeded = 300000
    case "Mercury":
      fuelNeeded = 500000
    case "Mars":
      fuelNeeded = 700000
    case "Jupiter":
      fuelNeeded = 15000000
    case "Saturn":
      fuelNeeded = 185000000
    default:
      fuelNeeded = 0
  }
  return fuelNeeded
}

func greenPlanet(planet string) {
  f.Println ("You reached your destination.")
  }

func cantFly() {

}

func flyToPlanet(planet string, fuel int) int{
  var fuelCost int
  var fuelRemaining int
  fuelRemaining = fuel
  fuelCost = calculateFuel(planet)
  if fuelRemaining > fuelCost{
    greenPlanet(planet) 
    } else {
         defer f.Println("We do not have the available fuel to fly there.")
      }
      return fuel
    }

func getDestination(planet string) string{
  return planet
}

func returnToHome(planet string, fuel int) {
  gasInExtraTanks := r.Intn(500000)
  chanceOfGasOnPlanet := r.Intn(100)
  waitTime := r.Intn(100)
  f.Println("Fuel needed to fly back home",calculateFuel(planet))
  if fuel > calculateFuel(planet) {
  f.Println("We still have enough gas to fly back to home!")
  } else if (gasInExtraTanks + fuel) < calculateFuel(planet) {
    fuelGauge(fuel)
    f.Print("Gas in extra tanks: ",gasInExtraTanks)
    f.Println("\nWe have to refill after the refill we are ready to take off!")
    fuel += gasInExtraTanks
    fuelGauge(fuel)
  } 
  if gasInExtraTanks < calculateFuel(planet) && fuel < calculateFuel(planet) {
    f.Printf("We need to find new Gas, let's search it on the Planet we are on.")
    f.Print("The chance we find a gas source is: ", chanceOfGasOnPlanet)
    f.Print("%\n")
    } 
     if fuel < calculateFuel(planet) && chanceOfGasOnPlanet > 50 {
        f.Println("Let's find the Gas")
        gasFound := r.Intn(1000000)
        f.Println("We found",gasFound,"l of gas")
        fuel += gasFound
        fuelGauge(fuel)
        } 
        if fuel >= calculateFuel(planet) {
          f.Println("Ready for takeoff!")
          fuelGauge(fuel-calculateFuel(planet))
        } else if waitTime != 0 {
         for waitTime <= chanceOfGasOnPlanet*r.Intn(10) {
           waitTime = r.Intn(1000)
         } 
                   f.Println("We have to mine for more fuel waiting time aprox.:", waitTime)
        }

        }
      

func main() {
  r.Seed(t.Now().UnixNano())
  // Test your functions!

  // Create `planetChoice` and `fuel`
  var fuel int
  var choice string
  fuel = 100000
  var planetChoice string
  planetChoice = "Mars" 
  flyToPlanet(planetChoice, fuel)
  f.Println("You wanna stay (y or n) ?")
  choice = "n"
  if choice == "y" {
    return
  } else if choice == "n" {
      returnToHome(planetChoice, fuel)
    } else {
        return
      }
  }
