package main

import (
	"context"
	"fmt"
	ch "test75_clickhouse/clickhouse"
	"testing"
	"time"
)

/*
CREATE TABLE my_first_table
(
    user_id UInt32,
    message String,
    timestamp DateTime,
    metric Float32
)
ENGINE = MergeTree
PRIMARY KEY (user_id, timestamp)
;

INSERT INTO my_first_table (user_id, message, timestamp, metric) VALUES
    (101, '你好,ClickHouse!',                                 now(),       -1.0    ),
    (102, '每批次插入大量数据行',                     yesterday(), 1.41421 ),
    (102, '根据常用查询对数据进行排序', today(),     2.718   ),
    (101, 'Granule 是数据读取的最小单元',      now() + 5,   3.14159 )
;

SELECT *
FROM my_first_table
ORDER BY timestamp
;
*/

type MyFirstTable struct {
	UserID    uint      `clickhouse:"user_id"`
	Message   string    `clickhouse:"message"`
	Timestamp time.Time `clickhouse:"timestamp"`
	Metric    float32   `clickhouse:"metric"`
}

func TestCreate01(t *testing.T) {

	ctx := context.Background()
	_ = ctx

	// 使用 CREATE TABLE 定义新表。
	//常规 SQL DDL 命令在 ClickHouse 中均可使用,但需注意一点——ClickHouse 中的表必须指定 ENGINE 子句。
	//使用 MergeTree 引擎可充分发挥 ClickHouse 的性能优势:

	// 创建表 SQL
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS my_first_table2
		(
			user_id UInt32,
			message String,
			timestamp DateTime,
			metric Float32
		)
		ENGINE = MergeTree
		PRIMARY KEY (user_id, timestamp)
	`
	// 在 ClickHouse 中，表中每一行的主键不要求唯一

	// 执行建表
	result, err := ch.ClickHouseCli.DB.Exec(createTableSQL)
	if err != nil {
		fmt.Println("create table failed:", err)
	}

	fmt.Println("table my_first_table created successfully")
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}

/*
| 类型                     | 说明     | 使用建议     |
| ---------------------- | ------ | -------- |
| UInt32 / UInt64        | 无符号整数  | ID、计数    |
| Int32 / Int64          | 有符号整数  | 差值       |
| Float32 / Float64      | 浮点     | 指标       |
| String                 | 字符串    | 慎用       |
| LowCardinality(String) | 低基数字符串 | **强烈推荐** |
| Date                   | 日期     | 分区       |
| DateTime               | 秒级时间   | 常用       |
| DateTime64(3)          | 毫秒     | 指标/日志    |
| Array(T)               | 数组     | 标签       |
| Map(K,V)               | KV     | 非结构化     |
*/
func TestCreate02(t *testing.T) {
	/*
		CREATE TABLE IF NOT EXISTS table2
		(
			id UInt32,
			name String,
			age Nullable(Int32),
			date_time DateTime,
			timestamp2 TIMESTAMP
		)
		ENGINE = MergeTree
		PARTITION BY toYYYYMM(date_time)
		ORDER BY (id, date_time)
		PRIMARY KEY (id, date_time)
		TTL timestamp + INTERVAL 30 DAY
		SETTINGS index_granularity = 8192;

		age Nullable(Int32), 表示age字段不能为空
		PARTITION BY（分区）分区不是索引，是目录级切分
		ORDER BY 的选型口诀：等值在前，范围在后
		PRIMARY KEY 稀疏索引（mark）通常 = ORDER BY 前缀
		TTL timestamp + INTERVAL 30 DAY：表示删除 30 天以前的数据



	*/
}

func TestCreate03(t *testing.T) {

	/*
		   	CREATE TABLE test_table_base
		      (
		          -- ===== 数值类型 =====
		          i8   Int8,
		          i16  Int16,
		          i32  Int32,
		          i64  Int64,

		          u8   UInt8,
		          u16  UInt16,
		          u32  UInt32,
		          u64  UInt64,

		          f32  Float32,
		          f64  Float64,

		          dec32  Decimal32(4),
		          dec64  Decimal64(8),
		          dec128 Decimal128(18),

		          -- ===== 布尔 =====
		          b Bool,

		          -- ===== 字符串 =====
		          s String,
		          lc LowCardinality(String),
		          fs FixedString(16),

		          -- ===== 用于排序的时间 =====
		          ts DateTime64(3)
		      )
		      ENGINE = MergeTree
		      PARTITION BY toYYYYMM(ts)
		      ORDER BY ts;

		String               → 普通字符串（啥都能放）
		LowCardinality(String) → 枚举/标签/状态值（重复很多）
			表里只存整数 ID
			真正字符串只存一份
			这一列的值一般是固定的一些值，例如：status，level，province...固定的一些枚举值
		FixedString(N)       → 长度固定的二进制字符串（不是给人看的）
			一般是存hash，二进制数据，等一些编码后的数据

	*/
	/*
		INSERT INTO test_table_base
		SELECT
		    number % 127 AS i8,
		    number % 32767 AS i16,
		    number AS i32,
		    number * 1000 AS i64,
		    number % 255 AS u8,
		    number % 65535 AS u16,
		    number AS u32,
		    number * 1000 AS u64,
		    number * 1.1 AS f32,
		    number * 1.111111 AS f64,
		    toDecimal32(number, 4) AS dec32,
		    toDecimal64(number, 8) AS dec64,
		    toDecimal128(number, 18) AS dec128,
		    number % 2 = 0 AS b,
		    concat('str_', toString(number)) AS s,
		    concat('lc_', toString(number % 5)) AS lc,
		    substring(toString(number), 1, 16) AS fs,
		    now64() + INTERVAL number SECOND AS ts
		FROM system.numbers
		LIMIT 10;
	*/
}

func TestCreate05(t *testing.T) {

	/*
	   	CREATE TABLE test_table_advance
	      (
	          -- ===== 时间类型 =====
	          d Date,
	          dt DateTime,
	          dt64 DateTime64(3),

	          -- ===== 枚举 =====
	          e8 Enum8('A' = 1, 'B' = 2),
	          e16 Enum16('X' = 100, 'Y' = 200),

	          -- ===== UUID / IP =====
	          uuid UUID,
	          ipv4 IPv4,
	          ipv6 IPv6,

	          -- ===== 用于排序的时间 =====
	          ts DateTime64(3)
	      )
	      ENGINE = MergeTree
	      PARTITION BY toYYYYMM(ts)
	      ORDER BY ts;

	*/
	/*
		INSERT INTO test_table_advance
		SELECT
		    today() + number AS d,                            -- Date
		    now() + INTERVAL number SECOND AS dt,            -- DateTime
		    now64() + INTERVAL number SECOND AS dt64,        -- DateTime64(3)
		    if(number % 2 = 0, 'A', 'B') AS e8,              -- Enum8
		    if(number % 2 = 0, 'X', 'Y') AS e16,            -- Enum16
		    generateUUIDv4() AS uuid,                        -- UUID
		    IPv4StringToNum(concat('192.168.1.', toString(number % 255))) AS ipv4,  -- IPv4
		    IPv6StringToNum(concat('::1:', toString(number))) AS ipv6,              -- IPv6
		    now64() + INTERVAL number SECOND AS ts           -- 排序时间列
		FROM system.numbers
		LIMIT 10;

	*/
}
func TestCreate06(t *testing.T) {

	/*
	   	CREATE TABLE test_table_complex
	      (
	          -- ===== 复杂类型 =====
	          arr_int Array(Int32),
	          arr_str Array(String),

	          tup Tuple(Int32, String, Float64),

	          mp Map(String, Int32),

	          -- ===== Nullable =====
	          n_int Nullable(Int32),
	          n_str Nullable(String),

	          -- ===== 特殊 =====
	          -- nothing Nothing,

	          -- ===== 用于排序的时间 =====
	          ts DateTime64(3)
	      )
	      ENGINE = MergeTree
	      PARTITION BY toYYYYMM(ts)
	      ORDER BY ts;

	*/

	/*
		INSERT INTO test_table_complex
		SELECT
		    [number, number+1, number+2] AS arr_int,                  -- Array(Int32)
		    ['str_' || toString(number), 'str_' || toString(number+1)] AS arr_str,  -- Array(String)
		    (number, 'tuple_' || toString(number), number*1.1) AS tup,             -- Tuple(Int32, String, Float64)
		    map('k1', number, 'k2', number+10) AS mp,                 -- Map(String, Int32)
		    if(number % 2 = 0, number, NULL) AS n_int,               -- Nullable(Int32)
		    if(number % 2 = 0, 'nullable_' || toString(number), NULL) AS n_str,  -- Nullable(String)
		    now64() + INTERVAL number SECOND AS ts                    -- 排序时间列
		FROM system.numbers
		LIMIT 10;

	*/
}

/*】
ENGINE = MergeTree
PARTITION BY ...
ORDER BY ...
PRIMARY KEY ...
SAMPLE BY ...
TTL ...
SETTINGS ...

ENGINE：存储引擎
PARTITION BY：分区，按时间、hash、表达式
ORDER BY：排序键，决定查询性能和索引
PRIMARY KEY：稀疏索引（一般 = ORDER BY 前缀）
SAMPLE BY：采样列，按列做快速采样
TTL：自动过期/冷热存储
SETTINGS：额外调优参数

| 引擎                                                                | 用途                          |
| ----------------------------------------------------------------- | --------------------------- |
| **MergeTree**                                                     | 最基础，单机 OLAP                 |
| **ReplacingMergeTree([version_column])**                          | 自动去重，可选版本列                  |
| **SummingMergeTree([columns])**                                   | 聚合列自动求和，适合指标表               |
| **AggregatingMergeTree([columns])**                               | 聚合态存储，存储 pre-aggregated 数据  |
| **GraphiteMergeTree(path, retention_policy, aggregation_method)** | 专门做 Graphite 时间序列           |
| **CollapsingMergeTree(sign_column)**                              | 可以 collapse insert 数据，通常做去重 |
| **VersionedCollapsingMergeTree(sign_column, version_column)**     | 支持版本去重                      |
| **StripeLog**                                                     | 写入非常快，几乎不排序，不建议分析表用         |
| **ReplicatedMergeTree**                                           | 集群副本表，支持分布式/副本              |
| **Distributed**                                                   | 虚拟表，跨节点分布式查询                |


*/

/*
| 操作                 | ClickHouse 示例                                                      | 作用                                |
| ------------------ | ------------------------------------------------------------------ | --------------------------------- |
| **CREATE TABLE**   | `CREATE TABLE my_table (id UInt32) ENGINE = MergeTree ORDER BY id` | 创建新表                              |
| **DROP TABLE**     | `DROP TABLE my_table`                                              | 删除表及数据                            |
| **TRUNCATE TABLE** | `TRUNCATE TABLE my_table`                                          | 清空表数据但保留表结构                       |
| **ALTER TABLE**    | `ALTER TABLE my_table ADD COLUMN new_col String`                   | 修改表结构：增/删列、修改类型、重命名列、修改 TTL、移动分区等 |
| **RENAME TABLE**   | `RENAME TABLE old_table TO new_table`                              | 重命名表                              |
*/
func TestCreate08(t *testing.T) {

	ctx := context.Background()
	_ = ctx

	create_sql :=
		`
		create table if not exists test_table_02
		(
			id Int64,
			msg String,
			level LowCardinality(String),
			ts DATETIME64(3)
		)
		engine = MergeTree()
		order by (id, ts)
	`

	execContext, err := ch.ClickHouseCli.DB.ExecContext(ctx, create_sql)
	if err != nil {
		fmt.Println("exec err:", err)
		return
	}
	fmt.Println(execContext)
	fmt.Println(execContext.RowsAffected())
	fmt.Println(execContext.LastInsertId())

	// Exec 或 ExecContext支持的sql：
	// CREATE TABLE, DROP TABLE, ALTER TABLE。返回行结果通常为空。
	// INSERT, UPDATE, DELETE（ClickHouse 不常用 UPDATE/DELETE，因为 ClickHouse 更偏向 OLAP 批量插入）
	// ⚠️注意：不要用SELECT查询语句

}

func TestCreate09(t *testing.T) {
	ctx := context.Background()
	_ = ctx
	// alter table test_table_02 add column age UInt8

	alter_sql := "alter table test_table_02 add column age UInt8"
	execContext, err := ch.ClickHouseCli.DB.ExecContext(ctx, alter_sql)
	if err != nil {
		fmt.Println("exec err:", err)
		return
	}
	fmt.Println(execContext)
	fmt.Println(execContext.RowsAffected())
	fmt.Println(execContext.LastInsertId())

}

func TestCreate10(t *testing.T) {
	ctx := context.Background()
	_ = ctx
	drop_sql := `drop table if exists test_table_02`
	execContext, err := ch.ClickHouseCli.DB.ExecContext(ctx, drop_sql)
	if err != nil {
		fmt.Println("exec err:", err)
		return
	}
	fmt.Println(execContext)
	fmt.Println(execContext.RowsAffected())
	fmt.Println(execContext.LastInsertId())
}
