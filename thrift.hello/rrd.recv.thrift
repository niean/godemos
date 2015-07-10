// namespace java rrd.recv.thrift
namespace go rrd.recv.thrift

struct GraphItem {
  1: string uuid,   // 全局唯一的id
  2: double value,  // 值，只有double一种类型
  3: i64 timestamp, // unix时间戳，秒为单位
  4: string dsType, // 数据类型，GAUGE或者COUNTER
  5: i32 step,      // 数据汇报周期
  6: i32 heartbeat, // 心跳周期，rrd中的概念，数据如果超过heartbeat时间内都没有上来，相应时间点的数据就是NULL了
  7: string min,    // 数据的最小值，比该值小则认为不合法，丢弃，如果该值设置为"U"，表示无限制
  8: string max,    // 数据的最大值，比该值大则认为不合法，丢弃，如果该值设置为"U"，表示无限制
}

service RRDHBaseBackend {
  void ping(),
  // 此处的返回值就是一个ErrorMessage即可，正常情况下留空，异常了就填充为错误消息
  string send(1:list<GraphItem> items)
}