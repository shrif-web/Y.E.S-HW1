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
  if (isNaN(Number(req.body.first) + Number(req.body.second))) {
    res.status(401).send("input integer number please")
    return
  }
  res.send(sha256((Number(req.body.first) + Number(req.body.second)).toString()))
});

app.get('/nodejs/write', function (req, res) {
  console.log("request line number:", req.query.number);
  console.log("path:"+process.cwd())
  if (!(req.query.number > 0 && req.query.number <= 100)) {
    res.status(401).send("input number range should be between 1-100")
    return
  }
  var str = fs.readFileSync('my_text.txt', 'utf8')
  res.send(str.split("\n")[req.query.number - 1])
});


app.listen(port, () => {
  console.log(`nodejs app listening at ${domain}:${port}`)
});
