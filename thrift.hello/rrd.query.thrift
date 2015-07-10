// namespace java rrd.recv.thrift
namespace go rrd.query.thrift

struct RRDData {
  1: i64 timestamp, // 数据项时间, unix时间戳, 秒为单位
  2: double value,  // 数据项取值, 只支持double类型
}

// 批量数据查询
struct QueryRequest {
  1: i64 startTs,   // 数据起始时间, unix时间戳, 秒为单位
  2: i64 endTs,     // 数据结束时间, unix时间戳, 秒为单位
  3: string cf,     // 归档函数, AVG|LAST|MIN|MAX
  4: string uuid,   // 数据项唯一标识
  5: string dsType, // 数据类型, GAUGE或者COUNTER
  6: i32 step,      // 数据周期, sec
}

struct QueryResponse {
  1: i64 startTs,   // 数据起始时间, unix.timestamp.in.sec
  2: i64 endTs,     // 数据结束时间, unix.timestamp.in.sec
  3: string cf,     // 归档函数, AVG|LAST|MIN|MAX
  4: string uuid,   // 数据项唯一标识
  5: string dsType, // 数据类型, GAUGE或者COUNTER
  6: i32 step,      // 数据周期, sec
  7: list<RRDData> values,  // 数据在指定时间段内的取值
}

// 最新数据查询
struct LastRequest {
  1: string uuid,   // 数据项唯一标识
  2: string dsType, // 数据类型, GAUGE或者COUNTER
  3: i32 step,      // 数据周期, sec
}

struct LastResponse {
  1: string uuid,   // 数据项唯一标识
  2: string dsType, // 数据类型, GAUGE或者COUNTER
  3: i32 step,      // 数据周期, sec
  4: RRDData value, // 最后收到的数据取值
}


service RRDHBaseQuery {
  void ping(),
  // 批量查询接口. 查询指定时间段内的数据、支持指定归档函数
  list<QueryResponse> query(1:list<QueryRequest> requests),
  // 最新数据查询接口. 查询指定数据项 最新上报的数据, COUNTER类型取差值、GAUGE类型取原值, 支持批量查询
  list<LastResponse> last(1:list<LastRequest> request)
}