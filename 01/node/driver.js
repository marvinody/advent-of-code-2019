const util = require('util');
const fs = require('fs');
const path = require('path')

const file = path.join(__dirname, '..', 'input')
const readFile = util.promisify(fs.readFile);


// Change this each time you make a new challenge probably

(async function () {
  // this is probably always the same
  const input = await readFile(file, 'utf8')

  // the rest changes
  const masses = input.split('\n').filter(l => l.length)
  const { findFuelForMass } = require('./code')

  const totalFuel = masses.reduce((acc, cur) =>
    acc + findFuelForMass(cur)
    , 0)

  console.log({ totalFuel })

})()
