const sha256 = require('js-sha256');
const express = require('express');
const app = express();
const cors = require('cors');
const bodyParser = require('body-parser');
const port = 8080;
const domain = "http://localhost"

app.use(cors());
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

app.post('/nodejs/sha256',function(req,res){
  console.log("request body:",req.body);
  res.json({
    result: sha256((Number(req.body.first)+Number(req.body.second)).toString())
  });
});

app.listen(port, () => {
  console.log(`nodejs app listening at ${domain}:${port}`)
});