package service

import (
	"UserService/dao"
	"errors"
	"github.com/garyburd/redigo/redis"
	"log"
	"strconv"
	"strings"
)

type LikeServiceImpl struct {
}

var OneMonth = 60 * 60 * 24 * 30

// IsLike 当前用户 uid 是否点赞了当前视频 vid
func (like *LikeServiceImpl) IsLike(uid int, vid int) (bool, error) {
	// 将 uid 转换为 string 类型
	strUID := strconv.FormatInt(int64(uid), 10)
	// 将 vid 转换为 string 类型
	strVID := strconv.FormatInt(int64(vid), 10)
	/*
		step1: 查询 RdLikeUID 中是否存在 key:strUID - value:strVID
	*/
	isExistUID, err := redis.Bool(RdLikeUID.Server.Do("EXISTS", strUID))
	if err != nil {
		log.Println(err.Error())
		return false, errors.New("查询是否存在 strUID 失败")
	}
	// 存在 strUID
	var isExistVID bool
	if isExistUID == true {
		isExistVID, err = redis.Bool(RdLikeUID.Server.Do("SIsMember", strUID, strVID))
		if err != nil {
			log.Println(err.Error())
			return false, errors.New("查询 strUID 中是否存在 strVID 失败")
		}
		// strUID 中存在 strVID
		if isExistVID == true {
			return isExistVID, nil
		}
	}
	/*
		step2: RdLikeUID 不存在 strUID 或是 strUID 中不存在 strVID，查询 RdLikeVID
	*/
	if isExistUID == false || isExistVID == false {
		isExistVID, err = redis.Bool(RdLikeVID.Server.Do("EXISTS", strVID))
		if err != nil {
			log.Println(err.Error())
			return false, errors.New("查询是否存在 strVID 失败")
		}
		// 存在 strVID
		if isExistVID == true {
			isExistUID, err = redis.Bool(RdLikeVID.Server.Do("SIsMember", strVID, strUID))
			if err != nil {
				log.Println(err.Error())
				return false, errors.New("查询 strVID 中是否存在 strUID 失败")
			}
			// strVID 中存在 strUID
			if isExistUID == true {
				return isExistUID, nil
			}
		}
	}
	//// 给 strUID 设置有效期一个月
	//_, err = RdLikeUID.Server.Do("Expire", strUID, time.Duration(OneMonth)*time.Second)
	//if err != nil {
	//	RdLikeUID.Server.Do("Del", strUID)
	//	log.Println(err.Error())
	//	return false, errors.New("strUID 设置有效期失败")
	//}
	/*
		step3: RdLikeUID RdLikeVID 中都没有对应 key-value，通过 uid 查询所有点赞 vid，并维护到 RdLikeUID 中
	*/
	err = like.UpRdUID(uid, strUID)
	if err != nil {
		return false, errors.New("维护 RdLikeUID 失败")
	}
	// 查询 RdLikeUID strUID 中是否存在 strVID
	isExistVID, err = redis.Bool(RdLikeUID.Server.Do("SIsMember", strUID, strVID))
	if err != nil {
		log.Println(err.Error())
		return false, errors.New("查询 strUID 中是否存在 strVID 失败")
	}
	return isExistVID, nil
}

// VideoLikeCount 当前视频 vid 被点赞数
func (like *LikeServiceImpl) VideoLikeCount(vid int) (int, error) {
	strVID := strconv.FormatInt(int64(vid), 10)
	isExistVID, err := redis.Bool(RdLikeVID.Server.Do("EXISTS", strVID))
	if err != nil {
		log.Println(err.Error())
		return 0, errors.New("查询是否存在 strVID 失败")
	}
	if isExistVID == true {
		// 获取集合中 uid 个数
		count, err1 := redis.Int(RdLikeVID.Server.Do("SCard", strVID))
		if err1 != nil {
			log.Println(err.Error())
			return 0, errors.New("获取集合中 uid 个数失败")
		}
		return count, nil
	}
	//_, err = RdLikeVID.Server.Do("Expire", strVID, time.Duration(OneMonth)*time.Second)
	//if err != nil {
	//	RdLikeVID.Server.Do("Del", strVID)
	//	log.Println(err.Error())
	//	return 0, errors.New("strVID 设置有效期失败")
	//}
	// 通过 vid 查询所有点赞 uid，并维护到 RdLikeVID 中
	err = like.UpRdVID(vid, strVID)
	if err != nil {
		return 0, errors.New("维护 RdLikeVID 失败")
	}
	// 再通过 set 集合中 uid 个数，获取点赞数量
	count, err2 := redis.Int(RdLikeVID.Server.Do("SCard", strVID))
	if err2 != nil {
		log.Println(err.Error())
		return 0, errors.New("获取集合中 uid 个数失败")
	}
	return count, nil
}

// LikeListCount 当前用户 uid 点赞视频总数
func (like *LikeServiceImpl) LikeListCount(uid int) (int, error) {
	strUID := strconv.FormatInt(int64(uid), 10)
	isExistUID, err := redis.Bool(RdLikeUID.Server.Do("EXISTS", strUID))
	if err != nil {
		log.Println(err.Error())
		return 0, errors.New("查询是否存在 strUID 失败")
	}
	if isExistUID == true {
		// 获取集合中 vid 个数
		count, err1 := redis.Int(RdLikeUID.Server.Do("SCard", strUID))
		if err1 != nil {
			log.Println(err.Error())
			return 0, errors.New("获取集合中 vid 个数失败")
		}
		return count, nil
	}
	//_, err = RdLikeUID.Server.Do("Expire", strUID, time.Duration(OneMonth)*time.Second)
	//if err != nil {
	//	RdLikeUID.Server.Do("Del", strUID)
	//	log.Println(err.Error())
	//	return 0, errors.New("strUID 设置有效期失败")
	//}
	err = like.UpRdUID(uid, strUID)
	if err != nil {
		return 0, errors.New("维护 RdLikeUID 失败")
	}
	count, err2 := redis.Int(RdLikeUID.Server.Do("SCard", strUID))
	if err2 != nil {
		log.Println(err.Error())
		return 0, errors.New("获取集合中 vid 个数失败")
	}
	return count, nil
}

// LikeAction 当前用户 uid 对视频 vid 的点赞操作，1-点赞，2-取消点赞 (还没调好)
func (like *LikeServiceImpl) LikeAction(uid int, vid int, act int) error {
	strUID := strconv.FormatInt(int64(uid), 10)
	strVID := strconv.FormatInt(int64(vid), 10)
	// 将点赞或者取消点赞的信息打入消息队列 RmqLikeAdd 或者 RmqLikeDel
	mss := strings.Builder{}
	mss.WriteString(strUID)
	mss.WriteString(" ")
	mss.WriteString(strVID)

	// 执行点赞操作
	if act == 1 {
		isExistUID, err := redis.Bool(RdLikeUID.Server.Do("EXISTS", strUID))
		if err != nil {
			log.Println(err.Error())
			return errors.New("查询是否存在 strUID 失败")
		}
		if isExistUID == true {
			// strUID 中添加当前 vid
			_, err = RdLikeUID.Server.Do("SAdd", strUID, vid)
			if err != nil {
				log.Println(err.Error())
				return errors.New("添加 vid 失败")
			}
			// 加入消息队列中准备对数据库操作
			RmqLikeAdd.Publish(mss.String())
		} else {
			//// 给 strUID 设置有效期一个月
			//_, err = RdLikeUID.Server.Do("Expire", strUID, time.Duration(OneMonth)*time.Second)
			//if err != nil {
			//	RdLikeUID.Server.Do("Del", strUID)
			//	log.Println(err.Error())
			//	return false, errors.New("strUID 设置有效期失败")
			//}
			err = like.UpRdUID(uid, strUID)
			if err != nil {
				return errors.New("维护 RdLikeUID 失败")
			}
			_, err = RdLikeUID.Server.Do("SAdd", strUID, vid)
			if err != nil {
				log.Println(err.Error())
				return errors.New("添加点赞当前视频失败")
			} else {
				RmqLikeAdd.Publish(mss.String())
			}
		}
		// 维护 RdbLikeVID
		isExistVID, err1 := redis.Bool(RdLikeVID.Server.Do("EXISTS", strVID))
		if err1 != nil {
			log.Println(err1.Error())
			return errors.New("查询是否存在 strVID 失败")
		}
		if isExistVID == true {
			_, err = RdLikeVID.Server.Do("SAdd", strVID, uid)
			if err != nil {
				log.Println(err.Error())
				return errors.New("添加 uid 失败")
			}
		} else {
			//_, err = RdLikeVID.Server.Do("Expire", strVID, time.Duration(OneMonth)*time.Second)
			//if err != nil {
			//	RdLikeVID.Server.Do("Del", strVID)
			//	log.Println(err.Error())
			//	return 0, errors.New("strVID 设置有效期失败")
			//}
			err = like.UpRdVID(vid, strVID)
			if err != nil {
				return errors.New("维护 RdLikeVID 失败")
			}
			_, err = RdLikeVID.Server.Do("SAdd", strVID, uid)
			if err != nil {
				log.Println(err.Error())
				return errors.New("添加 uid 失败")
			}
		}
	}
	// 执行取消赞操作
	if act == 2 {
		isExistUID, err := redis.Bool(RdLikeUID.Server.Do("EXISTS", strUID))
		if err != nil {
			log.Println(err.Error())
			return errors.New("查询是否存在 strUID 失败")
		}
		if isExistUID == true {
			// strUID 中删除当前 vid
			_, err = RdLikeUID.Server.Do("SRem", strUID, vid)
			if err != nil {
				log.Println(err.Error())
				return errors.New("删除 vid 失败")
			}
			// 加入消息队列中准备对数据库操作
			RmqLikeDel.Publish(mss.String())
		} else {
			//// 给 strUID 设置有效期一个月
			//_, err = RdLikeUID.Server.Do("Expire", strUID, time.Duration(OneMonth)*time.Second)
			//if err != nil {
			//	RdLikeUID.Server.Do("Del", strUID)
			//	log.Println(err.Error())
			//	return false, errors.New("strUID 设置有效期失败")
			//}
			err = like.UpRdUID(uid, strUID)
			if err != nil {
				return errors.New("维护 RdLikeUID 失败")
			}
			_, err = RdLikeUID.Server.Do("SRem", strUID, strVID)
			if err != nil {
				log.Println(err.Error())
				return errors.New("取消点赞当前视频失败")
			} else {
				RmqLikeDel.Publish(mss.String())
			}
		}
		// 维护 RdbLikeVID
		isExistVID, err1 := redis.Bool(RdLikeVID.Server.Do("EXISTS", strVID))
		if err1 != nil {
			log.Println(err1.Error())
			return errors.New("查询是否存在 strVID 失败")
		}
		if isExistVID == true {
			_, err = RdLikeVID.Server.Do("SRem", strVID, uid)
			if err != nil {
				log.Println(err.Error())
				return errors.New("删除 uid 失败")
			}
		} else {
			//_, err = RdLikeVID.Server.Do("Expire", strVID, time.Duration(OneMonth)*time.Second)
			//if err != nil {
			//	RdLikeVID.Server.Do("Del", strVID)
			//	log.Println(err.Error())
			//	return 0, errors.New("strVID 设置有效期失败")
			//}
			err = like.UpRdVID(vid, strVID)
			if err != nil {
				return errors.New("维护 RdLikeVID 失败")
			}
			_, err = RdLikeVID.Server.Do("SRem", strVID, uid)
			if err != nil {
				log.Println(err.Error())
				return errors.New("删除 uid 失败")
			}
		}
	}
	return nil
}

// UpRdUID 维护 RdLikeUID，遍历 vidList 加入
func (like *LikeServiceImpl) UpRdUID(uid int, strUID string) error {
	vidList, err := dao.LikeList(uid)
	if err != nil {
		return errors.New("维护 RdLikeUID 失败")
	}
	for _, likeVID := range vidList {
		_, err = RdLikeUID.Server.Do("SAdd", strUID, likeVID)
		if err != nil {
			RdLikeUID.Server.Do("Del", strUID)
			log.Println(err.Error())
			return errors.New("添加喜欢视频失败")
		}
	}
	return nil
}

// UpRdVID 维护 RdLikeVID，遍历 uidList 加入
func (like *LikeServiceImpl) UpRdVID(vid int, strVID string) error {
	uidList, err := dao.LikeUserList(vid)
	if err != nil {
		return errors.New("维护 RdLikeVID 失败")
	}
	for _, likeUID := range uidList {
		_, err = RdLikeVID.Server.Do("SAdd", strVID, likeUID)
		if err != nil {
			RdLikeUID.Server.Do("Del", strVID)
			log.Println(err.Error())
			return errors.New("添加喜欢用户失败")
		}
	}
	return nil
}
