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
           CPU Model  AMD EPYC 7K62 48-Core Processor
                Arch  amd64
                  OS  linux
                Time  20 s
               Score  4934
```

Partial test results:

| MODEL                                     | GHz  | SCORE | COMMENT              |
|-------------------------------------------|------|-------|----------------------|
| AMD EPYC 7K62 48-Core Processor           | 2.60 | 4934  | tencent-cloud-sa2-bj |
| Intel(R) Core(TM) i7-7700 CPU @ 3.60GHz   | 4.20 | 2913  | workstation-n2       |
| Intel(R) Core(TM) i5-10210U CPU @ 1.60GHz | 4.20 | 2815  | lenovo-thinkbook-14  |
| Intel Xeon Processor (Cascadelake)        | 2.29 | 2200  | ucloud-bj            |
| Intel(R) Xeon(R) Gold 6140 CPU @ 2.30GHz  | 2.29 | 1912  | digitalocean-sfo1    |
| Intel(R) Xeon(R) CPU E5-2680 v3 @ 2.50GHz | 2.50 | 1518  | linode-fremont       |

