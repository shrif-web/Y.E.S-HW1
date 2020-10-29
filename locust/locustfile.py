import time
from locust import HttpUser, task, constant

class GoUser(HttpUser):
    wait_time = constant(0.5)
    @task
    def sha(self):
        self.client.post("/go/sha256", json={"first":"2", "second":"3"})
    @task
    def write(self):
        self.client.get("/go/write?number=1")


class NodeUser(HttpUser):
    wait_time = constant(0.5)
    @task
    def sha(self):
        self.client.post("/nodejs/sha256", json={"first":"2", "second":"3"})
    @task
    def write(self):
        self.client.get("/nodejs/write?number=1")

