package cherryFacade

import (
	"strings"
	"sync"

	cconst "github.com/cherry-game/cherry/const"
	cerr "github.com/cherry-game/cherry/error"
	cstring "github.com/cherry-game/cherry/extend/string"
	ctime "github.com/cherry-game/cherry/extend/time"
	cproto "github.com/cherry-game/cherry/net/proto"
	"github.com/nats-io/nats.go"
)

type (
	Message struct {
		BuildTime  int64            // message build time(ms)
		PostTime   int64            // post to actor time(ms)
		Source     string           // 来源actor path
		Target     string           // 目标actor path
		targetPath *ActorPath       // 目标actor path对象
		FuncName   string           // 请求调用的函数名
		Session    *cproto.Session  // session of gateway
		Args       interface{}      // 请求的参数
		Header     nats.Header      // nats.Msg Header
		Reply      string           // nats.Msg reply subject
		IsCluster  bool             // 是否为集群消息
		ChanResult chan interface{} //
	}

	// ActorPath = NodeID . ActorID
	// ActorPath = NodeID . ActorID . ChildID
	ActorPath struct {
		NodeID  string
		ActorID string
		ChildID string
	}
)

//var (
//	messagePool = &sync.Pool{
//		New: func() interface{} {
//			return new(Message)
//		},
//	}
//)

func GetMessage() Message {
	msg := Message{
		BuildTime: ctime.Now().ToMillisecond(),
	}

	return msg
}

func BuildClusterMessage(packet *cproto.ClusterPacket) Message {
	message := Message{
		BuildTime: packet.BuildTime,
		Source:    packet.SourcePath,
		Target:    packet.TargetPath,
		FuncName:  packet.FuncName,
		IsCluster: true,
		Session:   packet.Session,
		Args:      packet.ArgBytes,
	}

	return message
}

//func (p *Message) Recycle() {
//	p.BuildTime = 0
//	p.PostTime = 0
//	p.Source = ""
//	p.Target = ""
//	p.targetPath = nil
//	p.FuncName = "_"
//	p.Session = nil
//	p.Args = nil
//	p.Err = nil
//	p.ClusterReply = nil
//	p.ChanResult = nil
//	p.IsCluster = false
//	messagePool.Put(p)
//}

func (p *Message) TargetPath() *ActorPath {
	if p.targetPath == nil {
		p.targetPath, _ = ToActorPath(p.Target)
	}
	return p.targetPath
}

func (p *Message) IsReply() bool {
	return p.Reply != ""
}

func (p *Message) Destory() {
	p.targetPath = nil
	p.Session = nil
	p.Args = nil
	p.Header = nil
	p.ChanResult = nil
}

func (p *ActorPath) IsChild() bool {
	return p.ChildID != ""
}

func (p *ActorPath) IsParent() bool {
	return p.ChildID == ""
}

// String
func (p *ActorPath) String() string {
	return NewChildPath(p.NodeID, p.ActorID, p.ChildID)
}

func NewActorPath(nodeID, actorID, childID string) *ActorPath {
	return &ActorPath{
		NodeID:  nodeID,
		ActorID: actorID,
		ChildID: childID,
	}
}

func NewChildPath(nodeID, actorID, childID interface{}) string {
	if childID == "" {
		return NewPath(nodeID, actorID)
	}
	return cstring.ToString(nodeID) + cconst.DOT + cstring.ToString(actorID) + cconst.DOT + cstring.ToString(childID)
}

func NewPath(nodeID, actorID interface{}) string {
	return cstring.ToString(nodeID) + cconst.DOT + cstring.ToString(actorID)
}

type actorPathResult struct {
	actorPath *ActorPath
	err       error
}

var (
	actorPathCache = &sync.Map{} // key: path string, value: *actorPathResult
)

func ToActorPath(path string) (*ActorPath, error) {
	if path == "" {
		return nil, cerr.ActorPathError
	}

	if cached, ok := actorPathCache.Load(path); ok {
		result := cached.(*actorPathResult)
		return result.actorPath, result.err
	}

	// 使用 IndexByte 手动分割，避免 strings.Split 创建切片
	// 查找第一个分隔符的位置
	firstDot := strings.IndexByte(path, '.')
	if firstDot == -1 {
		// 没有找到分隔符，无效路径
		return nil, cerr.ActorPathError
	}

	// 提取 NodeID (第一个点之前的部分)
	nodeID := path[:firstDot]

	// 查找第二个分隔符的位置（从第一个点之后开始）
	secondDot := strings.IndexByte(path[firstDot+1:], '.')
	if secondDot == -1 {
		// 只有两个部分: NodeID.ActorID
		actorID := path[firstDot+1:]
		actorPathCache.Store(path, &actorPathResult{actorPath: NewActorPath(nodeID, actorID, ""), err: nil})
		return NewActorPath(nodeID, actorID, ""), nil
	}

	// 有三个部分: NodeID.ActorID.ChildID
	// secondDot 是相对于 firstDot+1 的位置，需要加上偏移
	actorID := path[firstDot+1 : firstDot+1+secondDot]
	childID := path[firstDot+1+secondDot+1:]

	// 检查是否有第四个部分（无效）
	if strings.IndexByte(childID, '.') != -1 {
		return nil, cerr.ActorPathError
	}

	actorPathCache.Store(path, &actorPathResult{actorPath: NewActorPath(nodeID, actorID, childID), err: nil})
	return NewActorPath(nodeID, actorID, childID), nil
}
