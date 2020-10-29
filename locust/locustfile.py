import time
import random
from locust import HttpUser, task, constant , tag

max_line=3
class GoUser(HttpUser):
    wait_time = constant(1)

    @tag('go')
    @task
    def sha(self):
        self.client.post("/go/sha256", json={"first":"2", "second":"3"})

    @tag('go')
    @task
    def write(self):
        self.client.get("/go/write?number={}".format(random.randint(1,max_line)),name="/go/write")


class NodeUser(HttpUser):
    wait_time = constant(1)
    @tag('node')
    @task
    def sha(self):
        self.client.post("/nodejs/sha256", json={"first":"2", "second":"3"})

    @tag('node')
    @task
    def write(self):
        self.client.get("/nodejs/write?number={}".format(random.randint(1,max_line)),name="/node/write")

