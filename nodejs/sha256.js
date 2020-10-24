const express = require('express');
const app = express();
const cors = require('cors');
const port = 8080;

app.use(cors());
app.get('/nodejs/sha256', (req, res) => {
  res.json({ result: 12 });
  console.log(res);
})

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
})