package compose

import (
	"fmt"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app/components/calendar"
	"github.com/kudrykv/latex-yearly-planner/app/components/page"
	"github.com/kudrykv/latex-yearly-planner/app/config"
)

func Daily(cfg config.Config, tpls []string) (page.Modules, error) {
	if len(tpls) != 1 {
		return nil, fmt.Errorf("exppected one tpl, got %d %v", len(tpls), tpls)
	}

	modules := make(page.Modules, 0, 3)
	day := calendar.DayTime{Time: time.Date(cfg.Year, time.January, 1, 0, 0, 0, 0, time.Local)}

	for day.Year() == cfg.Year {
		modules = append(modules, page.Module{
			Cfg: cfg,
			Tpl: tpls[0],
			Body: map[string]interface{}{
				"Day":   day,
				"Hours": Hours(cfg.Layout.Numbers.DailyBottomHour, cfg.Layout.Numbers.DailyTopHour),
			},
		})

		day = day.AddDate(0, 0, 1)
	}

	return modules, nil
}

func Hours(from, to int) []calendar.DayTime {
	moment := time.Date(1, 1, 1, from, 0, 0, 0, time.Local)

	out := make([]calendar.DayTime, 0, to-from+1)

	for i := from; i <= to; i++ {
		out = append(out, calendar.DayTime{Time: moment})
		moment = moment.Add(time.Hour)
	}

	return out
}
