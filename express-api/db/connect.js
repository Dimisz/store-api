const mongoose = require('mongoose')

const connectDB = (dbURL) => {
  return mongoose.connect(dbURL)
    .then(() => console.log('connected to the db...'));
}

module.exports = connectDB