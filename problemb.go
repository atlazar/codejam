package main

import (
	"fmt"
	"errors"
	"sort"
)

type parent uint8

const (
	cameron parent = iota
	jamie   parent = iota
)

type interval struct {
	Begin, End int
	Person     parent
}

func SolveB() (string, error) {
	var cameronCount, jamieCount int
	if n, err := fmt.Scan(&cameronCount, &jamieCount); n != 2 || err != nil {
		return "", errors.New("unable to read activities number")
	}

	cameronActivities, err := readIntervals(cameronCount, cameron)
	if err != nil {
		return "", err
	}

	jamieActivities, err := readIntervals(jamieCount, jamie)
	if err != nil {
		return "", err
	}

	return countExchanges(cameronActivities, jamieActivities), nil
}

func readIntervals(count int, person parent) ([]interval, error) {
	cameronActivities := make([]interval, count)
	for i := 0; i < count; i++ {
		current := interval{Person: person}
		if n, err := fmt.Scan(&current.Begin, &current.End); n != 2 || err != nil {
			return nil, fmt.Errorf("unable to read %v %v interval", person, i+1)
		}
		cameronActivities[i] = current
	}
	return cameronActivities, nil
}

func countExchanges(cameronActivities, jamieActivities []interval) string {

	cameronCount := len(cameronActivities)
	jamieCount := len(jamieActivities)
	totalCount := cameronCount + jamieCount

	activities := make([]interval, totalCount)
	copy(activities[:], cameronActivities)
	copy(activities[cameronCount:], jamieActivities)
	sort.Slice(activities[:], func(i, j int) bool {
		return activities[i].Begin < activities[j].Begin
	});
	activities = append(activities, activities[0])
	activities[totalCount].Begin += 1440;
	activities[totalCount].End += 1440;

	//Count mandatory exchanges and time can be assign to any parent
	var exchanges, anyParentTime int
	for i := 0; i < totalCount; i++ {
		current := activities[i]
		next := activities[i+1]
		if current.Person != next.Person {
			exchanges++;
			anyParentTime += next.Begin - current.End
		}
	}

	cameronTime, cameronRemovable := countTime(cameron, totalCount, activities)
	jamieTime, jamieRemovable := countTime(jamie, totalCount, activities)

	if cameronOvercome := cameronTime - jamieTime; cameronOvercome > anyParentTime {
		exchanges += decrementTime(cameronRemovable, cameronOvercome-anyParentTime)
	} else if jamieOvercome := jamieTime - cameronTime; jamieOvercome > anyParentTime {
		exchanges += decrementTime(jamieRemovable, jamieOvercome-anyParentTime)
	}

	return fmt.Sprintf("%v", exchanges)
}
func decrementTime(removable []int, delta int) int {

	for i := 0; i < len(removable); i++ {
		delta -= 2 * removable[i]
		if delta <= 0 {
			return 2 * (i + 1)
		}
	}
	panic("Not enought removable intervals")
}
func countTime(forParent parent, totalCount int, activities []interval) (parentTime int, parentRemovable []int) {
	otherParent := jamie - forParent;
	for i := 0; i < totalCount; i++ {
		current := activities[i]
		if current.Person == otherParent {
			parentTime += current.End - current.Begin
			if next := activities[i+1]; next.Person == otherParent {
				gap := next.Begin - current.End
				parentTime += gap
				parentRemovable = append(parentRemovable, gap)
			}
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(parentRemovable)))
	return
}
