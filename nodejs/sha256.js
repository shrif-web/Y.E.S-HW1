const express = require('express')
const app = express()
const port = 8080

app.get('/nodejs/sha256',(req,res) =>{
    res.send("shallum")
})

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
})