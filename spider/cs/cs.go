package cs

//guid生成器 必须保证该结构体在程序中只初始化一次，否则有可能得到重复id
//依照snowflake算法得来

import (
	"errors"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

const (
	//Poch ( 2017-05-27 16:52:35.250507739 +0800 CST ).UnixNano() / 1e6
	Poch = 1495875155250
	//WorkerIDBits WorkerId所占的位
	WorkerIDBits = uint64(10)
	//SenquenceBits 序列号占的位
	SenquenceBits = uint64(12)
	//WorkerIDShift 参照
	WorkerIDShift = SenquenceBits
	//TimeStampShift 参照
	TimeStampShift = SenquenceBits + WorkerIDBits
	//SequenceMask 最大序列号值 4095(4096个)
	SequenceMask = int64(-1) ^ (int64(-1) << SenquenceBits)
	//MaxWorker 最大客户端标志值 1023(1024个)
	MaxWorker = int64(-1) ^ (int64(-1) << WorkerIDBits)
)

//GUID GUID定义
type GUID struct {
	sync.Mutex
	//Sequence 序列号
	Sequence int64
	//lastTimestamp 上一次时间戳
	lastTimeStamp int64
	//lastID 上一次生成的id
	lastID int64
	//WorkID
	WorkID int64
}

//NewGUID 获取一个GUID对象
func NewGUID(workID int64) (*GUID, error) {
	var g *GUID
	if workID > MaxWorker {
		return nil, errors.New("工作进程id超出最大值:" + strconv.FormatInt(MaxWorker, 10))
	}
	g = new(GUID)
	return g, nil
}

//milliseconds 获得当前毫秒时间
func (g *GUID) milliseconds() int64 {
	return time.Now().UnixNano() / 1e6
}

//NextID 获取一个GUID
func (g *GUID) NextID() (int64, error) {
	var ts int64
	var err error
	g.Lock()
	defer g.Unlock()
	ts = g.milliseconds()
	if ts == g.lastTimeStamp {
		g.Sequence = (g.Sequence + 1) & SequenceMask
		if g.Sequence == 0 {
			ts = g.timeStamp(ts)
		}
	} else {
		g.Sequence = 0
	}

	if ts < g.lastTimeStamp {
		err = errors.New("时钟过期")
		return 0, err
	}
	g.lastTimeStamp = ts
	ts = (ts-Poch)<<TimeStampShift | g.WorkID<<WorkerIDShift | g.Sequence
	return ts, nil
}

//timeStamp 获取一个可用时间基数
func (g *GUID) timeStamp(lastTimeStamp int64) int64 {
	ts := g.milliseconds()
	for {
		if ts < lastTimeStamp {
			ts = g.milliseconds()
		} else {
			break
		}
	}
	return ts
}

//GetIncreaseID 并发环境下生成一个增长的id,按需设置局部变量或者全局变量
func (g *GUID) GetIncreaseID(ID *uint64) uint64 {
	var n, v uint64
	for {
		v = atomic.LoadUint64(ID)
		n = v + 1
		if atomic.CompareAndSwapUint64(ID, v, n) {
			break
		}
	}
	return n
}
