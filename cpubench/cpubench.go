package main

import (
	"fmt"
	"log"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/alexeyco/simpletable"
	"github.com/hooto/hflag4g/hflag"
	"github.com/hooto/htoml4g/htoml"
	"github.com/shirou/gopsutil/v3/cpu"
)

var (
	CpuBenchBase2019 [2]int64 = [2]int64{1 << 16, 100e12}
)

type CpuBenchList struct {
	Items []*CpuBenchItem `toml:"items"`
}

type CpuBenchItem struct {
	Name      string       `toml:"name"`
	ModelName string       `toml:"model_name"`
	Score     int64        `toml:"score"`
	TestTime  string       `toml:"test_time"`
	CpuInfo   cpu.InfoStat `toml:"cpu_info"`
}

func main() {

	runtime.GOMAXPROCS(1)

	var benchResult *CpuBenchItem

	if _, ok := hflag.ValueOK("bench-off"); !ok {

		cpuInfo, err := cpu.Info()
		if err != nil {
			log.Fatal(err)
		}

		infos := []string{}
		infos = append(infos, []string{
			"Model", cpuInfo[0].ModelName,
			"Arch", runtime.GOARCH,
			"OS", runtime.GOOS,
		}...)

		sec, score := CpuBench(CpuBenchBase2019)

		infos = append(infos, []string{
			"Time", fmt.Sprintf("%d s", sec/1e9),
			"Score", fmt.Sprintf("%d", score),
		}...)

		for i := 1; i < len(infos); i += 2 {
			fmt.Printf("%20s  %s\n", infos[i-1], infos[i])
		}

		benchResult = &CpuBenchItem{
			Name:      hflag.Value("name").String(),
			ModelName: cpuInfo[0].ModelName,
			Score:     score,
			TestTime:  time.Now().Format("2006-01-02"),
			CpuInfo:   cpuInfo[0],
		}
	}

	if v, ok := hflag.ValueOK("table-list"); ok {

		var ls CpuBenchList
		cfgName := v.String()
		if cfgName == "" {
			cfgName = "cpubench.toml"
		}
		htoml.DecodeFromFile(&ls, cfgName)

		if benchResult != nil {
			hit := false
			for i, v2 := range ls.Items {
				if v2.ModelName == benchResult.ModelName {
					hit = true
					ls.Items[i] = benchResult
					break
				}
			}
			if !hit {
				ls.Items = append(ls.Items, benchResult)
			}
		}

		sort.Slice(ls.Items, func(i, j int) bool {
			return ls.Items[i].Score > ls.Items[j].Score
			// return strings.Compare(ls.Items[i].ModelName, ls.Items[j].ModelName) < 0
		})

		htoml.EncodeToFile(&ls, cfgName, nil)

		table := simpletable.New()

		table.Header = &simpletable.Header{
			Cells: []*simpletable.Cell{
				{Text: "NAME"},
				{Text: "SCORE"},
				{Text: "MODEL"},
				{Text: "GHz"},
				{Text: "DATE"},
			},
		}

		for _, v2 := range ls.Items {
			table.Body.Cells = append(table.Body.Cells, []*simpletable.Cell{
				{Text: v2.Name},
				{Text: fmt.Sprintf("%d", v2.Score)},
				{Text: v2.ModelName},
				{Text: fmt.Sprintf("%.2f", v2.CpuInfo.Mhz/1000)},
				{Text: v2.TestTime},
			})
		}

		tstr := strings.TrimSpace(table.String())
		if n := strings.Index(tstr, "\n"); n > 0 {
			tstr = tstr[n+1:]
		}
		if n := strings.LastIndex(tstr, "\n"); n > 0 {
			tstr = tstr[:n]
		}

		tstr = strings.Replace(tstr, "+--", "|--", -1)
		tstr = strings.Replace(tstr, "--+", "--|", -1)

		fmt.Println(tstr)
	}
}

// refer: J.Marchin PI
//  PI = [16/5 - 16/(3*53) + 16/(5*55) - 16/(7*57) + ......]
//  PI -= [4/239 - 4/(3*2393) + 4/(5*2395) - 4/(7*2397) + ......]
func CpuBench(base [2]int64) (int64, int64) {

	t := time.Now()

	L := int(base[0])
	N := (L)/4 + 1

	s := make([]int, N+3)
	w := make([]int, N+3)
	v := make([]int, N+3)
	q := make([]int, N+3)
	n := (int)(float64(L)/1.39793 + 1)

	w[0] = 16 * 5
	v[0] = 4 * 239

	for k := 1; k <= n; k++ {

		div(w, 25, w, N)
		div(v, 239, v, N)
		div(v, 239, v, N)
		sub(w, v, q, N)
		div(q, ((2 * k) - 1), q, N)

		if (k % 2) == 1 {
			add(s, q, s, N)
		} else {
			sub(s, q, s, N)
		}
	}

	if false {
		fmt.Printf("%d.", s[0])
		for k := 1; k < N; k++ {
			fmt.Printf("%04d", s[k])
		}
	}

	tl := time.Now().UnixNano() - t.UnixNano()
	return tl, base[1] / tl
}

func add(a []int, b []int, c []int, N int) {
	carry := 0
	for i := N + 1; i >= 0; i-- {
		c[i] = a[i] + b[i] + carry
		if c[i] < 10000 {
			carry = 0
		} else {
			c[i] = c[i] - 10000
			carry = 1
		}
	}
}

func sub(a []int, b []int, c []int, N int) {
	borrow := 0
	for i := N + 1; i >= 0; i-- {
		c[i] = a[i] - b[i] - borrow
		if c[i] >= 0 {
			borrow = 0
		} else {
			c[i] = c[i] + 10000
			borrow = 1
		}
	}
}

func div(a []int, b int, c []int, N int) {
	tmp, remain := 0, 0
	for i := 0; i <= (N + 1); i++ {
		tmp = a[i] + remain
		c[i] = tmp / b
		remain = (tmp % b) * 10000
	}
}
