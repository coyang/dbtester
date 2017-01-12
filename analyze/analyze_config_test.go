// Copyright 2017 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package analyze

import "testing"

func Test_readConfig(t *testing.T) {
	c, err := readConfig("analyze_config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	if c.Title != `Write 2M keys, 1000-client (etcd v3.1 100-conn), 8-byte key, 256-byte value` {
		t.Fatalf("unexpected Title %q", c.Title)
	}
	if c.AllAggregatedPath != "2017Q1-01-etcd-zookeeper-consul/01-write-2M-keys/aggregated.csv" {
		t.Fatalf("unexpected AllAggregatedPath %q", c.AllAggregatedPath)
	}

	if c.RawData[0].Legend != "etcd v3.1 (Go 1.7.4)" {
		t.Fatalf("unexpected c.RawData[0].Legend %q", c.RawData[0].Legend)
	}
	if c.RawData[0].DatabaseTag != "etcd-v3.1-go1.7.4" {
		t.Fatalf("unexpected c.RawData[0].DatabaseTag %q", c.RawData[0].DatabaseTag)
	}
	if c.RawData[0].OutputPath != "2017Q1-01-etcd-zookeeper-consul/01-write-2M-keys/etcd-v3.1-go1.7.4-aggregated.csv" {
		t.Fatalf("unexpected c.RawData[0].OutputPath %q", c.RawData[0].OutputPath)
	}
	if c.RawData[0].SourceSystemMetricsPaths[0] != "2017Q1-01-etcd-zookeeper-consul/01-write-2M-keys/etcd-v3.1-go1.7.4-1-system-metrics.csv" {
		t.Fatalf("unexpected c.RawData[0].SourceSystemMetricsPaths[0] %q", c.RawData[0].SourceSystemMetricsPaths[0])
	}
	if c.RawData[0].SourceBenchmarkMetricsPath != "2017Q1-01-etcd-zookeeper-consul/01-write-2M-keys/etcd-v3.1-go1.7.4-bench-result-timeseries.csv" {
		t.Fatalf("unexpected c.RawData[0].SourceBenchmarkMetricsPath %q", c.RawData[0].SourceBenchmarkMetricsPath)
	}
	if c.RawData[1].DatabaseTag != "zookeeper-r3.4.9-java8" {
		t.Fatalf("unexpected c.RawData[1].DatabaseTag %q", c.RawData[1].DatabaseTag)
	}

	if c.READMEConfig.OutputPath != "2017Q1-01-etcd-zookeeper-consul/01-write-2M-keys/README.md" {
		t.Fatalf("unexpected %s", c.READMEConfig.OutputPath)
	}
	if c.READMEConfig.Results[0].Images[0].ImageTitle != "2017Q1-01-etcd-zookeeper-consul/01-write-2M-keys/AVG-LATENCY-MS" {
		t.Fatalf("unexpected %s", c.READMEConfig.Results[0].Images[0].ImageTitle)
	}
	if c.READMEConfig.Results[0].Images[0].ImagePath != "https://storage.googleapis.com/dbtester-results/2017Q1-01-etcd-zookeeper-consul/01-write-2M-keys/AVG-LATENCY-MS.svg" {
		t.Fatalf("unexpected %s", c.READMEConfig.Results[0].Images[0].ImagePath)
	}
	if c.READMEConfig.Results[0].Images[0].ImageType != "remote" {
		t.Fatalf("unexpected %s", c.READMEConfig.Results[0].Images[0].ImageType)
	}
}