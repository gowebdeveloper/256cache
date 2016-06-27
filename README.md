# 单机极限

```
timeloveboy@timeloveboy-os:~$ ab -n 1000000 -c 1000 'http://localhost:8000/sg'
This is ApacheBench, Version 2.3 <$Revision: 1528965 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100000 requests
Completed 200000 requests
Completed 300000 requests
Completed 400000 requests
Completed 500000 requests
Completed 600000 requests
Completed 700000 requests
Completed 800000 requests
Completed 900000 requests
Completed 1000000 requests
Finished 1000000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8000

Document Path:          /sg
Document Length:        6 bytes

Concurrency Level:      1000
Time taken for tests:   65.854 seconds
Complete requests:      1000000
Failed requests:        0
Total transferred:      122000000 bytes
HTML transferred:       6000000 bytes
Requests per second:    15185.05 [#/sec] (mean)
Time per request:       65.854 [ms] (mean)
Time per request:       0.066 [ms] (mean, across all concurrent requests)
Transfer rate:          1809.16 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:       11   32   9.7     32      93
Processing:     5   34   9.8     34     105
Waiting:        2   19  10.3     17      72
Total:         46   66   3.9     66     156

Percentage of the requests served within a certain time (ms)
  50%     66
  66%     67
  75%     67
  80%     67
  90%     68
  95%     72
  98%     79
  99%     80
 100%    156 (longest request)

|   |
|---|
|   |

```