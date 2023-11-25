package script

import (
	"fmt"
	"os"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

func pressureTest(targeter vegeta.Targeter, rate vegeta.ConstantPacer, duration time.Duration) error {
	attacker := vegeta.NewAttacker()

	fmt.Printf("start pressure test, duration:%s , rate: %d/s\n\n", duration, rate.Freq)
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()

	report := vegeta.NewTextReporter(&metrics)
	report.Report(os.Stdout)

	return nil
}
