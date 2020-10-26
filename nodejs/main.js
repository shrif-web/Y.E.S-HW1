const sha256 = require('js-sha256');
const express = require('express');
const fs = require('fs');
const app = express();
const cors = require('cors');
const bodyParser = require('body-parser');
const port = 8080;
const domain = "http://localhost"

app.use(cors());
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

app.post('/nodejs/sha256', function (req, res) {
  console.log("request body:", req.body);
  if (isNaN(Number(req.body.first) + Number(req.body.second))){
    res.json({
      error: "input integer number please"
    });
    return
  }
  console.log(Number(req.body.second))
  res.json({
    result: sha256((Number(req.body.first) + Number(req.body.second)).toString())
  });
});

app.get('/nodejs/write', function (req, res) {
  console.log("request body:", req.body);
  if (!(req.body.number > 0 && req.body.number <= 100)) {
    res.json({
      error: "input number range should be between 1-100"
    });
    return
  }
  var kharkir = fs.readFileSync('my_text.txt', 'utf8')
  res.json({
    result: kharkir.split("\n")[req.body.number - 1]
  });
});


app.listen(port, () => {
  console.log(`nodejs app listening at ${domain}:${port}`)
});
