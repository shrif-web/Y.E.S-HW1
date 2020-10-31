const sha256 = require("js-sha256");
const express = require("express");
const fs = require("fs");
const app = express();
const cors = require("cors");
const bodyParser = require("body-parser");
const port = 8080;
const domain = "http://localhost";

app.use(cors());
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

app.post("/nodejs/sha256", function(req, res) {
  console.log("request body:", req.body);
  if (isNaN(Number(req.body.first) + Number(req.body.second))) {
    res.status(401).send("input integer number please");
    return;
  }
  res.send(
    sha256((Number(req.body.first) + Number(req.body.second)).toString())
  );
});

app.get("/nodejs/write", function(req, res) {
  console.log("request line number:", req.query.number);
  console.log("path:" + process.cwd());
  if (!(req.query.number > 0 && req.query.number <= 100)) {
    res.status(401).send("input number range should be between 1-100");
    return;
  }
  var str = fs.readFileSync("my_text.txt", "utf8");
  res.send(str.split("\n")[req.query.number - 1]);
});

app.use(function(req, res, next) {
  // console.log("request:", req.body, req.url);
  // res.status(404).send("Unable to find the requested resource!");
  console.log("here", req)
  if (req.accepts("html")) {
    console.log("yes")
    res.redirect(`/helloworld/404.html`);
  } else if (req.accepts("json")) {
    res.status(404);
    console.log("here2")
    // res.json({ message: `${req.method}: ${fullUrl(req)} is not supported!` });
  } else {
    res.status(404);
    res.type("txt").send("Not found");
  }
});

app.listen(port, () => {
  console.log(`nodejs app listening at ${domain}:${port}`);
});
