# Locust

Locust is a powerful tool for testing load of a website. We can define user Behavior in python code and swarm our system with millions of simultaneous users



## Install 
For installing locust, you should use a simple command using pip3:

> pip3 install locust

## Writing Test
For testing, we can simply create a python file called ```locustfile.py``` and we do the configuration for load test.

## Running

There are many ways to run the test. We used master-worker mode to run the tests.

So first run the command below on your master node:

```locust --master```

Then run next command on all worker nodes (master node can also be used as worker too):

```locust--worker --master-host=<master locust ip>```

After running commands, locust starts up a local web server on which you can visit in your browser. In this website you can enter the number of **users** and **hatch rate** for test.
In this page we can see the **requests per second**, **failure percentage** and other useful information.

We did 2 tests, one with 100000 users and 100 users spawned per second. This is the result for the first experiment :
```
 Name                                                          # reqs      # fails  |     Avg     Min     Max  Median  |   req/s failures/s
--------------------------------------------------------------------------------------------------------------------------------------------
 POST /go/sha256                                                27190   894(3.29%)  |    1465       2   70023    1300  |  139.80    4.60
 GET /go/write                                                  27222   885(3.25%)  |    1542       2   69349    1300  |  139.96    4.55
 GET /node/write                                                27111   831(3.07%)  |    1536       2   72830    1300  |  139.39    4.27
 POST /nodejs/sha256                                            26944   794(2.95%)  |    1495       2   69626    1300  |  138.53    4.08
--------------------------------------------------------------------------------------------------------------------------------------------
 Aggregated                                                    108467  3404(3.14%)  |    1509       2   72830    1300  |  557.67   17.50
 ```
And this is result for the second test with 100000 users and 1000 user spawned per request:
```
 Name                                                          # reqs      # fails  |     Avg     Min     Max  Median  |   req/s failures/s
--------------------------------------------------------------------------------------------------------------------------------------------
 POST /go/sha256                                                83277  8012(9.62%)  |    9547       3  174345    1600  |  172.66   16.61
 GET /go/write                                                  82829  7739(9.34%)  |    9295       3  170696    1600  |  171.73   16.05
 GET /node/write                                                81364  8125(9.99%)  |    9753       2  170623    1600  |  168.70   16.85
 POST /nodejs/sha256                                            81052  7864(9.70%)  |    9511       2  170636    1600  |  168.05   16.30
--------------------------------------------------------------------------------------------------------------------------------------------
 Aggregated                                                    328522 31740(9.66%)  |    9525       2  174345    1600  |  681.14   65.81
 ```
In the second test, when number of users reached 100000, failure percentage went up to 4%.
But as the test users doubled, failure percentage reached 10%.

### hardware configuration:
>number of workers: 3
>
>workers hardware: 2 cores of CPU with 2 GB of ram
>
>**Note:** yes.io server was also configured with 2 gigabytes of RAM and 2 cores of CPU
