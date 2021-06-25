package MessageMiddle

import "sync"

var themeId2SubId2NilMap = map[string]map[string]struct{}{}

var themeId2SubId2NilMapLock = &sync.Mutex{}

func LockThemeSubFn(fn func()) {
	LockFn(themeId2SubId2NilMapLock, fn)
}

func ThemeAddSub(themeId, subId string) error {
	LockThemeSubFn(func() {
		if themeId2SubId2NilMap[themeId] == nil {
			themeId2SubId2NilMap[themeId] = map[string]struct{}{}
		}
		themeId2SubId2NilMap[themeId][subId] = struct{}{}
	})
	return nil
}

func ThemeGetAllSubIdByThemeId(themeId string) []string {
	subIds := []string{}
	LockThemeSubFn(func() {
		if themeId2SubId2NilMap[themeId] == nil {
			return
		}
		for k := range themeId2SubId2NilMap[themeId] {
			subIds = append(subIds, k)
		}
	})
	return subIds
}
