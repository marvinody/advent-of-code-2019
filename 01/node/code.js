const findFuelForMass = mass => {
  if (mass < 0) return 0
  const fuelNeededForMass = Math.floor(mass / 3) - 2
  const fuelNeededForFuel = Math.max(0, findFuelForMass(fuelNeededForMass))
  return fuelNeededForMass + fuelNeededForFuel

}

module.exports = {
  findFuelForMass
}
