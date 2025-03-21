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

| NAME                      | SCORE | MODEL                                          | GHz  | Arch  | DATE       |
|---------------------------|-------|------------------------------------------------|------|-------|------------|
| workstation-n3            | 8060  | AMD Ryzen 5 5600G                              | 3.90 | x64   | 2023-06-21 |
| aliyun-ecs.g7a.large-c2m8 | 5809  | AMD EPYC 7T83 64-Core Processor                | 2.55 | x64   | 2022-11-04 |
| tencent-cloud-sa3-c2m4    | 5788  | AMD EPYC 7K83 64-Core Processor                | 2.55 | x64   | 2022-11-04 |
| aliyun-ecs.g7.large-c2m8  | 5740  | Intel(R) Xeon(R) Platinum 8369B CPU @ 2.70GHz  | 2.70 | x64   | 2022-11-04 |
| aliyun-ecs.c7.large-c2m8  | 5739  | Intel(R) Xeon(R) Platinum 8369B CPU @ 2.70GHz  | 2.70 | x64   | 2022-11-04 |
| aliyun-ecs.g8y.large-c2m8 | 5624  | Yitian 710                                     | 2.75 | arm64 | 2023-01-04 |
| tencent-cloud-s6-c2m2     | 5415  | Intel(R) Xeon(R) Platinum 8374C CPU @ 2.70GHz  | 2.70 | x64   | 2022-11-04 |
| linode-2g-fremont         | 5011  | AMD EPYC 7713 64-Core Processor                | 2.00 | x64   | 2024-04-08 |
| digitalocean-gpu-h100x1   | 4978  | Intel(R) Xeon(R) Platinum 8468                 | 2.10 | x64   | 2025-03-20 |
| tencent-cloud-sa2-c1m2-bj | 4934  | AMD EPYC 7K62 48-Core Processor                | 2.60 | x64   | 2020-07-22 |
| aliyun-ecs.g7r.large      | 4006  | Ampere(R) Altra(R) CPU @ 2.80GHz               | 2.80 | arm64 | 2022-11-04 |
| macbook-pro-2021          | 3697  | Apple M1                                       | 2.40 | arm64 | 2022-11-04 |
| workstation-n2            | 2913  | Intel(R) Core(TM) i7-7700 CPU @ 3.60GHz        | 4.20 | x64   | 2020-07-24 |
| lenovo-thinkbook-14       | 2815  | Intel(R) Core(TM) i5-10210U CPU @ 1.60GHz      | 4.20 | x64   | 2020-07-08 |
| tencent-cloud-c3-c4m8-bj  | 2693  | Intel(R) Xeon(R) Gold 6146 CPU @ 3.20GHz       | 3.19 | x64   | 2020-08-07 |
| ep-host-e5-2637           | 2536  | Intel(R) Xeon(R) CPU E5-2637 v2 @ 3.50GHz      | 3.80 | x64   | 2020-08-01 |
| aliyun-hfc6-c2m4-hz       | 2431  | Intel(R) Xeon(R) Platinum 8269CY CPU @ 2.50GHz | 3.10 | x64   | 2020-08-07 |
| vultr-hf-c1m1-frankfurt   | 2324  | Intel Core Processor (Skylake, IBRS)           | 3.79 | x64   | 2020-10-21 |
| vultr-c1m2-frankfurt      | 2297  | Intel Xeon Processor (Skylake, IBRS)           | 2.59 | x64   | 2020-10-27 |
| huawei-cloud-hecs-c4m8-bj | 2265  | Intel(R) Xeon(R) Gold 6278C CPU @ 2.60GHz      | 2.60 | x64   | 2020-09-15 |
| ucloud-kj-o-bj            | 2200  | Intel Xeon Processor (Cascadelake)             | 2.29 | x64   | 2020-07-08 |
| aliyun-c6e-c2m4-hz        | 2174  | Intel(R) Xeon(R) Platinum 8269CY CPU @ 2.50GHz | 2.50 | x64   | 2020-08-07 |
| aliyun-g6e-c2m8-hz        | 2174  | Intel(R) Xeon(R) Platinum 8269CY CPU @ 2.50GHz | 2.50 | x64   | 2020-08-07 |
| aliyun-s6-c1m2-hz         | 2166  | Intel(R) Xeon(R) Platinum 8269CY CPU @ 2.50GHz | 2.50 | x64   | 2020-08-07 |
| digitalocean-1g-sfo1      | 1912  | Intel(R) Xeon(R) Gold 6140 CPU @ 2.30GHz       | 2.29 | x64   | 2020-07-08 |
| ep-host-e5-2620           | 1639  | Intel(R) Xeon(R) CPU E5-2620 v2 @ 2.10GHz      | 2.60 | x64   | 2020-08-01 |
| linode-1g-fremont         | 1593  | Intel(R) Xeon(R) Gold 6148 CPU @ 2.40GHz       | 2.40 | x64   | 2020-07-26 |
| linode-1g-tokyo2          | 1531  | Intel(R) Xeon(R) CPU E5-2697 v4 @ 2.30GHz      | 2.30 | x64   | 2020-07-24 |
| linode-2g-fremont         | 1518  | Intel(R) Xeon(R) CPU E5-2680 v3 @ 2.50GHz      | 2.50 | x64   | 2020-07-08 |

