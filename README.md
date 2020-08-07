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

Partial test results (1x CPU core):

| MODEL                                          | GHz  | SCORE | COMMENT                   |
|------------------------------------------------|------|-------|---------------------------|
| AMD EPYC 7K62 48-Core Processor                | 2.60 | 4934  | tencent-cloud-sa2-c1m2-bj |
| Intel(R) Core(TM) i7-7700 CPU @ 3.60GHz        | 4.20 | 2913  | workstation-n2            |
| Intel(R) Core(TM) i5-10210U CPU @ 1.60GHz      | 4.20 | 2815  | lenovo-thinkbook-14       |
| Intel(R) Xeon(R) Gold 6146 CPU @ 3.20GHz       | 3.19 | 2693  | tencent-cloud-c3-c4m8-bj  |
| Intel(R) Xeon(R) CPU E5-2637 v2 @ 3.50GHz      | 3.80 | 2536  | ep-host-e5-2637           |
| Intel(R) Xeon(R) Platinum 8269CY CPU @ 2.50GHz | 3.10 | 2431  | aliyun-hfc6-c2m4-hz       |
| Intel Xeon Processor (Cascadelake)             | 2.29 | 2200  | ucloud-kj-o-bj            |
| Intel(R) Xeon(R) Platinum 8269CY CPU @ 2.50GHz | 2.50 | 2174  | aliyun-c6e-c2m4-hz        |
| Intel(R) Xeon(R) Platinum 8269CY CPU @ 2.50GHz | 2.50 | 2174  | aliyun-g6e-c2m8-hz        |
| Intel(R) Xeon(R) Platinum 8269CY CPU @ 2.50GHz | 2.50 | 2166  | aliyun-s6-c1m2-hz         |
| Intel(R) Xeon(R) Gold 6140 CPU @ 2.30GHz       | 2.29 | 1912  | digitalocean-1g-sfo1      |
| Intel(R) Xeon(R) CPU E5-2620 v2 @ 2.10GHz      | 2.60 | 1639  | ep-host-e5-2620           |
| Intel(R) Xeon(R) Gold 6148 CPU @ 2.40GHz       | 2.40 | 1593  | linode-1g-fremont         |
| Intel(R) Xeon(R) CPU E5-2697 v4 @ 2.30GHz      | 2.30 | 1531  | linode-1g-tokyo2          |
| Intel(R) Xeon(R) CPU E5-2680 v3 @ 2.50GHz      | 2.50 | 1518  | linode-2g-fremont         |

