// Code generated by $GOPATH/src/go-common/app/tool/cache/gen. DO NOT EDIT.

/*
  Package like is a generated cache proxy package.
  It is generated from:
  type _cache interface {
		// cache: -sync=true
		Like(c context.Context, id int64) (*likemdl.Item, error)
		// cache: -sync=true
		Likes(c context.Context, ids []int64) (map[int64]*likemdl.Item, error)
		// cache: -sync=true
		ActSubject(c context.Context, id int64) (*likemdl.SubjectItem, error)
		//cache: -sync=true -nullcache=-1 -check_null_code=$==-1
		LikeMissionBuff(ctx context.Context, sid int64, mid int64) (res int64, err error)
		//cache: -sync=true
		MissionGroupItems(ctx context.Context, lids []int64) (map[int64]*likemdl.MissionGroup, error)
		//cache: -sync=true -nullcache=-1 -check_null_code=$!=nil&&$==-1
		ActMission(ctx context.Context, sid int64, lid int64, mid int64) (res int64, err error)
		//cache:-sync=true
		ActLikeAchieves(ctx context.Context, sid int64) (res *likemdl.Achievements, err error)
		//cache:-sync=true
		ActMissionFriends(ctx context.Context, sid int64, lid int64) (res *likemdl.ActMissionGroups, err error)
		//cache:-sync=true
		ActUserAchieve(ctx context.Context, id int64) (res *likemdl.ActLikeUserAchievement, err error)
		// cache
		MatchSubjects(c context.Context, ids []int64) (map[int64]*likemdl.Object, error)
		// cache:-sync=true
		LikeContent(c context.Context, ids []int64) (map[int64]*likemdl.LikeContent, error)
		// cache
		SourceItemData(c context.Context, sid int64) ([]int64, error)
		// cache:-sync=true
		ActSubjectProtocol(c context.Context, sid int64) (res *likemdl.ActSubjectProtocol, err error)
	}
*/

package like

import (
	"context"

	likemdl "go-common/app/interface/main/activity/model/like"
	"go-common/library/stat/prom"
)

var _ _cache

// Like get data from cache if miss will call source method, then add to cache.
func (d *Dao) Like(c context.Context, id int64) (res *likemdl.Item, err error) {
	addCache := true
	res, err = d.CacheLike(c, id)
	if err != nil {
		addCache = false
		err = nil
	}
	if res != nil {
		prom.CacheHit.Incr("Like")
		return
	}
	prom.CacheMiss.Incr("Like")
	res, err = d.RawLike(c, id)
	if err != nil {
		return
	}
	miss := res
	if !addCache {
		return
	}
	d.AddCacheLike(c, id, miss)
	return
}

// Likes get data from cache if miss will call source method, then add to cache.
func (d *Dao) Likes(c context.Context, keys []int64) (res map[int64]*likemdl.Item, err error) {
	if len(keys) == 0 {
		return
	}
	addCache := true
	if res, err = d.CacheLikes(c, keys); err != nil {
		addCache = false
		res = nil
		err = nil
	}
	var miss []int64
	for _, key := range keys {
		if (res == nil) || (res[key] == nil) {
			miss = append(miss, key)
		}
	}
	prom.CacheHit.Add("Likes", int64(len(keys)-len(miss)))
	missLen := len(miss)
	if missLen == 0 {
		return
	}
	var missData map[int64]*likemdl.Item
	prom.CacheMiss.Add("Likes", int64(len(miss)))
	missData, err = d.RawLikes(c, miss)
	if res == nil {
		res = make(map[int64]*likemdl.Item, len(keys))
	}
	for k, v := range missData {
		res[k] = v
	}
	if err != nil {
		return
	}
	if !addCache {
		return
	}
	d.AddCacheLikes(c, missData)
	return
}

// ActSubject get data from cache if miss will call source method, then add to cache.
func (d *Dao) ActSubject(c context.Context, id int64) (res *likemdl.SubjectItem, err error) {
	addCache := true
	res, err = d.CacheActSubject(c, id)
	if err != nil {
		addCache = false
		err = nil
	}
	if res != nil {
		prom.CacheHit.Incr("ActSubject")
		return
	}
	prom.CacheMiss.Incr("ActSubject")
	res, err = d.RawActSubject(c, id)
	if err != nil {
		return
	}
	miss := res
	if !addCache {
		return
	}
	d.AddCacheActSubject(c, id, miss)
	return
}

// LikeMissionBuff get data from cache if miss will call source method, then add to cache.
func (d *Dao) LikeMissionBuff(c context.Context, id int64, mid int64) (res int64, err error) {
	addCache := true
	res, err = d.CacheLikeMissionBuff(c, id, mid)
	if err != nil {
		addCache = false
		err = nil
	}
	defer func() {
		if res == -1 {
			res = 0
		}
	}()
	if res != 0 {
		prom.CacheHit.Incr("LikeMissionBuff")
		return
	}
	prom.CacheMiss.Incr("LikeMissionBuff")
	res, err = d.RawLikeMissionBuff(c, id, mid)
	if err != nil {
		return
	}
	miss := res
	if miss == 0 {
		miss = -1
	}
	if !addCache {
		return
	}
	d.AddCacheLikeMissionBuff(c, id, miss, mid)
	return
}

// MissionGroupItems get data from cache if miss will call source method, then add to cache.
func (d *Dao) MissionGroupItems(c context.Context, keys []int64) (res map[int64]*likemdl.MissionGroup, err error) {
	if len(keys) == 0 {
		return
	}
	addCache := true
	if res, err = d.CacheMissionGroupItems(c, keys); err != nil {
		addCache = false
		res = nil
		err = nil
	}
	var miss []int64
	for _, key := range keys {
		if (res == nil) || (res[key] == nil) {
			miss = append(miss, key)
		}
	}
	prom.CacheHit.Add("MissionGroupItems", int64(len(keys)-len(miss)))
	missLen := len(miss)
	if missLen == 0 {
		return
	}
	var missData map[int64]*likemdl.MissionGroup
	prom.CacheMiss.Add("MissionGroupItems", int64(len(miss)))
	missData, err = d.RawMissionGroupItems(c, miss)
	if res == nil {
		res = make(map[int64]*likemdl.MissionGroup, len(keys))
	}
	for k, v := range missData {
		res[k] = v
	}
	if err != nil {
		return
	}
	if !addCache {
		return
	}
	d.AddCacheMissionGroupItems(c, missData)
	return
}

// ActMission get data from cache if miss will call source method, then add to cache.
func (d *Dao) ActMission(c context.Context, id int64, lid int64, mid int64) (res int64, err error) {
	addCache := true
	res, err = d.CacheActMission(c, id, lid, mid)
	if err != nil {
		addCache = false
		err = nil
	}
	defer func() {
		if res == -1 {
			res = 0
		}
	}()
	if res != 0 {
		prom.CacheHit.Incr("ActMission")
		return
	}
	prom.CacheMiss.Incr("ActMission")
	res, err = d.RawActMission(c, id, lid, mid)
	if err != nil {
		return
	}
	miss := res
	if miss == 0 {
		miss = -1
	}
	if !addCache {
		return
	}
	d.AddCacheActMission(c, id, miss, lid, mid)
	return
}

// ActLikeAchieves get data from cache if miss will call source method, then add to cache.
func (d *Dao) ActLikeAchieves(c context.Context, id int64) (res *likemdl.Achievements, err error) {
	addCache := true
	res, err = d.CacheActLikeAchieves(c, id)
	if err != nil {
		addCache = false
		err = nil
	}
	if res != nil {
		prom.CacheHit.Incr("ActLikeAchieves")
		return
	}
	prom.CacheMiss.Incr("ActLikeAchieves")
	res, err = d.RawActLikeAchieves(c, id)
	if err != nil {
		return
	}
	miss := res
	if !addCache {
		return
	}
	d.AddCacheActLikeAchieves(c, id, miss)
	return
}

// ActMissionFriends get data from cache if miss will call source method, then add to cache.
func (d *Dao) ActMissionFriends(c context.Context, id int64, lid int64) (res *likemdl.ActMissionGroups, err error) {
	addCache := true
	res, err = d.CacheActMissionFriends(c, id, lid)
	if err != nil {
		addCache = false
		err = nil
	}
	if res != nil {
		prom.CacheHit.Incr("ActMissionFriends")
		return
	}
	prom.CacheMiss.Incr("ActMissionFriends")
	res, err = d.RawActMissionFriends(c, id, lid)
	if err != nil {
		return
	}
	miss := res
	if !addCache {
		return
	}
	d.AddCacheActMissionFriends(c, id, miss, lid)
	return
}

// ActUserAchieve get data from cache if miss will call source method, then add to cache.
func (d *Dao) ActUserAchieve(c context.Context, id int64) (res *likemdl.ActLikeUserAchievement, err error) {
	addCache := true
	res, err = d.CacheActUserAchieve(c, id)
	if err != nil {
		addCache = false
		err = nil
	}
	if res != nil {
		prom.CacheHit.Incr("ActUserAchieve")
		return
	}
	prom.CacheMiss.Incr("ActUserAchieve")
	res, err = d.RawActUserAchieve(c, id)
	if err != nil {
		return
	}
	miss := res
	if !addCache {
		return
	}
	d.AddCacheActUserAchieve(c, id, miss)
	return
}

// MatchSubjects get data from cache if miss will call source method, then add to cache.
func (d *Dao) MatchSubjects(c context.Context, keys []int64) (res map[int64]*likemdl.Object, err error) {
	if len(keys) == 0 {
		return
	}
	addCache := true
	if res, err = d.CacheMatchSubjects(c, keys); err != nil {
		addCache = false
		res = nil
		err = nil
	}
	var miss []int64
	for _, key := range keys {
		if (res == nil) || (res[key] == nil) {
			miss = append(miss, key)
		}
	}
	prom.CacheHit.Add("MatchSubjects", int64(len(keys)-len(miss)))
	missLen := len(miss)
	if missLen == 0 {
		return
	}
	var missData map[int64]*likemdl.Object
	prom.CacheMiss.Add("MatchSubjects", int64(len(miss)))
	missData, err = d.RawMatchSubjects(c, miss)
	if res == nil {
		res = make(map[int64]*likemdl.Object, len(keys))
	}
	for k, v := range missData {
		res[k] = v
	}
	if err != nil {
		return
	}
	if !addCache {
		return
	}
	d.cache.Do(c, func(c context.Context) {
		d.AddCacheMatchSubjects(c, missData)
	})
	return
}

// LikeContent get data from cache if miss will call source method, then add to cache.
func (d *Dao) LikeContent(c context.Context, keys []int64) (res map[int64]*likemdl.LikeContent, err error) {
	if len(keys) == 0 {
		return
	}
	addCache := true
	if res, err = d.CacheLikeContent(c, keys); err != nil {
		addCache = false
		res = nil
		err = nil
	}
	var miss []int64
	for _, key := range keys {
		if (res == nil) || (res[key] == nil) {
			miss = append(miss, key)
		}
	}
	prom.CacheHit.Add("LikeContent", int64(len(keys)-len(miss)))
	missLen := len(miss)
	if missLen == 0 {
		return
	}
	var missData map[int64]*likemdl.LikeContent
	prom.CacheMiss.Add("LikeContent", int64(len(miss)))
	missData, err = d.RawLikeContent(c, miss)
	if res == nil {
		res = make(map[int64]*likemdl.LikeContent, len(keys))
	}
	for k, v := range missData {
		res[k] = v
	}
	if err != nil {
		return
	}
	if !addCache {
		return
	}
	d.AddCacheLikeContent(c, missData)
	return
}

// SourceItemData get data from cache if miss will call source method, then add to cache.
func (d *Dao) SourceItemData(c context.Context, id int64) (res []int64, err error) {
	addCache := true
	res, err = d.CacheSourceItemData(c, id)
	if err != nil {
		addCache = false
		err = nil
	}
	if len(res) != 0 {
		prom.CacheHit.Incr("SourceItemData")
		return
	}
	prom.CacheMiss.Incr("SourceItemData")
	res, err = d.RawSourceItemData(c, id)
	if err != nil {
		return
	}
	miss := res
	if !addCache {
		return
	}
	d.cache.Do(c, func(c context.Context) {
		d.AddCacheSourceItemData(c, id, miss)
	})
	return
}

// ActSubjectProtocol get data from cache if miss will call source method, then add to cache.
func (d *Dao) ActSubjectProtocol(c context.Context, id int64) (res *likemdl.ActSubjectProtocol, err error) {
	addCache := true
	res, err = d.CacheActSubjectProtocol(c, id)
	if err != nil {
		addCache = false
		err = nil
	}
	if res != nil {
		prom.CacheHit.Incr("ActSubjectProtocol")
		return
	}
	prom.CacheMiss.Incr("ActSubjectProtocol")
	res, err = d.RawActSubjectProtocol(c, id)
	if err != nil {
		return
	}
	miss := res
	if !addCache {
		return
	}
	d.AddCacheActSubjectProtocol(c, id, miss)
	return
}