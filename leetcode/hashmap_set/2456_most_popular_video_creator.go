package leetcode

// https://leetcode.com/problems/most-popular-video-creator/

// sum of views of ALL creator's videos
// find highest popularity and ID of most viewed videos

// views[i] - video with ids[i] - belongs to creators[i]
// map[creators[i]] = highest view video ids
// map[creators[i]] = max views
func mostPopularCreator(creators []string, ids []string, views []int) [][]string {
	popularity := make(map[string]int)
	mostViewed := make(map[string]int)
	mostViewedId := make(map[string]string)

	maxPop := 0

	for i := range views {
		creator := creators[i]
		popularity[creator] += views[i]
		if maxPop < popularity[creator] {
			maxPop = popularity[creator]
		}
		if views[i] > mostViewed[creator] {
			mostViewed[creator] = views[i]
			mostViewedId[creator] = ids[i]
		} else if views[i] == mostViewed[creator] {
			if mostViewedId[creator] == "" || mostViewedId[creator] > ids[i] {
				mostViewedId[creator] = ids[i]
			}
		}
	}

	res := [][]string{}
	for name, pop := range popularity {
		if pop == maxPop {
			pair := []string{name, mostViewedId[name]}
			res = append(res, pair)
		}
	}

	return res
}
