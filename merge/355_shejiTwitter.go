package main

type Twitte struct {
	time int
	tid  int
	next *Twitte
}

type User struct {
	userId   int
	tList    *Twitte
	followee []int
}

type Twitter struct {
	time  int
	users map[int]*User
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	t := Twitter{
		time:  0,
		users: make(map[int]*User),
	}
	return t
}

func ConstructTwitee(user *User, time int, tid int) {
	newt := &Twitte{
		time: time,
		tid:  tid,
		next: nil,
	}
	newt.next = user.tList
	user.tList = newt
}

func (this *Twitter) ConstructUser(uid int) {
	user := &User{
		userId:   uid,
		tList:    nil,
		followee: []int{uid},
	}

	this.users[uid] = user
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	if _, ok := this.users[userId]; !ok {
		this.ConstructUser(userId)
	}

	u := this.users[userId]

	ConstructTwitee(u, this.time, tweetId)
	this.time++
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {

	if _, ok := this.users[userId]; !ok {
		this.ConstructUser(userId)
	}
	followee := this.users[userId].followee
	maxCount := 10
	tCount := 0

	maxHeap := make([]*Twitte, 0)
	rst := []int{}
	for i := 0; i < len(followee); i++ {
		if this.users[followee[i]].tList == nil {
			continue
		}
		maxHeap = append(maxHeap, this.users[followee[i]].tList)
	}

	BuildHeap(maxHeap)

	for tCount < maxCount && len(maxHeap) > 0 {
		max := maxHeap[0]
		rst = append(rst, max.tid)
		tCount++

		max = max.next
		if max != nil {
			maxHeap[0] = max
			AdjustHeap(maxHeap, 0)
		} else {
			maxHeap[0] = maxHeap[len(maxHeap)-1]
			maxHeap = maxHeap[:len(maxHeap)-1]
			AdjustHeap(maxHeap, 0)
		}
	}

	return rst

}

func BuildHeap(heap []*Twitte) {

	length := len(heap)
	for i := length/2 - 1; i >= 0; i-- {
		AdjustHeap(heap, i)
	}
}

func AdjustHeap(heap []*Twitte, i int) {

	max := i
	left := 2*i + 1
	right := 2*i + 2
	if left < len(heap) && heap[max].time < heap[left].time {
		max = left
	}

	if right < len(heap) && heap[max].time < heap[right].time {
		max = right
	}

	if max != i {
		heap[i], heap[max] = heap[max], heap[i]
		AdjustHeap(heap, max)
	}
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := this.users[followeeId]; !ok {
		this.ConstructUser(followeeId)
	}

	if _, ok := this.users[followerId]; !ok {
		this.ConstructUser(followerId)
	}

	for _, v := range this.users[followerId].followee {
		if v == followeeId {
			return
		}
	}

	this.users[followerId].followee = append(this.users[followerId].followee, followeeId)
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {

	if followeeId == followerId {
		return
	}

	if _, ok := this.users[followeeId]; !ok {
		this.ConstructUser(followeeId)
	}

	if _, ok := this.users[followerId]; !ok {
		this.ConstructUser(followerId)
	}

	for i, v := range this.users[followerId].followee {
		if v == followeeId {
			this.users[followerId].followee = append(this.users[followerId].followee[:i], this.users[followerId].followee[i+1:]...)
			break
		}
	}
}
