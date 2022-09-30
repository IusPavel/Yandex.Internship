// B. Через тернии к клиенту
// Если по времени не уложимся, то попробовать переделать под чтение с файла (parseAndSortLogs)
// Также попробовать использовать в parseAndSortLogs указатель ([]*log)
package main

import (
	"fmt"
	"os"
	"sort"
)

type Rocket struct {
	ID				int
	MinutesInAction	int
	TimeAccepted	MyTime
}

type log struct {
	id			int
	day			int
	hour		int
	minute		int
	status		string
}

type MyTime struct {
	days	int
	hours	int
	minutes	int
}

func	NewMyTime(days, hours, minutes int) MyTime {
	return MyTime{
		days: days,
		hours: hours,
		minutes: minutes,
	}
}

func (t MyTime) SubMinutes(u MyTime) int {
	return (t.days - u.days) * 24 * 60 + (t.hours - u.hours) * 60 + (t.minutes - u.minutes)
}

func parseAndSortLogs(rows int) []log {
	logs := make([]log, 0, rows)
	for rows > 0 {
		var logTmp log
		fmt.Fscan(os.Stdin, &logTmp.day, &logTmp.hour, &logTmp.minute, &logTmp.id, &logTmp.status)
		if logTmp.status != "B" {
			logs = append(logs, logTmp)
		}
		rows--
	}
	sort.SliceStable(logs, func(i, j int) bool {
		if logs[i].day == logs[j].day && logs[i].hour == logs[j].hour {
			return logs[i].minute < logs[j].minute
		} else if logs[i].day == logs[j].day {
			return logs[i].hour < logs[j].hour
		} else {
			return logs[i].day < logs[j].day
		}
	})
	return logs
}

func splitLogsIntoRockets(logs []log) map[int]*Rocket {
	rockets := make(map[int]*Rocket)
	for _, v := range logs {
		if v.status == "A" {
			if _, ok := rockets[v.id]; !ok {
				rockets[v.id] = &Rocket{
					ID: v.id,
					MinutesInAction: 0,
					TimeAccepted: NewMyTime(v.day, v.hour, v.minute),
				}
			} else {
				r := rockets[v.id]
				r.TimeAccepted = NewMyTime(v.day, v.hour, v.minute)
			}
		} else if v.status == "S" || v.status == "C" {
			r := rockets[v.id]
			r.MinutesInAction += NewMyTime(v.day, v.hour, v.minute).SubMinutes(r.TimeAccepted)
		} else {
			panic("Impossible event. It's no way to reach this case")
		}
	}
	return rockets
}

func sortAndPrintMinutesByID(rockets map[int]*Rocket) {
	keys := make([]int, 0, len(rockets))
	for k := range rockets {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(a, b int) bool {
		return keys[a] < keys[b]
	})
	for i, v := range keys {
		fmt.Print(rockets[v].MinutesInAction)
		if i < len(keys) - 1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func main() {
	var rows int
	fmt.Fscanf(os.Stdin, "%d", &rows)
	logs := parseAndSortLogs(rows)
	rockets := splitLogsIntoRockets(logs)
	sortAndPrintMinutesByID(rockets)
}