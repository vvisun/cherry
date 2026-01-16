package cherryNatsCluster

const (
	localSubjectFormat      = "cherry-%s.local.%s.%s"   // cherry.{prefix}.local.{nodeType}.{nodeID}
	remoteSubjectFormat     = "cherry-%s.remote.%s.%s"  // cherry.{prefix}.remote.{nodeType}.{nodeID}
	remoteTypeSubjectFormat = "cherry-%s.remoteType.%s" // cherry.{prefix}.remoteType.{nodeType}
	replySubjectFormat      = "cherry-%s.reply.%s.%s"   // cherry.{prefix}.reply.{nodeType}.{nodeID}
)

const long_threshold = 32

const (
	cherryPrefix     = "cherry-"
	localSuffix      = ".local."
	remoteSuffix     = ".remote."
	remoteTypeSuffix = ".remoteType."
	replySuffix      = ".reply."
)

// GetLocalSubject local message nats chan
func GetLocalSubject(prefix, nodeType, nodeID string) string {
	// return fmt.Sprintf(localSubjectFormat, prefix, nodeType, nodeID)
	//长字符串优化
	// if len(prefix) > long_threshold && len(nodeType) > long_threshold && len(nodeID) > long_threshold {
	// 	//使用strings.Builder
	// 	var builder strings.Builder
	// 	builder.Grow(len(cherryPrefix) + len(prefix) + len(localSuffix) + len(nodeType) + 1 + len(nodeID))
	// 	builder.WriteString(cherryPrefix)
	// 	builder.WriteString(prefix)
	// 	builder.WriteString(localSuffix)
	// 	builder.WriteString(nodeType)
	// 	builder.WriteByte('.')
	// 	builder.WriteString(nodeID)
	// 	return builder.String()
	// }
	return cherryPrefix + prefix + localSuffix + nodeType + "." + nodeID
}

// GetRemoteSubject remote message nats chan
func GetRemoteSubject(prefix, nodeType, nodeID string) string {
	// return fmt.Sprintf(remoteSubjectFormat, prefix, nodeType, nodeID)
	//长字符串优化
	// if len(prefix) > long_threshold && len(nodeType) > long_threshold && len(nodeID) > long_threshold {
	// 	//使用strings.Builder
	// 	var builder strings.Builder
	// 	builder.Grow(len(cherryPrefix) + len(prefix) + len(remoteSuffix) + len(nodeType) + 1 + len(nodeID))
	// 	builder.WriteString(cherryPrefix)
	// 	builder.WriteString(prefix)
	// 	builder.WriteString(remoteSuffix)
	// 	builder.WriteString(nodeType)
	// 	builder.WriteByte('.')
	// 	builder.WriteString(nodeID)
	// 	return builder.String()
	// }
	return cherryPrefix + prefix + remoteSuffix + nodeType + "." + nodeID
}

// GetRemoteTypeSubject remote type message nats chan
func GetRemoteTypeSubject(prefix, nodeType string) string {
	// return fmt.Sprintf(remoteTypeSubjectFormat, prefix, nodeType)
	//长字符串优化
	// if len(prefix) > long_threshold && len(nodeType) > long_threshold {
	// 	//使用strings.Builder
	// 	var builder strings.Builder
	// 	builder.Grow(len(cherryPrefix) + len(prefix) + len(remoteTypeSuffix) + len(nodeType))
	// 	builder.WriteString(cherryPrefix)
	// 	builder.WriteString(prefix)
	// 	builder.WriteString(remoteTypeSuffix)
	// 	builder.WriteString(nodeType)
	// 	return builder.String()
	// }
	return cherryPrefix + prefix + remoteTypeSuffix + nodeType
}

// GetReplySubject reply message nats chan
func GetReplySubject(prefix, nodeType, nodeID string) string {
	// return fmt.Sprintf(replySubjectFormat, prefix, nodeType, nodeID)
	//长字符串优化
	// if len(prefix) > long_threshold && len(nodeType) > long_threshold && len(nodeID) > long_threshold {
	// 	//使用strings.Builder
	// 	var builder strings.Builder
	// 	builder.Grow(len(cherryPrefix) + len(prefix) + len(replySuffix) + len(nodeType) + 1 + len(nodeID))
	// 	builder.WriteString(cherryPrefix)
	// 	builder.WriteString(prefix)
	// 	builder.WriteString(replySuffix)
	// 	builder.WriteString(nodeType)
	// 	builder.WriteByte('.')
	// 	builder.WriteString(nodeID)
	// 	return builder.String()
	// }
	return cherryPrefix + prefix + replySuffix + nodeType + "." + nodeID
}
