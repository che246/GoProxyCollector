package scheduler

import (
	"github.com/AceDarkkinght/GoProxyCollector/collector"
	"github.com/AceDarkkinght/GoProxyCollector/storage"
	"github.com/AceDarkkinght/GoProxyCollector/util"
)

func Start(collector collector.Collector, storage storage.Storage) {
	if collector == nil {
		return
	}

	for {
		url := collector.Next()
		if url == "" {
			break
		}

		// Collect.
		results, err := collector.Collect(url)
		if err != nil {
			return
		}

		if len(results) == 0 {
			return
		}

		// Verify.
		for _, r := range results {
			if util.VerifyHTTP(r.Ip, r.Port) {
				storage.Update(r.Ip, r)
			}
		}
	}
}