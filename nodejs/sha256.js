const express = require('express');
const app = express();
const cors = require('cors');
const bodyParser = require('body-parser');
const port = 8080;

app.use(cors());
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

app.post('/nodejs/sha256',function(req,res){
  res.json({
    result: Number(req.body.f)+Number(req.body.s)
  });
});

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
});