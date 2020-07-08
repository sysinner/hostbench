# hostbench



## CPU Benchmark

Install:

``` shell
go get -u github.com/sysinner/hostbench/cpubench

go install github.com/sysinner/hostbench/cpubench
```

Usage:

``` shell
[root@localhost ~]# cpubench
           CPU Model  Intel(R) Xeon(R) CPU E5-2680 v3 @ 2.50GHz
                Arch  amd64
                  OS  linux
                Time  65 s
               Score  1518
```

Partial test results:

| CPU MODEL                                 | GHz  | SCORE | COMMENT             |
|-------------------------------------------|------|-------|---------------------|
| Intel Xeon Processor (Cascadelake)        | 2.29 | 2200  | ucloud-bj           |
| Intel(R) Core(TM) i5-10210U CPU @ 1.60GHz | 4.20 | 2815  | lenovo-thinkbook-14 |
| Intel(R) Xeon(R) CPU E5-2680 v3 @ 2.50GHz | 2.50 | 1518  | linode-fremont      |
| Intel(R) Xeon(R) Gold 6140 CPU @ 2.30GHz  | 2.29 | 1912  | digitalocean-sfo1   |


